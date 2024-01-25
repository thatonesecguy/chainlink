// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	client "github.com/smartcontractkit/chainlink/v2/common/client"

	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"

	fmt "fmt"

	logger "github.com/smartcontractkit/chainlink-common/pkg/logger"

	mock "github.com/stretchr/testify/mock"

	time "time"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
)

// TxmClient is an autogenerated mock type for the TxmClient type
type TxmClient[CHAIN_ID types.ID, ADDR types.Hashable, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee] struct {
	mock.Mock
}

// BatchGetReceipts provides a mock function with given fields: ctx, attempts
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) BatchGetReceipts(ctx context.Context, attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) ([]R, []error, error) {
	ret := _m.Called(ctx, attempts)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetReceipts")
	}

	var r0 []R
	var r1 []error
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) ([]R, []error, error)); ok {
		return rf(ctx, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) []R); ok {
		r0 = rf(ctx, attempts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]R)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) []error); ok {
		r1 = rf(ctx, attempts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r2 = rf(ctx, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BatchGetReceiptsWithFinalizedBlock provides a mock function with given fields: ctx, attempts, useFinalityTag, finalityDepth
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) BatchGetReceiptsWithFinalizedBlock(ctx context.Context, attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], useFinalityTag bool, finalityDepth uint32) (*big.Int, []R, []error, error) {
	ret := _m.Called(ctx, attempts, useFinalityTag, finalityDepth)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetReceiptsWithFinalizedBlock")
	}

	var r0 *big.Int
	var r1 []R
	var r2 []error
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, uint32) (*big.Int, []R, []error, error)); ok {
		return rf(ctx, attempts, useFinalityTag, finalityDepth)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, uint32) *big.Int); ok {
		r0 = rf(ctx, attempts, useFinalityTag, finalityDepth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, uint32) []R); ok {
		r1 = rf(ctx, attempts, useFinalityTag, finalityDepth)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]R)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, uint32) []error); ok {
		r2 = rf(ctx, attempts, useFinalityTag, finalityDepth)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]error)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bool, uint32) error); ok {
		r3 = rf(ctx, attempts, useFinalityTag, finalityDepth)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// BatchSendTransactions provides a mock function with given fields: ctx, attempts, bathSize, lggr
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) BatchSendTransactions(ctx context.Context, attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], bathSize int, lggr logger.SugaredLogger) ([]client.SendTxReturnCode, []error, time.Time, []int64, error) {
	ret := _m.Called(ctx, attempts, bathSize, lggr)

	if len(ret) == 0 {
		panic("no return value specified for BatchSendTransactions")
	}

	var r0 []client.SendTxReturnCode
	var r1 []error
	var r2 time.Time
	var r3 []int64
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) ([]client.SendTxReturnCode, []error, time.Time, []int64, error)); ok {
		return rf(ctx, attempts, bathSize, lggr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) []client.SendTxReturnCode); ok {
		r0 = rf(ctx, attempts, bathSize, lggr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]client.SendTxReturnCode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) []error); ok {
		r1 = rf(ctx, attempts, bathSize, lggr)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) time.Time); ok {
		r2 = rf(ctx, attempts, bathSize, lggr)
	} else {
		r2 = ret.Get(2).(time.Time)
	}

	if rf, ok := ret.Get(3).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) []int64); ok {
		r3 = rf(ctx, attempts, bathSize, lggr)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).([]int64)
		}
	}

	if rf, ok := ret.Get(4).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], int, logger.SugaredLogger) error); ok {
		r4 = rf(ctx, attempts, bathSize, lggr)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// CallContract provides a mock function with given fields: ctx, attempt, blockNumber
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CallContract(ctx context.Context, attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], blockNumber *big.Int) (fmt.Stringer, error) {
	ret := _m.Called(ctx, attempt, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CallContract")
	}

	var r0 fmt.Stringer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *big.Int) (fmt.Stringer, error)); ok {
		return rf(ctx, attempt, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *big.Int) fmt.Stringer); ok {
		r0 = rf(ctx, attempt, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fmt.Stringer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *big.Int) error); ok {
		r1 = rf(ctx, attempt, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfiguredChainID provides a mock function with given fields:
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) ConfiguredChainID() CHAIN_ID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConfiguredChainID")
	}

	var r0 CHAIN_ID
	if rf, ok := ret.Get(0).(func() CHAIN_ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	return r0
}

