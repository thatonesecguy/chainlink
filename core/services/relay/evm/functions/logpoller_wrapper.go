package functions

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/functions/generated/functions_coordinator"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/functions/generated/functions_router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/functions/config"
	evmRelayTypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type logPollerWrapper struct {
	services.StateMachine

	routerContract            *functions_router.FunctionsRouter
	pluginConfig              config.PluginConfig
	client                    client.Client
	logPoller                 logpoller.LogPoller
	subscribers               map[string]evmRelayTypes.RouteUpdateSubscriber
	activeCoordinator         common.Address
	proposedCoordinator       common.Address
	requestBlockOffset        int64
	responseBlockOffset       int64
	pastBlocksToPoll          int64
	logPollerCacheDurationSec int64
	detectedRequests          detectedEvents
	detectedResponses         detectedEvents
	mu                        sync.Mutex
	closeWait                 sync.WaitGroup
	stopCh                    services.StopChan
	lggr                      logger.Logger
}

type detectedEvent struct {
	requestId    [32]byte
	timeDetected time.Time
}

type detectedEvents struct {
	isPreviouslyDetected  map[[32]byte]struct{}
	detectedEventsOrdered []detectedEvent
}

const logPollerCacheDurationSecDefault = 300
const pastBlocksToPollDefault = 50
const maxLogsToProcess = 1000

var _ evmRelayTypes.LogPollerWrapper = &logPollerWrapper{}

var CommitmentABI = getCommitmentABI()

func getCommitmentABI() abi.Arguments {
	mustNewType := func(t string) abi.Type {
		result, err := abi.NewType(t, t, nil)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
		}
		return result
	}
	return abi.Arguments([]abi.Argument{
		{Type: mustNewType("bytes32")}, // RequestId
		{Type: mustNewType("address")}, // Coordinator
		{Type: mustNewType("uint96")},  // EstimatedTotalCostJuels
		{Type: mustNewType("address")}, // Client
		{Type: mustNewType("uint64")},  // SubscriptionId
		{Type: mustNewType("uint32")},  // CallbackGasLimit
		{Type: mustNewType("uint72")},  // AdminFee
		{Type: mustNewType("uint72")},  // DonFee
		{Type: mustNewType("uint40")},  // GasOverheadBeforeCallback
		{Type: mustNewType("uint40")},  // GasOverheadAfterCallback
		{Type: mustNewType("uint32")},  // TimeoutTimestamp
	})
}

