package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"
	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/common/types"
)

var (
	// ErrInvalidChainID is returned when the chain ID is invalid
	ErrInvalidChainID = errors.New("invalid chain ID")
	// ErrTxnNotFound is returned when a transaction is not found
	ErrTxnNotFound = errors.New("transaction not found")
	// ErrExistingIdempotencyKey is returned when a transaction with the same idempotency key already exists
	ErrExistingIdempotencyKey = errors.New("transaction with idempotency key already exists")
	// ErrAddressNotFound is returned when an address is not found
	ErrAddressNotFound = errors.New("address not found")
	// ErrSequenceNotFound is returned when a sequence is not found
	ErrSequenceNotFound = errors.New("sequence not found")
	// ErrCouldNotGetReceipt is the error string we save if we reach our finality depth for a confirmed transaction without ever getting a receipt
	// This most likely happened because an external wallet used the account for this nonce
	ErrCouldNotGetReceipt = errors.New("could not get receipt")
)

// inMemoryStore is a simple wrapper around a persistent TxStore and KeyStore. It handles all the transaction state in memory.
type inMemoryStore[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
] struct {
	lggr    logger.SugaredLogger
	chainID CHAIN_ID

	maxUnstarted      uint64
	keyStore          txmgrtypes.KeyStore[ADDR, CHAIN_ID, SEQ]
	persistentTxStore txmgrtypes.TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]

	addressStatesLock sync.RWMutex
	addressStates     map[ADDR]*addressState[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]
}

// NewInMemoryStore returns a new inMemoryStore
func NewInMemoryStore[
	CHAIN_ID types.ID,
	ADDR, TX_HASH, BLOCK_HASH types.Hashable,
	R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH],
	SEQ types.Sequence,
	FEE feetypes.Fee,
](
	ctx context.Context,
	lggr logger.SugaredLogger,
	chainID CHAIN_ID,
	keyStore txmgrtypes.KeyStore[ADDR, CHAIN_ID, SEQ],
	persistentTxStore txmgrtypes.TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE],
	config txmgrtypes.InMemoryStoreConfig,
) (*inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE], error) {
	ms := inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{
		lggr:              lggr,
		chainID:           chainID,
		keyStore:          keyStore,
		persistentTxStore: persistentTxStore,

		addressStates: map[ADDR]*addressState[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{},
	}

	ms.maxUnstarted = config.MaxQueued()
	if ms.maxUnstarted <= 0 {
		ms.maxUnstarted = 10000
	}

	addressesToTxs := map[ADDR][]txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{}
	// populate all enabled addresses
	enabledAddresses, err := keyStore.EnabledAddressesForChain(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("new_in_memory_store: %w", err)
	}
	for _, addr := range enabledAddresses {
		addressesToTxs[addr] = []txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{}
	}

	txs, err := persistentTxStore.GetAllTransactions(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("address_state: initialization: %w", err)
	}

	for _, tx := range txs {
		at, exists := addressesToTxs[tx.FromAddress]
		if !exists {
			at = []txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{}
		}
		at = append(at, tx)
		addressesToTxs[tx.FromAddress] = at
	}
	for fromAddr, txs := range addressesToTxs {
		ms.addressStates[fromAddr] = newAddressState[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE](lggr, chainID, fromAddr, ms.maxUnstarted, txs)
	}

	return &ms, nil
}

// CreateTransaction creates a new transaction for a given txRequest.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CreateTransaction(
	ctx context.Context,
	txRequest txmgrtypes.TxRequest[ADDR, TX_HASH],
	chainID CHAIN_ID,
) (
	txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	error,
) {
	return txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{}, nil
}

// FindTxWithIdempotencyKey returns a transaction with the given idempotency key
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxWithIdempotencyKey(ctx context.Context, idempotencyKey string, chainID CHAIN_ID) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

// CheckTxQueueCapacity checks if the queue capacity has been reached for a given address
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CheckTxQueueCapacity(ctx context.Context, fromAddress ADDR, maxQueuedTransactions uint64, chainID CHAIN_ID) error {
	return nil
}

// FindLatestSequence returns the latest sequence number for a given address
// It is used to initialize the in-memory sequence map in the broadcaster
// TODO(jtw): this is until we have a abstracted Sequencer Component which can be used instead
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindLatestSequence(ctx context.Context, fromAddress ADDR, chainID CHAIN_ID) (seq SEQ, err error) {
	return seq, nil
}

