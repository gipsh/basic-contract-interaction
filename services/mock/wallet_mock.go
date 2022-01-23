// Code generated by MockGen. DO NOT EDIT.
// Source: wallet.go

// Package mock_services is a generated GoMock package.
package mock

import (
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWalletService is a mock of WalletService interface
type MockWalletService struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceMockRecorder
}

// MockWalletServiceMockRecorder is the mock recorder for MockWalletService
type MockWalletServiceMockRecorder struct {
	mock *MockWalletService
}

// NewMockWalletService creates a new mock instance
func NewMockWalletService(ctrl *gomock.Controller) *MockWalletService {
	mock := &MockWalletService{ctrl: ctrl}
	mock.recorder = &MockWalletServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWalletService) EXPECT() *MockWalletServiceMockRecorder {
	return m.recorder
}

// GetPrivateKey mocks base method
func (m *MockWalletService) GetPrivateKey(arg0 common.Address) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateKey", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateKey indicates an expected call of GetPrivateKey
func (mr *MockWalletServiceMockRecorder) GetPrivateKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateKey", reflect.TypeOf((*MockWalletService)(nil).GetPrivateKey), arg0)
}

// AddAddress mocks base method
func (m *MockWalletService) AddAddress(arg0 common.Address, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAddress", arg0, arg1)
}

// AddAddress indicates an expected call of AddAddress
func (mr *MockWalletServiceMockRecorder) AddAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockWalletService)(nil).AddAddress), arg0, arg1)
}

// AddAddressPK mocks base method
func (m *MockWalletService) AddAddressPK(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAddressPK", arg0)
}

// AddAddressPK indicates an expected call of AddAddressPK
func (mr *MockWalletServiceMockRecorder) AddAddressPK(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddressPK", reflect.TypeOf((*MockWalletService)(nil).AddAddressPK), arg0)
}