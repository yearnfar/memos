// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yearnfar/memos/internal/module/auth/model"
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

// Authenticate mocks base method.
func (m *MockService) Authenticate(ctx context.Context, tokenStr, keyId string) (*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, tokenStr, keyId)
	ret0, _ := ret[0].(*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockServiceMockRecorder) Authenticate(ctx, tokenStr, keyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockService)(nil).Authenticate), ctx, tokenStr, keyId)
}

// GenerateAccessToken mocks base method.
func (m *MockService) GenerateAccessToken(ctx context.Context, userId int32, audience, keyId string, expirationTime time.Time) (*model.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, userId, audience, keyId, expirationTime)
	ret0, _ := ret[0].(*model.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockServiceMockRecorder) GenerateAccessToken(ctx, userId, audience, keyId, expirationTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockService)(nil).GenerateAccessToken), ctx, userId, audience, keyId, expirationTime)
}

// SignIn mocks base method.
func (m *MockService) SignIn(ctx context.Context, req *model.SignInRequest) (*model.SignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, req)
	ret0, _ := ret[0].(*model.SignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockServiceMockRecorder) SignIn(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockService)(nil).SignIn), ctx, req)
}