// CountUnconfirmedTransactions returns the number of unconfirmed transactions for a given address.
// Unconfirmed transactions are transactions that have been broadcast but not confirmed on-chain.
// NOTE(jtw): used to calculate total inflight transactions
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnconfirmedTransactions(ctx context.Context, fromAddress ADDR, chainID CHAIN_ID) (uint32, error) {
	return 0, nil
}

// CountUnstartedTransactions returns the number of unstarted transactions for a given address.
// Unstarted transactions are transactions that have not been broadcast yet.
// NOTE(jtw): used to calculate total inflight transactions
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnstartedTransactions(ctx context.Context, fromAddress ADDR, chainID CHAIN_ID) (uint32, error) {
	return 0, nil
}

// UpdateTxUnstartedToInProgress updates a transaction from unstarted to in_progress.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxUnstartedToInProgress(
	ctx context.Context,
	tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
) error {
	return nil
}

// GetTxInProgress returns the in_progress transaction for a given address.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetTxInProgress(ctx context.Context, fromAddress ADDR) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

// UpdateTxAttemptInProgressToBroadcast updates a transaction attempt from in_progress to broadcast.
// It also updates the transaction state to unconfirmed.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxAttemptInProgressToBroadcast(
	ctx context.Context,
	tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	newAttemptState txmgrtypes.TxAttemptState,
) error {
	return nil
}

// FindNextUnstartedTransactionFromAddress returns the next unstarted transaction for a given address.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindNextUnstartedTransactionFromAddress(_ context.Context, tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], fromAddress ADDR, chainID CHAIN_ID) error {
	return nil
}

