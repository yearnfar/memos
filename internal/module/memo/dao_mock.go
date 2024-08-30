// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package memo is a generated GoMock package.
package memo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yearnfar/memos/internal/module/memo/model"
)

// MockDAO is a mock of DAO interface.
type MockDAO struct {
	ctrl     *gomock.Controller
	recorder *MockDAOMockRecorder
}

// MockDAOMockRecorder is the mock recorder for MockDAO.
type MockDAOMockRecorder struct {
	mock *MockDAO
}

// NewMockDAO creates a new mock instance.
func NewMockDAO(ctrl *gomock.Controller) *MockDAO {
	mock := &MockDAO{ctrl: ctrl}
	mock.recorder = &MockDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAO) EXPECT() *MockDAOMockRecorder {
	return m.recorder
}

// CreateMemo mocks base method.
func (m *MockDAO) CreateMemo(ctx context.Context, memo *model.Memo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMemo", ctx, memo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMemo indicates an expected call of CreateMemo.
func (mr *MockDAOMockRecorder) CreateMemo(ctx, memo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMemo", reflect.TypeOf((*MockDAO)(nil).CreateMemo), ctx, memo)
}

// CreateReaction mocks base method.
func (m_2 *MockDAO) CreateReaction(ctx context.Context, m *model.Reaction) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateReaction", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReaction indicates an expected call of CreateReaction.
func (mr *MockDAOMockRecorder) CreateReaction(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReaction", reflect.TypeOf((*MockDAO)(nil).CreateReaction), ctx, m)
}

// CreateResource mocks base method.
func (m_2 *MockDAO) CreateResource(ctx context.Context, m *model.Resource) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateResource", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResource indicates an expected call of CreateResource.
func (mr *MockDAOMockRecorder) CreateResource(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResource", reflect.TypeOf((*MockDAO)(nil).CreateResource), ctx, m)
}

// DeleteMemoById mocks base method.
func (m *MockDAO) DeleteMemoById(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMemoById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMemoById indicates an expected call of DeleteMemoById.
func (mr *MockDAOMockRecorder) DeleteMemoById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMemoById", reflect.TypeOf((*MockDAO)(nil).DeleteMemoById), ctx, id)
}

// DeleteMemoRelations mocks base method.
func (m *MockDAO) DeleteMemoRelations(ctx context.Context, req *model.DeleteMemoRelationsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMemoRelations", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMemoRelations indicates an expected call of DeleteMemoRelations.
func (mr *MockDAOMockRecorder) DeleteMemoRelations(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMemoRelations", reflect.TypeOf((*MockDAO)(nil).DeleteMemoRelations), ctx, req)
}

// DeleteResourceById mocks base method.
func (m *MockDAO) DeleteResourceById(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourceById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResourceById indicates an expected call of DeleteResourceById.
func (mr *MockDAOMockRecorder) DeleteResourceById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceById", reflect.TypeOf((*MockDAO)(nil).DeleteResourceById), ctx, id)
}

// FindInboxes mocks base method.
func (m *MockDAO) FindInboxes(ctx context.Context, req *model.FindInboxesRequest) ([]*model.Inbox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInboxes", ctx, req)
	ret0, _ := ret[0].([]*model.Inbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInboxes indicates an expected call of FindInboxes.
func (mr *MockDAOMockRecorder) FindInboxes(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInboxes", reflect.TypeOf((*MockDAO)(nil).FindInboxes), ctx, req)
}

// FindMemo mocks base method.
func (m *MockDAO) FindMemo(ctx context.Context, req *model.FindMemoRequest) (*model.Memo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemo", ctx, req)
	ret0, _ := ret[0].(*model.Memo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemo indicates an expected call of FindMemo.
func (mr *MockDAOMockRecorder) FindMemo(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemo", reflect.TypeOf((*MockDAO)(nil).FindMemo), ctx, req)
}

// FindMemoOrganizers mocks base method.
func (m *MockDAO) FindMemoOrganizers(ctx context.Context, req *model.FindMemoOrganizersRequest) ([]*model.MemoOrganizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemoOrganizers", ctx, req)
	ret0, _ := ret[0].([]*model.MemoOrganizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoOrganizers indicates an expected call of FindMemoOrganizers.
func (mr *MockDAOMockRecorder) FindMemoOrganizers(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoOrganizers", reflect.TypeOf((*MockDAO)(nil).FindMemoOrganizers), ctx, req)
}

// FindMemoRelations mocks base method.
func (m *MockDAO) FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) ([]*model.MemoRelation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemoRelations", ctx, req)
	ret0, _ := ret[0].([]*model.MemoRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoRelations indicates an expected call of FindMemoRelations.
func (mr *MockDAOMockRecorder) FindMemoRelations(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoRelations", reflect.TypeOf((*MockDAO)(nil).FindMemoRelations), ctx, req)
}

// FindMemos mocks base method.
func (m *MockDAO) FindMemos(ctx context.Context, req *model.FindMemosRequest) ([]*model.Memo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemos", ctx, req)
	ret0, _ := ret[0].([]*model.Memo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemos indicates an expected call of FindMemos.
func (mr *MockDAOMockRecorder) FindMemos(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemos", reflect.TypeOf((*MockDAO)(nil).FindMemos), ctx, req)
}

// FindReactions mocks base method.
func (m *MockDAO) FindReactions(ctx context.Context, req *model.FindReactionsRequest) ([]*model.Reaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindReactions", ctx, req)
	ret0, _ := ret[0].([]*model.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReactions indicates an expected call of FindReactions.
func (mr *MockDAOMockRecorder) FindReactions(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReactions", reflect.TypeOf((*MockDAO)(nil).FindReactions), ctx, req)
}

// FindResource mocks base method.
func (m *MockDAO) FindResource(ctx context.Context, req *model.FindResourceRequest) (*model.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindResource", ctx, req)
	ret0, _ := ret[0].(*model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResource indicates an expected call of FindResource.
func (mr *MockDAOMockRecorder) FindResource(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResource", reflect.TypeOf((*MockDAO)(nil).FindResource), ctx, req)
}

// FindResources mocks base method.
func (m *MockDAO) FindResources(ctx context.Context, req *model.FindResourcesRequest) ([]*model.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindResources", ctx, req)
	ret0, _ := ret[0].([]*model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResources indicates an expected call of FindResources.
func (mr *MockDAOMockRecorder) FindResources(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResources", reflect.TypeOf((*MockDAO)(nil).FindResources), ctx, req)
}

// FindWorkspaceSettings mocks base method.
func (m *MockDAO) FindWorkspaceSettings(ctx context.Context, req *model.FindWorkspaceSettingsRequest) ([]*model.WorkspaceSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWorkspaceSettings", ctx, req)
	ret0, _ := ret[0].([]*model.WorkspaceSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWorkspaceSettings indicates an expected call of FindWorkspaceSettings.
func (mr *MockDAOMockRecorder) FindWorkspaceSettings(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWorkspaceSettings", reflect.TypeOf((*MockDAO)(nil).FindWorkspaceSettings), ctx, req)
}

// ReadLocalFile mocks base method.
func (m *MockDAO) ReadLocalFile(ctx context.Context, fpath, name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLocalFile", ctx, fpath, name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLocalFile indicates an expected call of ReadLocalFile.
func (mr *MockDAOMockRecorder) ReadLocalFile(ctx, fpath, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLocalFile", reflect.TypeOf((*MockDAO)(nil).ReadLocalFile), ctx, fpath, name)
}

// RemoveLocalFile mocks base method.
func (m *MockDAO) RemoveLocalFile(ctx context.Context, fpath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLocalFile", ctx, fpath)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLocalFile indicates an expected call of RemoveLocalFile.
func (mr *MockDAOMockRecorder) RemoveLocalFile(ctx, fpath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLocalFile", reflect.TypeOf((*MockDAO)(nil).RemoveLocalFile), ctx, fpath)
}

// SaveLocalFile mocks base method.
func (m *MockDAO) SaveLocalFile(ctx context.Context, fpath string, blob []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLocalFile", ctx, fpath, blob)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLocalFile indicates an expected call of SaveLocalFile.
func (mr *MockDAOMockRecorder) SaveLocalFile(ctx, fpath, blob interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLocalFile", reflect.TypeOf((*MockDAO)(nil).SaveLocalFile), ctx, fpath, blob)
}

// UpdateMemo mocks base method.
func (m *MockDAO) UpdateMemo(ctx context.Context, memo *model.Memo, update map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMemo", ctx, memo, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMemo indicates an expected call of UpdateMemo.
func (mr *MockDAOMockRecorder) UpdateMemo(ctx, memo, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMemo", reflect.TypeOf((*MockDAO)(nil).UpdateMemo), ctx, memo, update)
}

// UpdateResource mocks base method.
func (m_2 *MockDAO) UpdateResource(ctx context.Context, m *model.Resource, update map[string]any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateResource", ctx, m, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResource indicates an expected call of UpdateResource.
func (mr *MockDAOMockRecorder) UpdateResource(ctx, m, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockDAO)(nil).UpdateResource), ctx, m, update)
}

// UpsertMemoRelation mocks base method.
func (m_2 *MockDAO) UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpsertMemoRelation", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMemoRelation indicates an expected call of UpsertMemoRelation.
func (mr *MockDAOMockRecorder) UpsertMemoRelation(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMemoRelation", reflect.TypeOf((*MockDAO)(nil).UpsertMemoRelation), ctx, m)
}

// UpsertWorkspaceSetting mocks base method.
func (m *MockDAO) UpsertWorkspaceSetting(ctx context.Context, setting *model.WorkspaceSetting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkspaceSetting", ctx, setting)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkspaceSetting indicates an expected call of UpsertWorkspaceSetting.
func (mr *MockDAOMockRecorder) UpsertWorkspaceSetting(ctx, setting interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkspaceSetting", reflect.TypeOf((*MockDAO)(nil).UpsertWorkspaceSetting), ctx, setting)
}
