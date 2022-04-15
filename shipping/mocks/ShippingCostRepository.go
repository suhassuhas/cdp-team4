// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/domain"
	errs "github.com/swiggy-2022-bootcamp/cdp-team4/shipping/utils/errs"

	mock "github.com/stretchr/testify/mock"
)

// ShippingCostRepository is an autogenerated mock type for the ShippingCostRepository type
type ShippingCostRepository struct {
	mock.Mock
}

// DeleteShippingCostByCity provides a mock function with given fields: _a0
func (_m *ShippingCostRepository) DeleteShippingCostByCity(_a0 string) (bool, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// FindShippingCostByCity provides a mock function with given fields: _a0
func (_m *ShippingCostRepository) FindShippingCostByCity(_a0 string) (*domain.ShippingCost, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.ShippingCost
	if rf, ok := ret.Get(0).(func(string) *domain.ShippingCost); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ShippingCost)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// InsertShippingCost provides a mock function with given fields: _a0
func (_m *ShippingCostRepository) InsertShippingCost(_a0 domain.ShippingCost) (bool, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.ShippingCost) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.ShippingCost) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdateShippingCost provides a mock function with given fields: _a0
func (_m *ShippingCostRepository) UpdateShippingCost(_a0 domain.ShippingCost) (bool, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.ShippingCost) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(domain.ShippingCost) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
