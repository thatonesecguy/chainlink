package evm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"go.uber.org/multierr"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	txmgrcommon "github.com/smartcontractkit/chainlink/v2/common/txmgr"
	txm "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/functions/config"
	functionsRelay "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/functions"
	evmRelayTypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type functionsProvider struct {
	services.StateMachine
	configWatcher       *configWatcher
	contractTransmitter ContractTransmitter
	logPollerWrapper    evmRelayTypes.LogPollerWrapper
}

var _ evmRelayTypes.FunctionsProvider = (*functionsProvider)(nil)

func (p *functionsProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	return p.contractTransmitter
}

func (p *functionsProvider) LogPollerWrapper() evmRelayTypes.LogPollerWrapper {
	return p.logPollerWrapper
}

func (p *functionsProvider) FunctionsEvents() commontypes.FunctionsEvents {
	// TODO (FUN-668): implement
	return nil
}

func (p *functionsProvider) Start(ctx context.Context) error {
	return p.StartOnce("FunctionsProvider", func() error {
		if err := p.configWatcher.Start(ctx); err != nil {
			return err
		}
		return p.logPollerWrapper.Start(ctx)
	})
}

func (p *functionsProvider) Close() error {
	return p.StopOnce("FunctionsProvider", func() (err error) {
		err = multierr.Combine(err, p.logPollerWrapper.Close())
		err = multierr.Combine(err, p.configWatcher.Close())
		return
	})
}

// Forward all calls to the underlying configWatcher
func (p *functionsProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return p.configWatcher.OffchainConfigDigester()
}

func (p *functionsProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return p.configWatcher.ContractConfigTracker()
}

func (p *functionsProvider) HealthReport() map[string]error {
	return p.configWatcher.HealthReport()
}

func (p *functionsProvider) Name() string {
	return p.configWatcher.Name()
}

func (p *functionsProvider) ChainReader() commontypes.ContractReader {
	return nil
}

func (p *functionsProvider) Codec() commontypes.Codec {
	return nil
}

func NewFunctionsProvider(ctx context.Context, chain legacyevm.Chain, rargs commontypes.RelayArgs, pargs commontypes.PluginArgs, lggr logger.Logger, ethKeystore keystore.Eth, pluginType functionsRelay.FunctionsPluginType) (evmRelayTypes.FunctionsProvider, error) {
	relayOpts := evmRelayTypes.NewRelayOpts(rargs)
	relayConfig, err := relayOpts.RelayConfig()
	if err != nil {
		return nil, err
	}
	expectedChainID := relayConfig.ChainID.String()
	if expectedChainID != chain.ID().String() {
		return nil, fmt.Errorf("internal error: chain id in spec does not match this relayer's chain: have %s expected %s", relayConfig.ChainID.String(), chain.ID().String())
	}
	if err != nil {
		return nil, err
	}
	if !common.IsHexAddress(rargs.ContractID) {
		return nil, errors.Errorf("invalid contractID, expected hex address")
	}
	var pluginConfig config.PluginConfig
	if err2 := json.Unmarshal(pargs.PluginConfig, &pluginConfig); err2 != nil {
		return nil, err2
	}
	routerContractAddress := common.HexToAddress(rargs.ContractID)
	logPollerWrapper, err := functionsRelay.NewLogPollerWrapper(routerContractAddress, pluginConfig, chain.Client(), chain.LogPoller(), lggr)
	if err != nil {
		return nil, err
	}
	configWatcher, err := newFunctionsConfigProvider(ctx, pluginType, chain, rargs, relayConfig.FromBlock, logPollerWrapper, lggr)
	if err != nil {
		return nil, err
	}
	var contractTransmitter ContractTransmitter
	if relayConfig.SendingKeys != nil {
		contractTransmitter, err = newFunctionsContractTransmitter(ctx, pluginConfig.ContractVersion, rargs, configWatcher, ethKeystore, logPollerWrapper, lggr)
		if err != nil {
			return nil, err
		}
	} else {
		lggr.Warn("no sending keys configured for functions plugin, not starting contract transmitter")
	}
	return &functionsProvider{
		configWatcher:       configWatcher,
		contractTransmitter: contractTransmitter,
		logPollerWrapper:    logPollerWrapper,
	}, nil
}