func NewLogPollerWrapper(routerContractAddress common.Address, pluginConfig config.PluginConfig, client client.Client, logPoller logpoller.LogPoller, lggr logger.Logger) (evmRelayTypes.LogPollerWrapper, error) {
	routerContract, err := functions_router.NewFunctionsRouter(routerContractAddress, client)
	if err != nil {
		return nil, err
	}
	blockOffset := int64(pluginConfig.MinIncomingConfirmations) - 1
	if blockOffset < 0 {
		lggr.Warnw("invalid minIncomingConfirmations, using 1 instead", "minIncomingConfirmations", pluginConfig.MinIncomingConfirmations)
		blockOffset = 0
	}
	requestBlockOffset := int64(pluginConfig.MinRequestConfirmations) - 1
	if requestBlockOffset < 0 {
		lggr.Warnw("invalid minRequestConfirmations, using minIncomingConfirmations instead", "minRequestConfirmations", pluginConfig.MinRequestConfirmations)
		requestBlockOffset = blockOffset
	}
	responseBlockOffset := int64(pluginConfig.MinResponseConfirmations) - 1
	if responseBlockOffset < 0 {
		lggr.Warnw("invalid minResponseConfirmations, using minIncomingConfirmations instead", "minResponseConfirmations", pluginConfig.MinResponseConfirmations)
		responseBlockOffset = blockOffset
	}
	logPollerCacheDurationSec := int64(pluginConfig.LogPollerCacheDurationSec)
	if logPollerCacheDurationSec <= 0 {
		lggr.Warnw("invalid logPollerCacheDuration, using 300 instead", "logPollerCacheDurationSec", logPollerCacheDurationSec)
		logPollerCacheDurationSec = logPollerCacheDurationSecDefault
	}
	pastBlocksToPoll := int64(pluginConfig.PastBlocksToPoll)
	if pastBlocksToPoll <= 0 {
		lggr.Warnw("invalid pastBlocksToPoll, using 50 instead", "pastBlocksToPoll", pastBlocksToPoll)
		pastBlocksToPoll = pastBlocksToPollDefault
	}
	if blockOffset >= pastBlocksToPoll || requestBlockOffset >= pastBlocksToPoll || responseBlockOffset >= pastBlocksToPoll {
		lggr.Errorw("invalid config: number of required confirmation blocks >= pastBlocksToPoll", "pastBlocksToPoll", pastBlocksToPoll, "minIncomingConfirmations", pluginConfig.MinIncomingConfirmations, "minRequestConfirmations", pluginConfig.MinRequestConfirmations, "minResponseConfirmations", pluginConfig.MinResponseConfirmations)
		return nil, errors.Errorf("invalid config: number of required confirmation blocks >= pastBlocksToPoll")
	}

	return &logPollerWrapper{
		routerContract:            routerContract,
		pluginConfig:              pluginConfig,
		requestBlockOffset:        requestBlockOffset,
		responseBlockOffset:       responseBlockOffset,
		pastBlocksToPoll:          pastBlocksToPoll,
		logPollerCacheDurationSec: logPollerCacheDurationSec,
		detectedRequests:          detectedEvents{isPreviouslyDetected: make(map[[32]byte]struct{})},
		detectedResponses:         detectedEvents{isPreviouslyDetected: make(map[[32]byte]struct{})},
		logPoller:                 logPoller,
		client:                    client,
		subscribers:               make(map[string]evmRelayTypes.RouteUpdateSubscriber),
		stopCh:                    make(services.StopChan),
		lggr:                      lggr.Named("LogPollerWrapper"),
	}, nil
}

func (l *logPollerWrapper) Start(context.Context) error {
	return l.StartOnce("LogPollerWrapper", func() error {
		l.lggr.Infow("starting LogPollerWrapper", "routerContract", l.routerContract.Address().Hex(), "contractVersion", l.pluginConfig.ContractVersion)
		l.mu.Lock()
		defer l.mu.Unlock()
		if l.pluginConfig.ContractVersion != 1 {
			return errors.New("only contract version 1 is supported")
		}
		l.closeWait.Add(1)
		go l.checkForRouteUpdates()
		return nil
	})
}

func (l *logPollerWrapper) Close() error {
	return l.StopOnce("LogPollerWrapper", func() (err error) {
		l.lggr.Info("closing LogPollerWrapper")
		close(l.stopCh)
		l.closeWait.Wait()
		return nil
	})
}

func (l *logPollerWrapper) HealthReport() map[string]error {
	return map[string]error{l.Name(): l.Ready()}
}

func (l *logPollerWrapper) Name() string { return l.lggr.Name() }

