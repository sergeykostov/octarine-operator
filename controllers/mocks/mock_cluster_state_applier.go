// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware/cbcontainers-operator/controllers (interfaces: ClusterStateApplier)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/vmware/cbcontainers-operator/api/v1"
	models "github.com/vmware/cbcontainers-operator/cbcontainers/models"
	options "github.com/vmware/cbcontainers-operator/cbcontainers/state/applyment/options"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterStateApplier is a mock of ClusterStateApplier interface.
type MockClusterStateApplier struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStateApplierMockRecorder
}

// MockClusterStateApplierMockRecorder is the mock recorder for MockClusterStateApplier.
type MockClusterStateApplierMockRecorder struct {
	mock *MockClusterStateApplier
}

// NewMockClusterStateApplier creates a new mock instance.
func NewMockClusterStateApplier(ctrl *gomock.Controller) *MockClusterStateApplier {
	mock := &MockClusterStateApplier{ctrl: ctrl}
	mock.recorder = &MockClusterStateApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterStateApplier) EXPECT() *MockClusterStateApplierMockRecorder {
	return m.recorder
}

// ApplyDesiredState mocks base method.
func (m *MockClusterStateApplier) ApplyDesiredState(arg0 context.Context, arg1 *v1.CBContainersAgent, arg2 *models.RegistrySecretValues, arg3 client.Client, arg4 options.OwnerSetter) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDesiredState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyDesiredState indicates an expected call of ApplyDesiredState.
func (mr *MockClusterStateApplierMockRecorder) ApplyDesiredState(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDesiredState", reflect.TypeOf((*MockClusterStateApplier)(nil).ApplyDesiredState), arg0, arg1, arg2, arg3, arg4)
}

// GetPriorityClassEmptyK8sObject mocks base method.
func (m *MockClusterStateApplier) GetPriorityClassEmptyK8sObject() client.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriorityClassEmptyK8sObject")
	ret0, _ := ret[0].(client.Object)
	return ret0
}

// GetPriorityClassEmptyK8sObject indicates an expected call of GetPriorityClassEmptyK8sObject.
func (mr *MockClusterStateApplierMockRecorder) GetPriorityClassEmptyK8sObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriorityClassEmptyK8sObject", reflect.TypeOf((*MockClusterStateApplier)(nil).GetPriorityClassEmptyK8sObject))
}
