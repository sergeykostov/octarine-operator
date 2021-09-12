// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware/cbcontainers-operator/cbcontainers/state/cluster (interfaces: ClusterChildK8sObjectApplier)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/vmware/cbcontainers-operator/api/v1"
	options "github.com/vmware/cbcontainers-operator/cbcontainers/state/applyment/options"
	cluster "github.com/vmware/cbcontainers-operator/cbcontainers/state/cluster"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterChildK8sObjectApplier is a mock of ClusterChildK8sObjectApplier interface.
type MockClusterChildK8sObjectApplier struct {
	ctrl     *gomock.Controller
	recorder *MockClusterChildK8sObjectApplierMockRecorder
}

// MockClusterChildK8sObjectApplierMockRecorder is the mock recorder for MockClusterChildK8sObjectApplier.
type MockClusterChildK8sObjectApplierMockRecorder struct {
	mock *MockClusterChildK8sObjectApplier
}

// NewMockClusterChildK8sObjectApplier creates a new mock instance.
func NewMockClusterChildK8sObjectApplier(ctrl *gomock.Controller) *MockClusterChildK8sObjectApplier {
	mock := &MockClusterChildK8sObjectApplier{ctrl: ctrl}
	mock.recorder = &MockClusterChildK8sObjectApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterChildK8sObjectApplier) EXPECT() *MockClusterChildK8sObjectApplierMockRecorder {
	return m.recorder
}

// ApplyClusterChildK8sObject mocks base method.
func (m *MockClusterChildK8sObjectApplier) ApplyClusterChildK8sObject(arg0 context.Context, arg1 *v1.CBContainersAgent, arg2 client.Client, arg3 cluster.ClusterChildK8sObject, arg4 ...*options.ApplyOptions) (bool, client.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyClusterChildK8sObject", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(client.Object)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApplyClusterChildK8sObject indicates an expected call of ApplyClusterChildK8sObject.
func (mr *MockClusterChildK8sObjectApplierMockRecorder) ApplyClusterChildK8sObject(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyClusterChildK8sObject", reflect.TypeOf((*MockClusterChildK8sObjectApplier)(nil).ApplyClusterChildK8sObject), varargs...)
}