// methods of LogPollerWrapper
func (l *logPollerWrapper) LatestEvents(ctx context.Context) ([]evmRelayTypes.OracleRequest, []evmRelayTypes.OracleResponse, error) {
	l.mu.Lock()
	coordinators := []common.Address{}
	if l.activeCoordinator != (common.Address{}) {
		coordinators = append(coordinators, l.activeCoordinator)
	}
	if l.proposedCoordinator != (common.Address{}) && l.activeCoordinator != l.proposedCoordinator {
		coordinators = append(coordinators, l.proposedCoordinator)
	}
	latest, err := l.logPoller.LatestBlock(ctx)
	if err != nil {
		l.mu.Unlock()
		return nil, nil, err
	}
	latestBlockNum := latest.BlockNumber
	startBlockNum := latestBlockNum - l.pastBlocksToPoll
	if startBlockNum < 0 {
		startBlockNum = 0
	}
	l.mu.Unlock()

	// outside of the lock
	resultsReq := []evmRelayTypes.OracleRequest{}
	resultsResp := []evmRelayTypes.OracleResponse{}
	if len(coordinators) == 0 {
		l.lggr.Debug("LatestEvents: no non-zero coordinators to check")
		return resultsReq, resultsResp, errors.New("no non-zero coordinators to check")
	}

	for _, coordinator := range coordinators {
		requestEndBlock := latestBlockNum - l.requestBlockOffset
		requestLogs, err := l.logPoller.Logs(ctx, startBlockNum, requestEndBlock, functions_coordinator.FunctionsCoordinatorOracleRequest{}.Topic(), coordinator)
		if err != nil {
			l.lggr.Errorw("LatestEvents: fetching request logs from LogPoller failed", "startBlock", startBlockNum, "endBlock", requestEndBlock)
			return nil, nil, err
		}
		l.lggr.Debugw("LatestEvents: fetched request logs", "nRequestLogs", len(requestLogs), "latestBlock", latest, "startBlock", startBlockNum, "endBlock", requestEndBlock)
		requestLogs = l.filterPreviouslyDetectedEvents(requestLogs, &l.detectedRequests, "requests")
		responseEndBlock := latestBlockNum - l.responseBlockOffset
		responseLogs, err := l.logPoller.Logs(ctx, startBlockNum, responseEndBlock, functions_coordinator.FunctionsCoordinatorOracleResponse{}.Topic(), coordinator)
		if err != nil {
			l.lggr.Errorw("LatestEvents: fetching response logs from LogPoller failed", "startBlock", startBlockNum, "endBlock", responseEndBlock)
			return nil, nil, err
		}
		l.lggr.Debugw("LatestEvents: fetched request logs", "nResponseLogs", len(responseLogs), "latestBlock", latest, "startBlock", startBlockNum, "endBlock", responseEndBlock)
		responseLogs = l.filterPreviouslyDetectedEvents(responseLogs, &l.detectedResponses, "responses")

		parsingContract, err := functions_coordinator.NewFunctionsCoordinator(coordinator, l.client)
		if err != nil {
			l.lggr.Error("LatestEvents: creating a contract instance for parsing failed")
			return nil, nil, err
		}

		l.lggr.Debugw("LatestEvents: parsing logs", "nRequestLogs", len(requestLogs), "nResponseLogs", len(responseLogs), "coordinatorAddress", coordinator.Hex())
		for _, log := range requestLogs {
			gethLog := log.ToGethLog()
			oracleRequest, err := parsingContract.ParseOracleRequest(gethLog)
			if err != nil {
				l.lggr.Errorw("LatestEvents: failed to parse a request log, skipping", "err", err)
				continue
			}

			commitmentBytes, err := CommitmentABI.Pack(
				oracleRequest.Commitment.RequestId,
				oracleRequest.Commitment.Coordinator,
				oracleRequest.Commitment.EstimatedTotalCostJuels,
				oracleRequest.Commitment.Client,
				oracleRequest.Commitment.SubscriptionId,
				oracleRequest.Commitment.CallbackGasLimit,
				oracleRequest.Commitment.AdminFee,
				oracleRequest.Commitment.DonFee,
				oracleRequest.Commitment.GasOverheadBeforeCallback,
				oracleRequest.Commitment.GasOverheadAfterCallback,
				oracleRequest.Commitment.TimeoutTimestamp,
			)
			if err != nil {
				l.lggr.Errorw("LatestEvents: failed to pack commitment bytes, skipping", err)
			}

			resultsReq = append(resultsReq, evmRelayTypes.OracleRequest{
				RequestId:           oracleRequest.RequestId,
				RequestingContract:  oracleRequest.RequestingContract,
				RequestInitiator:    oracleRequest.RequestInitiator,
				SubscriptionId:      oracleRequest.SubscriptionId,
				SubscriptionOwner:   oracleRequest.SubscriptionOwner,
				Data:                oracleRequest.Data,
				DataVersion:         oracleRequest.DataVersion,
				Flags:               oracleRequest.Flags,
				CallbackGasLimit:    oracleRequest.CallbackGasLimit,
				TxHash:              oracleRequest.Raw.TxHash,
				OnchainMetadata:     commitmentBytes,
				CoordinatorContract: coordinator,
			})
		}
		for _, log := range responseLogs {
			gethLog := log.ToGethLog()
			oracleResponse, err := parsingContract.ParseOracleResponse(gethLog)
			if err != nil {
				l.lggr.Errorw("LatestEvents: failed to parse a response log, skipping")
				continue
			}
			resultsResp = append(resultsResp, evmRelayTypes.OracleResponse{
				RequestId: oracleResponse.RequestId,
			})
		}
	}

	l.lggr.Debugw("LatestEvents: done", "nRequestLogs", len(resultsReq), "nResponseLogs", len(resultsResp), "startBlock", startBlockNum, "endBlock", latestBlockNum)
	return resultsReq, resultsResp, nil
}

