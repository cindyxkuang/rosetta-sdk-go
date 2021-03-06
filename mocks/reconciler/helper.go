// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package reconciler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// BlockExists provides a mock function with given fields: ctx, block
func (_m *Helper) BlockExists(ctx context.Context, block *types.BlockIdentifier) (bool, error) {
	ret := _m.Called(ctx, block)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *types.BlockIdentifier) bool); ok {
		r0 = rf(ctx, block)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.BlockIdentifier) error); ok {
		r1 = rf(ctx, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComputedBalance provides a mock function with given fields: ctx, account, currency, headBlock
func (_m *Helper) ComputedBalance(ctx context.Context, account *types.AccountIdentifier, currency *types.Currency, headBlock *types.BlockIdentifier) (*types.Amount, *types.BlockIdentifier, error) {
	ret := _m.Called(ctx, account, currency, headBlock)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) *types.Amount); ok {
		r0 = rf(ctx, account, currency, headBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 *types.BlockIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) *types.BlockIdentifier); ok {
		r1 = rf(ctx, account, currency, headBlock)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.BlockIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) error); ok {
		r2 = rf(ctx, account, currency, headBlock)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CurrentBlock provides a mock function with given fields: ctx
func (_m *Helper) CurrentBlock(ctx context.Context) (*types.BlockIdentifier, error) {
	ret := _m.Called(ctx)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context) *types.BlockIdentifier); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LiveBalance provides a mock function with given fields: ctx, account, currency, block
func (_m *Helper) LiveBalance(ctx context.Context, account *types.AccountIdentifier, currency *types.Currency, block *types.BlockIdentifier) (*types.Amount, *types.BlockIdentifier, error) {
	ret := _m.Called(ctx, account, currency, block)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) *types.Amount); ok {
		r0 = rf(ctx, account, currency, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 *types.BlockIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) *types.BlockIdentifier); ok {
		r1 = rf(ctx, account, currency, block)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.BlockIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.AccountIdentifier, *types.Currency, *types.BlockIdentifier) error); ok {
		r2 = rf(ctx, account, currency, block)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