// FinalizedBlockHash provides a mock function with given fields: ctx
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FinalizedBlockHash(ctx context.Context) (BLOCK_HASH, *big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FinalizedBlockHash")
	}

	var r0 BLOCK_HASH
	var r1 *big.Int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (BLOCK_HASH, *big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) BLOCK_HASH); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(BLOCK_HASH)
	}

	if rf, ok := ret.Get(1).(func(context.Context) *big.Int); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PendingSequenceAt provides a mock function with given fields: ctx, addr
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PendingSequenceAt(ctx context.Context, addr ADDR) (SEQ, error) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for PendingSequenceAt")
	}

	var r0 SEQ
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR) (SEQ, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR) SEQ); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(SEQ)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendEmptyTransaction provides a mock function with given fields: ctx, newTxAttempt, seq, gasLimit, fee, fromAddress
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SendEmptyTransaction(ctx context.Context, newTxAttempt func(SEQ, uint32, FEE, ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error), seq SEQ, gasLimit uint32, fee FEE, fromAddress ADDR) (string, error) {
	ret := _m.Called(ctx, newTxAttempt, seq, gasLimit, fee, fromAddress)

	if len(ret) == 0 {
		panic("no return value specified for SendEmptyTransaction")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(SEQ, uint32, FEE, ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error), SEQ, uint32, FEE, ADDR) (string, error)); ok {
		return rf(ctx, newTxAttempt, seq, gasLimit, fee, fromAddress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(SEQ, uint32, FEE, ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error), SEQ, uint32, FEE, ADDR) string); ok {
		r0 = rf(ctx, newTxAttempt, seq, gasLimit, fee, fromAddress)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(SEQ, uint32, FEE, ADDR) (txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error), SEQ, uint32, FEE, ADDR) error); ok {
		r1 = rf(ctx, newTxAttempt, seq, gasLimit, fee, fromAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransactionReturnCode provides a mock function with given fields: ctx, tx, attempt, lggr
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SendTransactionReturnCode(ctx context.Context, tx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], lggr logger.SugaredLogger) (client.SendTxReturnCode, error) {
	ret := _m.Called(ctx, tx, attempt, lggr)

	if len(ret) == 0 {
		panic("no return value specified for SendTransactionReturnCode")
	}

	var r0 client.SendTxReturnCode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.SugaredLogger) (client.SendTxReturnCode, error)); ok {
		return rf(ctx, tx, attempt, lggr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.SugaredLogger) client.SendTxReturnCode); ok {
		r0 = rf(ctx, tx, attempt, lggr)
	} else {
		r0 = ret.Get(0).(client.SendTxReturnCode)
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], logger.SugaredLogger) error); ok {
		r1 = rf(ctx, tx, attempt, lggr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequenceAt provides a mock function with given fields: ctx, addr, blockNum
func (_m *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SequenceAt(ctx context.Context, addr ADDR, blockNum *big.Int) (SEQ, error) {
	ret := _m.Called(ctx, addr, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for SequenceAt")
	}

	var r0 SEQ
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, *big.Int) (SEQ, error)); ok {
		return rf(ctx, addr, blockNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, *big.Int) SEQ); ok {
		r0 = rf(ctx, addr, blockNum)
	} else {
		r0 = ret.Get(0).(SEQ)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, *big.Int) error); ok {
		r1 = rf(ctx, addr, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTxmClient creates a new instance of TxmClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxmClient[CHAIN_ID types.ID, ADDR types.Hashable, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee](t interface {
	mock.TestingT
	Cleanup(func())
}) *TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE] {
	mock := &TxmClient[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}