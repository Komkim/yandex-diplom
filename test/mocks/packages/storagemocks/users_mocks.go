// Code generated by MockGen. DO NOT EDIT.
// Source: ../../storage/repository/users.go

// Package storagemocks is a generated GoMock package.
package storagemocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsers is a mock of Users interface
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockUsers) Register(login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUsersMockRecorder) Register(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsers)(nil).Register), login, password)
}

// Login mocks base method
func (m *MockUsers) Login(login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockUsersMockRecorder) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsers)(nil).Login), login, password)
}