func (l *logPollerWrapper) filterPreviouslyDetectedEvents(logs []logpoller.Log, detectedEvents *detectedEvents, filterType string) []logpoller.Log {
	if len(logs) > maxLogsToProcess {
		l.lggr.Errorw("filterPreviouslyDetectedEvents: too many logs to process, only processing latest maxLogsToProcess logs", "filterType", filterType, "nLogs", len(logs), "maxLogsToProcess", maxLogsToProcess)
		logs = logs[len(logs)-maxLogsToProcess:]
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	filteredLogs := []logpoller.Log{}
	for _, log := range logs {
		var requestId [32]byte
		if len(log.Topics) < 2 || len(log.Topics[1]) != 32 {
			l.lggr.Errorw("filterPreviouslyDetectedEvents: invalid log, skipping", "filterType", filterType, "log", log)
			continue
		}
		copy(requestId[:], log.Topics[1]) // requestId is the second topic (1st topic is the event signature)
		if _, ok := detectedEvents.isPreviouslyDetected[requestId]; !ok {
			filteredLogs = append(filteredLogs, log)
			detectedEvents.isPreviouslyDetected[requestId] = struct{}{}
			detectedEvents.detectedEventsOrdered = append(detectedEvents.detectedEventsOrdered, detectedEvent{requestId: requestId, timeDetected: time.Now()})
		}
	}
	expiredRequests := 0
	for _, detectedEvent := range detectedEvents.detectedEventsOrdered {
		expirationTime := time.Now().Add(-time.Second * time.Duration(l.logPollerCacheDurationSec))
		if !detectedEvent.timeDetected.Before(expirationTime) {
			break
		}
		delete(detectedEvents.isPreviouslyDetected, detectedEvent.requestId)
		expiredRequests++
	}
	detectedEvents.detectedEventsOrdered = detectedEvents.detectedEventsOrdered[expiredRequests:]
	l.lggr.Debugw("filterPreviouslyDetectedEvents: done", "filterType", filterType, "nLogs", len(logs), "nFilteredLogs", len(filteredLogs), "nExpiredRequests", expiredRequests, "previouslyDetectedCacheSize", len(detectedEvents.detectedEventsOrdered))
	return filteredLogs
}

// "internal" method called only by EVM relayer components
func (l *logPollerWrapper) SubscribeToUpdates(ctx context.Context, subscriberName string, subscriber evmRelayTypes.RouteUpdateSubscriber) {
	if l.pluginConfig.ContractVersion == 0 {
		// in V0, immediately set contract address to Oracle contract and never update again
		if err := subscriber.UpdateRoutes(ctx, l.routerContract.Address(), l.routerContract.Address()); err != nil {
			l.lggr.Errorw("LogPollerWrapper: Failed to update routes", "subscriberName", subscriberName, "err", err)
		}
	} else if l.pluginConfig.ContractVersion == 1 {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.subscribers[subscriberName] = subscriber
	}
}

func (l *logPollerWrapper) checkForRouteUpdates() {
	defer l.closeWait.Done()
	freqSec := l.pluginConfig.ContractUpdateCheckFrequencySec
	if freqSec == 0 {
		l.lggr.Errorw("LogPollerWrapper: ContractUpdateCheckFrequencySec is zero - route update checks disabled")
		return
	}

	updateOnce := func() {
		// NOTE: timeout == frequency here, could be changed to a separate config value
		timeout := time.Duration(l.pluginConfig.ContractUpdateCheckFrequencySec) * time.Second
		ctx, cancel := l.stopCh.CtxCancel(context.WithTimeout(context.Background(), timeout))
		defer cancel()
		active, proposed, err := l.getCurrentCoordinators(ctx)
		if err != nil {
			l.lggr.Errorw("LogPollerWrapper: error calling getCurrentCoordinators", "err", err)
			return
		}

		l.handleRouteUpdate(ctx, active, proposed)
	}

	updateOnce() // update once right away
	ticker := time.NewTicker(time.Duration(freqSec) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-l.stopCh:
			return
		case <-ticker.C:
			updateOnce()
		}
	}
}

