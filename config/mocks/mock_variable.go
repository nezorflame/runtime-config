// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kihamo/runtime-config/config (interfaces: Variable)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	config "github.com/kihamo/runtime-config/config"
	reflect "reflect"
)

// MockVariable is a mock of Variable interface
type MockVariable struct {
	ctrl     *gomock.Controller
	recorder *MockVariableMockRecorder
}

// MockVariableMockRecorder is the mock recorder for MockVariable
type MockVariableMockRecorder struct {
	mock *MockVariable
}

// NewMockVariable creates a new mock instance
func NewMockVariable(ctrl *gomock.Controller) *MockVariable {
	mock := &MockVariable{ctrl: ctrl}
	mock.recorder = &MockVariableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVariable) EXPECT() *MockVariableMockRecorder {
	return m.recorder
}

// Group mocks base method
func (m *MockVariable) Group() string {
	ret := m.ctrl.Call(m, "Group")
	ret0, _ := ret[0].(string)
	return ret0
}

// Group indicates an expected call of Group
func (mr *MockVariableMockRecorder) Group() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockVariable)(nil).Group))
}

// IsEditable mocks base method
func (m *MockVariable) IsEditable() bool {
	ret := m.ctrl.Call(m, "IsEditable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEditable indicates an expected call of IsEditable
func (mr *MockVariableMockRecorder) IsEditable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEditable", reflect.TypeOf((*MockVariable)(nil).IsEditable))
}

// IsOptional mocks base method
func (m *MockVariable) IsOptional() bool {
	ret := m.ctrl.Call(m, "IsOptional")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOptional indicates an expected call of IsOptional
func (mr *MockVariableMockRecorder) IsOptional() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOptional", reflect.TypeOf((*MockVariable)(nil).IsOptional))
}

// Name mocks base method
func (m *MockVariable) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockVariableMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockVariable)(nil).Name))
}

// Type mocks base method
func (m *MockVariable) Type() config.VariableType {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(config.VariableType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockVariableMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockVariable)(nil).Type))
}

// Usage mocks base method
func (m *MockVariable) Usage() string {
	ret := m.ctrl.Call(m, "Usage")
	ret0, _ := ret[0].(string)
	return ret0
}

// Usage indicates an expected call of Usage
func (mr *MockVariableMockRecorder) Usage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Usage", reflect.TypeOf((*MockVariable)(nil).Usage))
}

// Value mocks base method
func (m *MockVariable) Value() config.Value {
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(config.Value)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockVariableMockRecorder) Value() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockVariable)(nil).Value))
}
