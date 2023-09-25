// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	block "github.com/skip-mev/block-sdk/block"

	log "cosmossdk.io/log"

	math "cosmossdk.io/math"

	mempool "github.com/cosmos/cosmos-sdk/types/mempool"

	mock "github.com/stretchr/testify/mock"

	proposals "github.com/skip-mev/block-sdk/block/proposals"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Lane is an autogenerated mock type for the Lane type
type Lane struct {
	mock.Mock
}

// Compare provides a mock function with given fields: ctx, this, other
func (_m *Lane) Compare(ctx types.Context, this types.Tx, other types.Tx) int {
	ret := _m.Called(ctx, this, other)

	var r0 int
	if rf, ok := ret.Get(0).(func(types.Context, types.Tx, types.Tx) int); ok {
		r0 = rf(ctx, this, other)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Contains provides a mock function with given fields: tx
func (_m *Lane) Contains(tx types.Tx) bool {
	ret := _m.Called(tx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Tx) bool); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CountTx provides a mock function with given fields:
func (_m *Lane) CountTx() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetMaxBlockSpace provides a mock function with given fields:
func (_m *Lane) GetMaxBlockSpace() math.LegacyDec {
	ret := _m.Called()

	var r0 math.LegacyDec
	if rf, ok := ret.Get(0).(func() math.LegacyDec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	return r0
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *Lane) Insert(_a0 context.Context, _a1 types.Tx) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *Lane) Logger() log.Logger {
	ret := _m.Called()

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func() log.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// Match provides a mock function with given fields: ctx, tx
func (_m *Lane) Match(ctx types.Context, tx types.Tx) bool {
	ret := _m.Called(ctx, tx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.Tx) bool); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Lane) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrepareLane provides a mock function with given fields: ctx, proposal, limit, next
func (_m *Lane) PrepareLane(ctx types.Context, proposal proposals.Proposal, limit proposals.LaneLimits, next block.PrepareLanesHandler) (proposals.Proposal, error) {
	ret := _m.Called(ctx, proposal, limit, next)

	var r0 proposals.Proposal
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, proposals.Proposal, proposals.LaneLimits, block.PrepareLanesHandler) (proposals.Proposal, error)); ok {
		return rf(ctx, proposal, limit, next)
	}
	if rf, ok := ret.Get(0).(func(types.Context, proposals.Proposal, proposals.LaneLimits, block.PrepareLanesHandler) proposals.Proposal); ok {
		r0 = rf(ctx, proposal, limit, next)
	} else {
		r0 = ret.Get(0).(proposals.Proposal)
	}

	if rf, ok := ret.Get(1).(func(types.Context, proposals.Proposal, proposals.LaneLimits, block.PrepareLanesHandler) error); ok {
		r1 = rf(ctx, proposal, limit, next)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessLane provides a mock function with given fields: ctx, partialProposal, next
func (_m *Lane) ProcessLane(ctx types.Context, partialProposal []types.Tx, next block.ProcessLanesHandler) (types.Context, error) {
	ret := _m.Called(ctx, partialProposal, next)

	var r0 types.Context
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, []types.Tx, block.ProcessLanesHandler) (types.Context, error)); ok {
		return rf(ctx, partialProposal, next)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []types.Tx, block.ProcessLanesHandler) types.Context); ok {
		r0 = rf(ctx, partialProposal, next)
	} else {
		r0 = ret.Get(0).(types.Context)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []types.Tx, block.ProcessLanesHandler) error); ok {
		r1 = rf(ctx, partialProposal, next)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: _a0
func (_m *Lane) Remove(_a0 types.Tx) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Tx) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: _a0, _a1
func (_m *Lane) Select(_a0 context.Context, _a1 [][]byte) mempool.Iterator {
	ret := _m.Called(_a0, _a1)

	var r0 mempool.Iterator
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte) mempool.Iterator); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mempool.Iterator)
		}
	}

	return r0
}

// SetAnteHandler provides a mock function with given fields: antehander
func (_m *Lane) SetAnteHandler(antehander types.AnteHandler) {
	_m.Called(antehander)
}

// SetIgnoreList provides a mock function with given fields: ignoreList
func (_m *Lane) SetIgnoreList(ignoreList []block.Lane) {
	_m.Called(ignoreList)
}

// NewLane creates a new instance of Lane. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLane(t interface {
	mock.TestingT
	Cleanup(func())
}) *Lane {
	mock := &Lane{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
