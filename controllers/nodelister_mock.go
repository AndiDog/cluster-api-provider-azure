/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ./azuremanagedmachinepool_reconciler.go

// Package mock_controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-11-01/compute"
	gomock "go.uber.org/mock/gomock"
)

// MockNodeLister is a mock of NodeLister interface.
type MockNodeLister struct {
	ctrl     *gomock.Controller
	recorder *MockNodeListerMockRecorder
}

// MockNodeListerMockRecorder is the mock recorder for MockNodeLister.
type MockNodeListerMockRecorder struct {
	mock *MockNodeLister
}

// NewMockNodeLister creates a new mock instance.
func NewMockNodeLister(ctrl *gomock.Controller) *MockNodeLister {
	mock := &MockNodeLister{ctrl: ctrl}
	mock.recorder = &MockNodeListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeLister) EXPECT() *MockNodeListerMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockNodeLister) List(arg0 context.Context, arg1 string) ([]compute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNodeListerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNodeLister)(nil).List), arg0, arg1)
}

// ListInstances mocks base method.
func (m *MockNodeLister) ListInstances(arg0 context.Context, arg1, arg2 string) ([]compute.VirtualMachineScaleSetVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", arg0, arg1, arg2)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockNodeListerMockRecorder) ListInstances(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockNodeLister)(nil).ListInstances), arg0, arg1, arg2)
}
