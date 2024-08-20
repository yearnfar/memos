// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package memo is a generated GoMock package.
package memo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yearnfar/memos/internal/module/memo/model"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ListInboxes mocks base method.
func (m *MockService) ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboxes", ctx, req)
	ret0, _ := ret[0].([]*model.Inbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInboxes indicates an expected call of ListInboxes.
func (mr *MockServiceMockRecorder) ListInboxes(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboxes", reflect.TypeOf((*MockService)(nil).ListInboxes), ctx, req)
}
