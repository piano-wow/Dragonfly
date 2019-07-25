// Code generated by MockGen. DO NOT EDIT.
// Source: supernode/daemon/mgr/cdn_mgr.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/dragonflyoss/Dragonfly/apis/types"
)

// MockCDNMgr is a mock of CDNMgr interface
type MockCDNMgr struct {
	ctrl     *gomock.Controller
	recorder *MockCDNMgrMockRecorder
}

// MockCDNMgrMockRecorder is the mock recorder for MockCDNMgr
type MockCDNMgrMockRecorder struct {
	mock *MockCDNMgr
}

// NewMockCDNMgr creates a new mock instance
func NewMockCDNMgr(ctrl *gomock.Controller) *MockCDNMgr {
	mock := &MockCDNMgr{ctrl: ctrl}
	mock.recorder = &MockCDNMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCDNMgr) EXPECT() *MockCDNMgrMockRecorder {
	return m.recorder
}

// TriggerCDN mocks base method
func (m *MockCDNMgr) TriggerCDN(ctx context.Context, taskInfo *types.TaskInfo) (*types.TaskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerCDN", ctx, taskInfo)
	ret0, _ := ret[0].(*types.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerCDN indicates an expected call of TriggerCDN
func (mr *MockCDNMgrMockRecorder) TriggerCDN(ctx, taskInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerCDN", reflect.TypeOf((*MockCDNMgr)(nil).TriggerCDN), ctx, taskInfo)
}

// GetHTTPPath mocks base method
func (m *MockCDNMgr) GetHTTPPath(ctx context.Context, taskID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPPath", ctx, taskID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPPath indicates an expected call of GetHTTPPath
func (mr *MockCDNMgrMockRecorder) GetHTTPPath(ctx, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPPath", reflect.TypeOf((*MockCDNMgr)(nil).GetHTTPPath), ctx, taskID)
}

// GetStatus mocks base method
func (m *MockCDNMgr) GetStatus(ctx context.Context, taskID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", ctx, taskID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockCDNMgrMockRecorder) GetStatus(ctx, taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockCDNMgr)(nil).GetStatus), ctx, taskID)
}

// Delete mocks base method
func (m *MockCDNMgr) Delete(ctx context.Context, taskID string, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, taskID, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCDNMgrMockRecorder) Delete(ctx, taskID, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCDNMgr)(nil).Delete), ctx, taskID, force)
}
