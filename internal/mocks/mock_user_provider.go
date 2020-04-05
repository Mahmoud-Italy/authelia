// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/internal/authentication (interfaces: UserProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	authentication "github.com/authelia/authelia/internal/authentication"
)

// MockUserProvider is a mock of UserProvider interface
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// CheckUserPassword mocks base method
func (m *MockUserProvider) CheckUserPassword(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPassword", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserPassword indicates an expected call of CheckUserPassword
func (mr *MockUserProviderMockRecorder) CheckUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPassword", reflect.TypeOf((*MockUserProvider)(nil).CheckUserPassword), arg0, arg1)
}

// GetDetails mocks base method
func (m *MockUserProvider) GetDetails(arg0 string) (*authentication.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails", arg0)
	ret0, _ := ret[0].(*authentication.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails
func (mr *MockUserProviderMockRecorder) GetDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockUserProvider)(nil).GetDetails), arg0)
}

// UpdatePassword mocks base method
func (m *MockUserProvider) UpdatePassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserProviderMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserProvider)(nil).UpdatePassword), arg0, arg1)
}
