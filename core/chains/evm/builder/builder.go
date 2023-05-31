package builder

import (
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/forwarders"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

func NewTxm(
	db *sqlx.DB,
	cfg txmgr.Config,
	client evmclient.Client,
	lggr logger.Logger,
	logPoller logpoller.LogPoller,
	keyStore keystore.Eth,
	eventBroadcaster pg.EventBroadcaster,
	estimator gas.EvmFeeEstimator,
) (txm txmgr.EvmTxManager,
	err error,
) {
	var fwdMgr txmgr.EvmFwdMgr

	if cfg.EvmUseForwarders() {
		fcfg := forwarders.Config{
			EvmFinalityDepth:    int(cfg.EvmFinalityDepth()),
			LogSQL:              cfg.Database().LogSQL,
			DefaultQueryTimeout: cfg.Database().DatabaseDefaultQueryTimeout(),
		}
		fwdMgr = forwarders.NewFwdMgr(db, client, logPoller, lggr, fcfg)
	} else {
		lggr.Info("EvmForwarderManager: Disabled")
	}
	checker := &txmgr.CheckerFactory{Client: client}
	// create tx attempt builder
	txAttemptBuilder := txmgr.NewEvmTxAttemptBuilder(*client.ConfiguredChainID(), cfg, keyStore, estimator)
	txStore := txmgr.NewTxStore(db, lggr, cfg.Database())
	txNonceSyncer := txmgr.NewNonceSyncer(txStore, lggr, client, keyStore)

	txmCfg := txmgr.NewEvmTxmConfig(cfg)       // wrap Evm specific config
	txmClient := txmgr.NewEvmTxmClient(client) // wrap Evm specific client
	broadcasterCfg := types.BroadcasterConfig[*assets.Wei]{
		FallbackPollInterval:    txmCfg.FallbackPollInterval(),
		MaxInFlightTransactions: txmCfg.MaxInFlightTransactions(),
		IsL2:                    txmCfg.IsL2(),
		MaxFeePrice:             txmCfg.MaxFeePrice(),
		FeePriceDefault:         txmCfg.FeePriceDefault(),
	}
	ethBroadcaster := txmgr.NewEvmBroadcaster(txStore, txmClient, broadcasterCfg, keyStore, eventBroadcaster, txAttemptBuilder, txNonceSyncer, lggr, checker, cfg.EvmNonceAutoSync())

	ethConfirmerCfg := types.ConfirmerConfig[*assets.Wei]{
		RPCDefaultBatchSize:     txmCfg.RPCDefaultBatchSize(),
		UseForwarders:           txmCfg.UseForwarders(),
		FeeBumpTxDepth:          txmCfg.FeeBumpTxDepth(),
		MaxInFlightTransactions: txmCfg.MaxInFlightTransactions(),
		FeeLimitDefault:         txmCfg.FeeLimitDefault(),

		FeeBumpThreshold: txmCfg.FeeBumpThreshold(),
		FinalityDepth:    txmCfg.FinalityDepth(),
		MaxFeePrice:      txmCfg.MaxFeePrice(),
		FeeBumpPercent:   txmCfg.FeeBumpPercent(),

		DefaultQueryTimeout: txmCfg.Database().DatabaseDefaultQueryTimeout(),
	}
	ethConfirmer := txmgr.NewEvmConfirmer(txStore, txmClient, ethConfirmerCfg, keyStore, txAttemptBuilder, lggr)
	var ethResender *txmgr.EvmResender
	if cfg.EthTxResendAfterThreshold() > 0 {
		ethResender = txmgr.NewEvmResender(lggr, txStore, txmClient, keyStore, txmgr.DefaultResenderPollInterval, txmCfg)
	}
	txm = txmgr.NewEvmTxm(txmClient.ConfiguredChainID(), txmCfg, keyStore, lggr, checker, fwdMgr, txAttemptBuilder, txStore, txNonceSyncer, ethBroadcaster, ethConfirmer, ethResender)
	return txm, nil
}
