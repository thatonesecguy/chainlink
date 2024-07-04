// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"
)

// UniversalEstimatorClient is an autogenerated mock type for the universalEstimatorClient type
type UniversalEstimatorClient struct {
	mock.Mock
}

// FeeHistory provides a mock function with given fields: ctx, blockCount, rewardPercentiles
func (_m *UniversalEstimatorClient) FeeHistory(ctx context.Context, blockCount uint64, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	ret := _m.Called(ctx, blockCount, rewardPercentiles)

	if len(ret) == 0 {
		panic("no return value specified for FeeHistory")
	}

	var r0 *ethereum.FeeHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []float64) (*ethereum.FeeHistory, error)); ok {
		return rf(ctx, blockCount, rewardPercentiles)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []float64) *ethereum.FeeHistory); ok {
		r0 = rf(ctx, blockCount, rewardPercentiles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethereum.FeeHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, []float64) error); ok {
		r1 = rf(ctx, blockCount, rewardPercentiles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *UniversalEstimatorClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUniversalEstimatorClient creates a new instance of UniversalEstimatorClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUniversalEstimatorClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *UniversalEstimatorClient {
	mock := &UniversalEstimatorClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