// SaveReplacementInProgressAttempt saves a replacement attempt for a transaction that is in_progress.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveReplacementInProgressAttempt(
	ctx context.Context,
	oldAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	replacementAttempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
) error {
	if oldAttempt.State != txmgrtypes.TxAttemptInProgress || replacementAttempt.State != txmgrtypes.TxAttemptInProgress {
		return fmt.Errorf("expected attempts to be in_progress")
	}
	if oldAttempt.ID == 0 {
		return fmt.Errorf("expected oldAttempt to have an ID")
	}

	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	var as *addressState[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]
	for _, vas := range ms.addressStates {
		if vas.hasTx(oldAttempt.TxID) {
			as = vas
			break
		}
	}
	if as == nil {
		return fmt.Errorf("save_replacement_in_progress_attempt: %w: %q", ErrTxnNotFound, oldAttempt.TxID)
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.SaveReplacementInProgressAttempt(ctx, oldAttempt, replacementAttempt); err != nil {
		return fmt.Errorf("save_replacement_in_progress_attempt: %w", err)
	}

	// Update in memory store
	// delete the old attempt
	as.deleteTxAttempt(oldAttempt.TxID, oldAttempt.ID)
	// add the new attempt
	if err := as.addTxAttempt(*replacementAttempt); err != nil {
		return fmt.Errorf("save_replacement_in_progress_attempt: failed to add a replacement transaction attempt: %w", err)
	}

	return nil
}

// UpdateTxFatalError updates a transaction to fatal_error.
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxFatalError(ctx context.Context, tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	return nil
}

// Close closes the inMemoryStore
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Close() {
}

// Abandon removes all transactions for a given address
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Abandon(ctx context.Context, chainID CHAIN_ID, addr ADDR) error {
	return nil
}

// SetBroadcastBeforeBlockNum sets the broadcast_before_block_num for a given chain ID
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SetBroadcastBeforeBlockNum(ctx context.Context, blockNum int64, chainID CHAIN_ID) error {
	return nil
}

// FindTxAttemptsConfirmedMissingReceipt returns all transactions that are confirmed but missing a receipt
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsConfirmedMissingReceipt(ctx context.Context, chainID CHAIN_ID) (
	[]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	error,
) {
	return nil, nil
}

// UpdateBroadcastAts updates the broadcast_at time for a given set of attempts
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateBroadcastAts(ctx context.Context, now time.Time, txIDs []int64) error {
	return nil
}

// UpdateTxsUnconfirmed updates the unconfirmed transactions for a given set of ids
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxsUnconfirmed(ctx context.Context, txIDs []int64) error {
	return nil
}

// FindTxAttemptsRequiringReceiptFetch returns all transactions that are missing a receipt
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringReceiptFetch(ctx context.Context, chainID CHAIN_ID) (
	attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	err error,
) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesPendingCallback(ctx context.Context, blockNum int64, chainID CHAIN_ID) (
	[]txmgrtypes.ReceiptPlus[R],
	error,
) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxCallbackCompleted(ctx context.Context, pipelineTaskRunRid uuid.UUID, chainId CHAIN_ID) error {
	return nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveFetchedReceipts(ctx context.Context, receipts []R, chainID CHAIN_ID) error {
	return nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesByMetaFieldAndStates(ctx context.Context, metaField string, metaValue string, states []txmgrtypes.TxState, chainID *big.Int) (
	[]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	error,
) {
	return nil, nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithMetaFieldByStates(ctx context.Context, metaField string, states []txmgrtypes.TxState, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithMetaFieldByReceiptBlockNum(ctx context.Context, metaField string, blockNum int64, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithAttemptsAndReceiptsByIdsAndState(ctx context.Context, ids []big.Int, states []txmgrtypes.TxState, chainID *big.Int) (tx []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], err error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PruneUnstartedTxQueue(ctx context.Context, queueSize uint32, subject uuid.UUID) ([]int64, error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) ReapTxHistory(ctx context.Context, minBlockNumberToKeep int64, timeThreshold time.Time, chainID CHAIN_ID) error {
	if ms.chainID.String() != chainID.String() {
		panic("invalid chain ID")
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.ReapTxHistory(ctx, minBlockNumberToKeep, timeThreshold, chainID); err != nil {
		return err
	}

	// Update in memory store
	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	for _, as := range ms.addressStates {
		as.reapConfirmedTxs(minBlockNumberToKeep, timeThreshold)
		as.reapFatalErroredTxs(timeThreshold)
	}

	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountTransactionsByState(_ context.Context, state txmgrtypes.TxState, chainID CHAIN_ID) (uint32, error) {
	return 0, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) DeleteInProgressAttempt(ctx context.Context, attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	if attempt.State != txmgrtypes.TxAttemptInProgress {
		return fmt.Errorf("DeleteInProgressAttempt: expected attempt state to be in_progress")
	}
	if attempt.ID == 0 {
		return fmt.Errorf("DeleteInProgressAttempt: expected attempt to have an id")
	}

	// Check if fromaddress enabled
	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	var as *addressState[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]
	for _, vas := range ms.addressStates {
		if vas.hasTx(attempt.TxID) {
			as = vas
			break
		}
	}
	if as == nil {
		return fmt.Errorf("delete_in_progress_attempt: %w: %q", ErrTxnNotFound, attempt.TxID)
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.DeleteInProgressAttempt(ctx, attempt); err != nil {
		return fmt.Errorf("delete_in_progress_attempt: %w", err)
	}

	// Update in memory store
	as.deleteTxAttempt(attempt.TxID, attempt.ID)

	return nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringResubmissionDueToInsufficientFunds(_ context.Context, address ADDR, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringResend(_ context.Context, olderThan time.Time, maxInFlightTransactions uint32, chainID CHAIN_ID, address ADDR) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxWithSequence(_ context.Context, fromAddress ADDR, seq SEQ) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTransactionsConfirmedInBlockRange(_ context.Context, highBlockNumber, lowBlockNumber int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindEarliestUnconfirmedBroadcastTime(ctx context.Context, chainID CHAIN_ID) (null.Time, error) {
	return null.Time{}, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindEarliestUnconfirmedTxAttemptBlock(ctx context.Context, chainID CHAIN_ID) (null.Int, error) {
	return null.Int{}, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetInProgressTxAttempts(ctx context.Context, address ADDR, chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetNonFatalTransactions(ctx context.Context, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetTxByID(_ context.Context, id int64) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) HasInProgressTransaction(_ context.Context, account ADDR, chainID CHAIN_ID) (bool, error) {
	return false, nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) LoadTxAttempts(_ context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PreloadTxes(_ context.Context, attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveConfirmedMissingReceiptAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInProgressAttempt(ctx context.Context, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInsufficientFundsAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveSentAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	return nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxForRebroadcast(ctx context.Context, etx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], etxAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	as, ok := ms.addressStates[etx.FromAddress]
	if !ok {
		return nil
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.UpdateTxForRebroadcast(ctx, etx, etxAttempt); err != nil {
		return err
	}

	// Update in memory store
	return as.moveConfirmedToUnconfirmed(etxAttempt)
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) IsTxFinalized(ctx context.Context, blockHeight int64, txID int64, chainID CHAIN_ID) (bool, error) {
	return false, nil
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringGasBump(ctx context.Context, address ADDR, blockNum, gasBumpThreshold, depth int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	return nil, nil
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkAllConfirmedMissingReceipt(ctx context.Context, chainID CHAIN_ID) error {
	if ms.chainID.String() != chainID.String() {
		panic("invalid chain ID")
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.MarkAllConfirmedMissingReceipt(ctx, chainID); err != nil {
		return err
	}

	// Update in memory store
	var errs error
	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	for _, as := range ms.addressStates {
		// Get the max confirmed sequence
		filter := func(tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) bool { return true }
		states := []txmgrtypes.TxState{TxConfirmed}
		txs := as.findTxs(states, filter)
		var maxConfirmedSequence SEQ
		for _, tx := range txs {
			if tx.Sequence == nil {
				continue
			}
			if (*tx.Sequence).Int64() > maxConfirmedSequence.Int64() {
				maxConfirmedSequence = *tx.Sequence
			}
		}

		// Mark all unconfirmed txs with a sequence less than the max confirmed sequence as confirmed_missing_receipt
		filter = func(tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) bool {
			if tx.Sequence == nil {
				return false
			}

			return (*tx.Sequence).Int64() < maxConfirmedSequence.Int64()
		}
		states = []txmgrtypes.TxState{TxUnconfirmed}
		txs = as.findTxs(states, filter)
		for _, tx := range txs {
			if err := as.moveUnconfirmedToConfirmedMissingReceipt(tx.ID); err != nil {
				err = fmt.Errorf("mark_all_confirmed_missing_receipt: address: %s: %w", as.fromAddress, err)
				errs = errors.Join(errs, err)
			}
		}
	}

	return errs
}
func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkOldTxesMissingReceiptAsErrored(ctx context.Context, blockNum int64, finalityDepth uint32, chainID CHAIN_ID) error {
	if ms.chainID.String() != chainID.String() {
		panic(fmt.Sprintf(ErrInvalidChainID.Error()+": %s", chainID.String()))
	}

	// Persist to persistent storage
	if err := ms.persistentTxStore.MarkOldTxesMissingReceiptAsErrored(ctx, blockNum, finalityDepth, chainID); err != nil {
		return err
	}

	// Update in memory store
	type result struct {
		ID                         int64
		Sequence                   SEQ
		FromAddress                ADDR
		MaxBroadcastBeforeBlockNum int64
		TxHashes                   []TX_HASH
	}
	var results []result
	cutoff := blockNum - int64(finalityDepth)
	if cutoff <= 0 {
		return nil
	}
	filter := func(tx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) bool {
		if len(tx.TxAttempts) == 0 {
			return false
		}
		var maxBroadcastBeforeBlockNum int64
		for i := 0; i < len(tx.TxAttempts); i++ {
			txAttempt := tx.TxAttempts[i]
			if txAttempt.BroadcastBeforeBlockNum == nil {
				continue
			}
			if *txAttempt.BroadcastBeforeBlockNum > maxBroadcastBeforeBlockNum {
				maxBroadcastBeforeBlockNum = *txAttempt.BroadcastBeforeBlockNum
			}
		}
		return maxBroadcastBeforeBlockNum < cutoff
	}
	var errs error
	ms.addressStatesLock.RLock()
	defer ms.addressStatesLock.RUnlock()
	for _, as := range ms.addressStates {
		states := []txmgrtypes.TxState{TxConfirmedMissingReceipt}
		txs := as.findTxs(states, filter)
		for _, tx := range txs {
			if err := as.moveTxToFatalError(tx.ID, null.StringFrom(ErrCouldNotGetReceipt.Error())); err != nil {
				err = fmt.Errorf("mark_old_txes_missing_receipt_as_errored: address: %s: %w", as.fromAddress, err)
				errs = errors.Join(errs, err)
				continue
			}
			hashes := make([]TX_HASH, len(tx.TxAttempts))
			maxBroadcastBeforeBlockNum := int64(0)
			for i, attempt := range tx.TxAttempts {
				hashes[i] = attempt.Hash
				if attempt.BroadcastBeforeBlockNum != nil {
					if *attempt.BroadcastBeforeBlockNum > maxBroadcastBeforeBlockNum {
						maxBroadcastBeforeBlockNum = *attempt.BroadcastBeforeBlockNum
					}
				}
			}
			rr := result{
				ID:                         tx.ID,
				Sequence:                   *tx.Sequence,
				FromAddress:                tx.FromAddress,
				MaxBroadcastBeforeBlockNum: maxBroadcastBeforeBlockNum,
				TxHashes:                   hashes,
			}
			results = append(results, rr)
		}
	}

	for _, r := range results {
		ms.lggr.Criticalw(fmt.Sprintf("eth_tx with ID %v expired without ever getting a receipt for any of our attempts. "+
			"Current block height is %v, transaction was broadcast before block height %v. This transaction may not have not been sent and will be marked as fatally errored. "+
			"This can happen if there is another instance of chainlink running that is using the same private key, or if "+
			"an external wallet has been used to send a transaction from account %s with nonce %v."+
			" Please note that Chainlink requires exclusive ownership of it's private keys and sharing keys across multiple"+
			" chainlink instances, or using the chainlink keys with an external wallet is NOT SUPPORTED and WILL lead to missed transactions",
			r.ID, blockNum, r.MaxBroadcastBeforeBlockNum, r.FromAddress, r.Sequence), "ethTxID", r.ID, "sequence", r.Sequence, "fromAddress", r.FromAddress, "txHashes", r.TxHashes)
	}

	return errs
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) deepCopyTx(
	tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE] {
	copyTx := txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{
		ID:                 tx.ID,
		IdempotencyKey:     tx.IdempotencyKey,
		Sequence:           tx.Sequence,
		FromAddress:        tx.FromAddress,
		ToAddress:          tx.ToAddress,
		EncodedPayload:     make([]byte, len(tx.EncodedPayload)),
		Value:              *new(big.Int).Set(&tx.Value),
		FeeLimit:           tx.FeeLimit,
		Error:              tx.Error,
		BroadcastAt:        tx.BroadcastAt,
		InitialBroadcastAt: tx.InitialBroadcastAt,
		CreatedAt:          tx.CreatedAt,
		State:              tx.State,
		TxAttempts:         make([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], len(tx.TxAttempts)),
		Meta:               tx.Meta,
		Subject:            tx.Subject,
		ChainID:            tx.ChainID,
		PipelineTaskRunID:  tx.PipelineTaskRunID,
		MinConfirmations:   tx.MinConfirmations,
		TransmitChecker:    tx.TransmitChecker,
		SignalCallback:     tx.SignalCallback,
		CallbackCompleted:  tx.CallbackCompleted,
	}

	// Copy the EncodedPayload
	copy(copyTx.EncodedPayload, tx.EncodedPayload)

	// Copy the TxAttempts
	for i, attempt := range tx.TxAttempts {
		copyTx.TxAttempts[i] = ms.deepCopyTxAttempt(copyTx, attempt)
	}

	return &copyTx
}

func (ms *inMemoryStore[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) deepCopyTxAttempt(
	tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
	attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE],
) txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE] {
	copyAttempt := txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]{
		ID:                      attempt.ID,
		TxID:                    attempt.TxID,
		Tx:                      tx,
		TxFee:                   attempt.TxFee,
		ChainSpecificFeeLimit:   attempt.ChainSpecificFeeLimit,
		SignedRawTx:             make([]byte, len(attempt.SignedRawTx)),
		Hash:                    attempt.Hash,
		CreatedAt:               attempt.CreatedAt,
		BroadcastBeforeBlockNum: attempt.BroadcastBeforeBlockNum,
		State:                   attempt.State,
		Receipts:                make([]txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], len(attempt.Receipts)),
		TxType:                  attempt.TxType,
	}

	copy(copyAttempt.SignedRawTx, attempt.SignedRawTx)
	copy(copyAttempt.Receipts, attempt.Receipts)

	return copyAttempt
}