func (l *logPollerWrapper) getCurrentCoordinators(ctx context.Context) (common.Address, common.Address, error) {
	if l.pluginConfig.ContractVersion == 0 {
		return l.routerContract.Address(), l.routerContract.Address(), nil
	}
	var donId [32]byte
	copy(donId[:], []byte(l.pluginConfig.DONID))

	activeCoordinator, err := l.routerContract.GetContractById(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, donId)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	proposedCoordinator, err := l.routerContract.GetProposedContractById(&bind.CallOpts{
		Pending: false,
		Context: ctx,
	}, donId)
	if err != nil {
		return activeCoordinator, l.proposedCoordinator, nil
	}

	return activeCoordinator, proposedCoordinator, nil
}

func (l *logPollerWrapper) handleRouteUpdate(ctx context.Context, activeCoordinator common.Address, proposedCoordinator common.Address) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if activeCoordinator == (common.Address{}) {
		l.lggr.Error("LogPollerWrapper: cannot update activeCoordinator to zero address")
		return
	}

	if activeCoordinator == l.activeCoordinator && proposedCoordinator == l.proposedCoordinator {
		l.lggr.Debug("LogPollerWrapper: no changes to routes")
		return
	}
	errActive := l.registerFilters(ctx, activeCoordinator)
	errProposed := l.registerFilters(ctx, proposedCoordinator)
	if errActive != nil || errProposed != nil {
		l.lggr.Errorw("LogPollerWrapper: Failed to register filters", "errorActive", errActive, "errorProposed", errProposed)
		return
	}

	l.lggr.Debugw("LogPollerWrapper: new routes", "activeCoordinator", activeCoordinator.Hex(), "proposedCoordinator", proposedCoordinator.Hex())

	l.activeCoordinator = activeCoordinator
	l.proposedCoordinator = proposedCoordinator

	for _, subscriber := range l.subscribers {
		err := subscriber.UpdateRoutes(ctx, activeCoordinator, proposedCoordinator)
		if err != nil {
			l.lggr.Errorw("LogPollerWrapper: Failed to update routes", "err", err)
		}
	}

	filters := l.logPoller.GetFilters()
	for _, filter := range filters {
		if filter.Name[:len(l.filterPrefix())] != l.filterPrefix() {
			continue
		}
		if filter.Name == l.filterName(l.activeCoordinator) || filter.Name == l.filterName(l.proposedCoordinator) {
			continue
		}
		if err := l.logPoller.UnregisterFilter(ctx, filter.Name); err != nil {
			l.lggr.Errorw("LogPollerWrapper: Failed to unregister filter", "filterName", filter.Name, "err", err)
		}
		l.lggr.Debugw("LogPollerWrapper: Successfully unregistered filter", "filterName", filter.Name)
	}
}

func (l *logPollerWrapper) filterPrefix() string {
	return "FunctionsLogPollerWrapper:" + l.pluginConfig.DONID
}

func (l *logPollerWrapper) filterName(addr common.Address) string {
	return logpoller.FilterName(l.filterPrefix(), addr.String())
}

func (l *logPollerWrapper) registerFilters(ctx context.Context, coordinatorAddress common.Address) error {
	if (coordinatorAddress == common.Address{}) {
		return nil
	}
	return l.logPoller.RegisterFilter(
		ctx,
		logpoller.Filter{
			Name: l.filterName(coordinatorAddress),
			EventSigs: []common.Hash{
				functions_coordinator.FunctionsCoordinatorOracleRequest{}.Topic(),
				functions_coordinator.FunctionsCoordinatorOracleResponse{}.Topic(),
			},
			Addresses: []common.Address{coordinatorAddress},
		})
}
