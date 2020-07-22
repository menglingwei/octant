// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	schema "k8s.io/apimachinery/pkg/runtime/schema"

	portforward "github.com/vmware-tanzu/octant/internal/portforward"
	action "github.com/vmware-tanzu/octant/pkg/action"
)

// MockPortForwarder is a mock of PortForwarder interface
type MockPortForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockPortForwarderMockRecorder
}

// MockPortForwarderMockRecorder is the mock recorder for MockPortForwarder
type MockPortForwarderMockRecorder struct {
	mock *MockPortForwarder
}

// NewMockPortForwarder creates a new mock instance
func NewMockPortForwarder(ctrl *gomock.Controller) *MockPortForwarder {
	mock := &MockPortForwarder{ctrl: ctrl}
	mock.recorder = &MockPortForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortForwarder) EXPECT() *MockPortForwarderMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockPortForwarder) List(ctx context.Context) []portforward.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]portforward.State)
	return ret0
}

// List indicates an expected call of List
func (mr *MockPortForwarderMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPortForwarder)(nil).List), ctx)
}

// Get mocks base method
func (m *MockPortForwarder) Get(id string) (portforward.State, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(portforward.State)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPortForwarderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortForwarder)(nil).Get), id)
}

// Create mocks base method
func (m *MockPortForwarder) Create(ctx context.Context, alerter *action.Alerter, gvk schema.GroupVersionKind, name, namespace string, remotePort uint16) (portforward.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, alerter, gvk, name, namespace, remotePort)
	ret0, _ := ret[0].(portforward.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPortForwarderMockRecorder) Create(ctx, alerter, gvk, name, namespace, remotePort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPortForwarder)(nil).Create), ctx, alerter, gvk, name, namespace, remotePort)
}

// FindTarget mocks base method
func (m *MockPortForwarder) FindTarget(namespace string, gvk schema.GroupVersionKind, name string) ([]portforward.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTarget", namespace, gvk, name)
	ret0, _ := ret[0].([]portforward.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTarget indicates an expected call of FindTarget
func (mr *MockPortForwarderMockRecorder) FindTarget(namespace, gvk, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTarget", reflect.TypeOf((*MockPortForwarder)(nil).FindTarget), namespace, gvk, name)
}

// FindPod mocks base method
func (m *MockPortForwarder) FindPod(namespace string, gvk schema.GroupVersionKind, name string) ([]portforward.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPod", namespace, gvk, name)
	ret0, _ := ret[0].([]portforward.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPod indicates an expected call of FindPod
func (mr *MockPortForwarderMockRecorder) FindPod(namespace, gvk, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPod", reflect.TypeOf((*MockPortForwarder)(nil).FindPod), namespace, gvk, name)
}

// Stop mocks base method
func (m *MockPortForwarder) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockPortForwarderMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPortForwarder)(nil).Stop))
}

// StopForwarder mocks base method
func (m *MockPortForwarder) StopForwarder(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopForwarder", id)
}

// StopForwarder indicates an expected call of StopForwarder
func (mr *MockPortForwarderMockRecorder) StopForwarder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopForwarder", reflect.TypeOf((*MockPortForwarder)(nil).StopForwarder), id)
}
