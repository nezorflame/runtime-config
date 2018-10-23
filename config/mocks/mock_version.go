// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kihamo/runtime-config/config (interfaces: Version)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVersion is a mock of Version interface
type MockVersion struct {
	ctrl     *gomock.Controller
	recorder *MockVersionMockRecorder
}

// MockVersionMockRecorder is the mock recorder for MockVersion
type MockVersionMockRecorder struct {
	mock *MockVersion
}

// NewMockVersion creates a new mock instance
func NewMockVersion(ctrl *gomock.Controller) *MockVersion {
	mock := &MockVersion{ctrl: ctrl}
	mock.recorder = &MockVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVersion) EXPECT() *MockVersionMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockVersion) ID() string {
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockVersionMockRecorder) ID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockVersion)(nil).ID))
}

// ProjectID mocks base method
func (m *MockVersion) ProjectID() string {
	ret := m.ctrl.Call(m, "ProjectID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProjectID indicates an expected call of ProjectID
func (mr *MockVersionMockRecorder) ProjectID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectID", reflect.TypeOf((*MockVersion)(nil).ProjectID))
}
