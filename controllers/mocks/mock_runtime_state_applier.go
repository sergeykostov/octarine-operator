// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware/cbcontainers-operator/controllers (interfaces: RuntimeStateApplier)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/vmware/cbcontainers-operator/api/v1"
	options "github.com/vmware/cbcontainers-operator/cbcontainers/state/applyment/options"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockRuntimeStateApplier is a mock of RuntimeStateApplier interface.
type MockRuntimeStateApplier struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeStateApplierMockRecorder
}

// MockRuntimeStateApplierMockRecorder is the mock recorder for MockRuntimeStateApplier.
type MockRuntimeStateApplierMockRecorder struct {
	mock *MockRuntimeStateApplier
}

// NewMockRuntimeStateApplier creates a new mock instance.
func NewMockRuntimeStateApplier(ctrl *gomock.Controller) *MockRuntimeStateApplier {
	mock := &MockRuntimeStateApplier{ctrl: ctrl}
	mock.recorder = &MockRuntimeStateApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntimeStateApplier) EXPECT() *MockRuntimeStateApplierMockRecorder {
	return m.recorder
}

// ApplyDesiredState mocks base method.
func (m *MockRuntimeStateApplier) ApplyDesiredState(arg0 context.Context, arg1 *v1.CBContainersRuntime, arg2 client.Client, arg3 options.OwnerSetter) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDesiredState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyDesiredState indicates an expected call of ApplyDesiredState.
func (mr *MockRuntimeStateApplierMockRecorder) ApplyDesiredState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDesiredState", reflect.TypeOf((*MockRuntimeStateApplier)(nil).ApplyDesiredState), arg0, arg1, arg2, arg3)
}
