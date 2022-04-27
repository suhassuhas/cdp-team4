// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain (interfaces: TransactionService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/domain"
	errs "github.com/swiggy-2022-bootcamp/cdp-team4/transaction/utils/errs"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionService) CreateTransaction(arg0 string, arg1 int) (string, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionServiceMockRecorder) CreateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionService)(nil).CreateTransaction), arg0, arg1)
}

// GetTransactionById mocks base method.
func (m *MockTransactionService) GetTransactionById(arg0 string) (*domain.Transaction, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionById", arg0)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetTransactionById indicates an expected call of GetTransactionById.
func (mr *MockTransactionServiceMockRecorder) GetTransactionById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionById", reflect.TypeOf((*MockTransactionService)(nil).GetTransactionById), arg0)
}

// GetTransactionByUserId mocks base method.
func (m *MockTransactionService) GetTransactionByUserId(arg0 string) (*domain.Transaction, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByUserId", arg0)
	ret0, _ := ret[0].(*domain.Transaction)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetTransactionByUserId indicates an expected call of GetTransactionByUserId.
func (mr *MockTransactionServiceMockRecorder) GetTransactionByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByUserId", reflect.TypeOf((*MockTransactionService)(nil).GetTransactionByUserId), arg0)
}

// UpdateTransactionByUserId mocks base method.
func (m *MockTransactionService) UpdateTransactionByUserId(arg0 string, arg1 int) (bool, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionByUserId", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// UpdateTransactionByUserId indicates an expected call of UpdateTransactionByUserId.
func (mr *MockTransactionServiceMockRecorder) UpdateTransactionByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionByUserId", reflect.TypeOf((*MockTransactionService)(nil).UpdateTransactionByUserId), arg0, arg1)
}
