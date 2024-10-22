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

// CreateActivity mocks base method.
func (m *MockDAO) CreateActivity(ctx context.Context, memo *model.Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActivity", ctx, memo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockDAOMockRecorder) CreateActivity(ctx, memo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockDAO)(nil).CreateActivity), ctx, memo)
}

// CreateInbox mocks base method.
func (m *MockDAO) CreateInbox(ctx context.Context, inbox *model.Inbox) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInbox", ctx, inbox)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInbox indicates an expected call of CreateInbox.
func (mr *MockDAOMockRecorder) CreateInbox(ctx, inbox interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInbox", reflect.TypeOf((*MockDAO)(nil).CreateInbox), ctx, inbox)
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
func (m *MockDAO) DeleteMemoRelations(ctx context.Context, where []string, args []any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMemoRelations", ctx, where, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMemoRelations indicates an expected call of DeleteMemoRelations.
func (mr *MockDAOMockRecorder) DeleteMemoRelations(ctx, where, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMemoRelations", reflect.TypeOf((*MockDAO)(nil).DeleteMemoRelations), ctx, where, args)
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

// FindActivities mocks base method.
func (m *MockDAO) FindActivities(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Activity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindActivities", varargs...)
	ret0, _ := ret[0].([]*model.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActivities indicates an expected call of FindActivities.
func (mr *MockDAOMockRecorder) FindActivities(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActivities", reflect.TypeOf((*MockDAO)(nil).FindActivities), varargs...)
}

// FindActivity mocks base method.
func (m *MockDAO) FindActivity(ctx context.Context, where []string, args []any, fields ...string) (*model.Activity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindActivity", varargs...)
	ret0, _ := ret[0].(*model.Activity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActivity indicates an expected call of FindActivity.
func (mr *MockDAOMockRecorder) FindActivity(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActivity", reflect.TypeOf((*MockDAO)(nil).FindActivity), varargs...)
}

// FindInbox mocks base method.
func (m *MockDAO) FindInbox(ctx context.Context, where []string, args []any, fields ...string) (*model.Inbox, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindInbox", varargs...)
	ret0, _ := ret[0].(*model.Inbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInbox indicates an expected call of FindInbox.
func (mr *MockDAOMockRecorder) FindInbox(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInbox", reflect.TypeOf((*MockDAO)(nil).FindInbox), varargs...)
}

// FindInboxes mocks base method.
func (m *MockDAO) FindInboxes(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Inbox, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindInboxes", varargs...)
	ret0, _ := ret[0].([]*model.Inbox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInboxes indicates an expected call of FindInboxes.
func (mr *MockDAOMockRecorder) FindInboxes(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInboxes", reflect.TypeOf((*MockDAO)(nil).FindInboxes), varargs...)
}

// FindMemo mocks base method.
func (m *MockDAO) FindMemo(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemo", varargs...)
	ret0, _ := ret[0].(*model.MemoInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemo indicates an expected call of FindMemo.
func (mr *MockDAOMockRecorder) FindMemo(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemo", reflect.TypeOf((*MockDAO)(nil).FindMemo), varargs...)
}

// FindMemoByID mocks base method.
func (m *MockDAO) FindMemoByID(ctx context.Context, id int32, fields ...string) (*model.MemoInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemoByID", varargs...)
	ret0, _ := ret[0].(*model.MemoInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoByID indicates an expected call of FindMemoByID.
func (mr *MockDAOMockRecorder) FindMemoByID(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoByID", reflect.TypeOf((*MockDAO)(nil).FindMemoByID), varargs...)
}

// FindMemoOrganizer mocks base method.
func (m *MockDAO) FindMemoOrganizer(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoOrganizer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemoOrganizer", varargs...)
	ret0, _ := ret[0].(*model.MemoOrganizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoOrganizer indicates an expected call of FindMemoOrganizer.
func (mr *MockDAOMockRecorder) FindMemoOrganizer(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoOrganizer", reflect.TypeOf((*MockDAO)(nil).FindMemoOrganizer), varargs...)
}

// FindMemoOrganizers mocks base method.
func (m *MockDAO) FindMemoOrganizers(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoOrganizer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemoOrganizers", varargs...)
	ret0, _ := ret[0].([]*model.MemoOrganizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoOrganizers indicates an expected call of FindMemoOrganizers.
func (mr *MockDAOMockRecorder) FindMemoOrganizers(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoOrganizers", reflect.TypeOf((*MockDAO)(nil).FindMemoOrganizers), varargs...)
}

// FindMemoRelation mocks base method.
func (m *MockDAO) FindMemoRelation(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoRelation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemoRelation", varargs...)
	ret0, _ := ret[0].(*model.MemoRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoRelation indicates an expected call of FindMemoRelation.
func (mr *MockDAOMockRecorder) FindMemoRelation(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoRelation", reflect.TypeOf((*MockDAO)(nil).FindMemoRelation), varargs...)
}

// FindMemoRelations mocks base method.
func (m *MockDAO) FindMemoRelations(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoRelation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemoRelations", varargs...)
	ret0, _ := ret[0].([]*model.MemoRelation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemoRelations indicates an expected call of FindMemoRelations.
func (mr *MockDAOMockRecorder) FindMemoRelations(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemoRelations", reflect.TypeOf((*MockDAO)(nil).FindMemoRelations), varargs...)
}

// FindMemos mocks base method.
func (m *MockDAO) FindMemos(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindMemos", varargs...)
	ret0, _ := ret[0].([]*model.MemoInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemos indicates an expected call of FindMemos.
func (mr *MockDAOMockRecorder) FindMemos(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemos", reflect.TypeOf((*MockDAO)(nil).FindMemos), varargs...)
}

// FindReaction mocks base method.
func (m *MockDAO) FindReaction(ctx context.Context, where []string, args []any, fields ...string) (*model.Reaction, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindReaction", varargs...)
	ret0, _ := ret[0].(*model.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReaction indicates an expected call of FindReaction.
func (mr *MockDAOMockRecorder) FindReaction(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReaction", reflect.TypeOf((*MockDAO)(nil).FindReaction), varargs...)
}

// FindReactions mocks base method.
func (m *MockDAO) FindReactions(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Reaction, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindReactions", varargs...)
	ret0, _ := ret[0].([]*model.Reaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindReactions indicates an expected call of FindReactions.
func (mr *MockDAOMockRecorder) FindReactions(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindReactions", reflect.TypeOf((*MockDAO)(nil).FindReactions), varargs...)
}

// FindResource mocks base method.
func (m *MockDAO) FindResource(ctx context.Context, where []string, args []any, fields ...string) (*model.Resource, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindResource", varargs...)
	ret0, _ := ret[0].(*model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResource indicates an expected call of FindResource.
func (mr *MockDAOMockRecorder) FindResource(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResource", reflect.TypeOf((*MockDAO)(nil).FindResource), varargs...)
}

// FindResourceByID mocks base method.
func (m *MockDAO) FindResourceByID(ctx context.Context, id int32, fields ...string) (*model.Resource, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindResourceByID", varargs...)
	ret0, _ := ret[0].(*model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResourceByID indicates an expected call of FindResourceByID.
func (mr *MockDAOMockRecorder) FindResourceByID(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResourceByID", reflect.TypeOf((*MockDAO)(nil).FindResourceByID), varargs...)
}

// FindResources mocks base method.
func (m *MockDAO) FindResources(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Resource, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindResources", varargs...)
	ret0, _ := ret[0].([]*model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindResources indicates an expected call of FindResources.
func (mr *MockDAOMockRecorder) FindResources(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindResources", reflect.TypeOf((*MockDAO)(nil).FindResources), varargs...)
}

// FindWorkspaceSetting mocks base method.
func (m *MockDAO) FindWorkspaceSetting(ctx context.Context, where []string, args []any, fields ...string) (*model.WorkspaceSetting, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindWorkspaceSetting", varargs...)
	ret0, _ := ret[0].(*model.WorkspaceSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWorkspaceSetting indicates an expected call of FindWorkspaceSetting.
func (mr *MockDAOMockRecorder) FindWorkspaceSetting(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWorkspaceSetting", reflect.TypeOf((*MockDAO)(nil).FindWorkspaceSetting), varargs...)
}

// FindWorkspaceSettings mocks base method.
func (m *MockDAO) FindWorkspaceSettings(ctx context.Context, where []string, args []any, fields ...string) ([]*model.WorkspaceSetting, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, where, args}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindWorkspaceSettings", varargs...)
	ret0, _ := ret[0].([]*model.WorkspaceSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWorkspaceSettings indicates an expected call of FindWorkspaceSettings.
func (mr *MockDAOMockRecorder) FindWorkspaceSettings(ctx, where, args interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, where, args}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWorkspaceSettings", reflect.TypeOf((*MockDAO)(nil).FindWorkspaceSettings), varargs...)
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

// UpsertMemoOrganizer mocks base method.
func (m_2 *MockDAO) UpsertMemoOrganizer(ctx context.Context, m *model.MemoOrganizer) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpsertMemoOrganizer", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMemoOrganizer indicates an expected call of UpsertMemoOrganizer.
func (mr *MockDAOMockRecorder) UpsertMemoOrganizer(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMemoOrganizer", reflect.TypeOf((*MockDAO)(nil).UpsertMemoOrganizer), ctx, m)
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