func newFunctionsConfigProvider(ctx context.Context, pluginType functionsRelay.FunctionsPluginType, chain legacyevm.Chain, args commontypes.RelayArgs, fromBlock uint64, logPollerWrapper evmRelayTypes.LogPollerWrapper, lggr logger.Logger) (*configWatcher, error) {
	if !common.IsHexAddress(args.ContractID) {
		return nil, errors.Errorf("invalid contractID, expected hex address")
	}

	routerContractAddress := common.HexToAddress(args.ContractID)

	cp, err := functionsRelay.NewFunctionsConfigPoller(pluginType, chain.LogPoller(), lggr)
	if err != nil {
		return nil, err
	}
	logPollerWrapper.SubscribeToUpdates(ctx, "FunctionsConfigPoller", cp)

	offchainConfigDigester := functionsRelay.NewFunctionsOffchainConfigDigester(pluginType, chain.ID().Uint64())
	logPollerWrapper.SubscribeToUpdates(ctx, "FunctionsOffchainConfigDigester", offchainConfigDigester)

	return newConfigWatcher(lggr, routerContractAddress, offchainConfigDigester, cp, chain, fromBlock, args.New), nil
}

func newFunctionsContractTransmitter(ctx context.Context, contractVersion uint32, rargs commontypes.RelayArgs, configWatcher *configWatcher, ethKeystore keystore.Eth, logPollerWrapper evmRelayTypes.LogPollerWrapper, lggr logger.Logger) (ContractTransmitter, error) {
	var relayConfig evmRelayTypes.RelayConfig
	if err := json.Unmarshal(rargs.RelayConfig, &relayConfig); err != nil {
		return nil, err
	}
	var fromAddresses []common.Address
	sendingKeys := relayConfig.SendingKeys
	if !relayConfig.EffectiveTransmitterID.Valid {
		return nil, errors.New("EffectiveTransmitterID must be specified")
	}
	effectiveTransmitterAddress := common.HexToAddress(relayConfig.EffectiveTransmitterID.String)

	sendingKeysLength := len(sendingKeys)
	if sendingKeysLength == 0 {
		return nil, errors.New("no sending keys provided")
	}

	// If we are using multiple sending keys, then a forwarder is needed to rotate transmissions.
	// Ensure that this forwarder is not set to a local sending key, and ensure our sending keys are enabled.
	for _, s := range sendingKeys {
		if sendingKeysLength > 1 && s == effectiveTransmitterAddress.String() {
			return nil, errors.New("the transmitter is a local sending key with transaction forwarding enabled")
		}
		if err := ethKeystore.CheckEnabled(ctx, common.HexToAddress(s), configWatcher.chain.Config().EVM().ChainID()); err != nil {
			return nil, errors.Wrap(err, "one of the sending keys given is not enabled")
		}
		fromAddresses = append(fromAddresses, common.HexToAddress(s))
	}

	strategy := txmgrcommon.NewQueueingTxStrategy(rargs.ExternalJobID, relayConfig.DefaultTransactionQueueDepth)

	var checker txm.TransmitCheckerSpec
	if relayConfig.SimulateTransactions {
		checker.CheckerType = txm.TransmitCheckerTypeSimulate
	}

	functionsTransmitter, err := functionsRelay.NewFunctionsContractTransmitter(
		configWatcher.chain.Client(),
		OCR2AggregatorTransmissionContractABI,
		configWatcher.chain.LogPoller(),
		lggr,
		contractVersion,
		configWatcher.chain.TxManager(),
		fromAddresses,
		effectiveTransmitterAddress,
		strategy,
		checker,
		configWatcher.chain.ID(),
		ethKeystore,
	)
	if err != nil {
		return nil, err
	}
	logPollerWrapper.SubscribeToUpdates(ctx, "FunctionsConfigTransmitter", functionsTransmitter)
	return functionsTransmitter, err
}
