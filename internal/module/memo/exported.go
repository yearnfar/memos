// Code generated by service-export(1.0.0). DO NOT EDIT.
// source: service.go

package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

var defaultService Service

func Register(s Service) {
	defaultService = s
}

func ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListInboxes 失败，服务未注册")
	}
	v1, v2 := defaultService.ListInboxes(ctx, req)
	return v1, v2
}

func ListMemos(ctx context.Context, req *model.ListMemosRequest) ([]*model.Memo, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListMemos 失败，服务未注册")
	}
	v1, v2 := defaultService.ListMemos(ctx, req)
	return v1, v2
}

func CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (*model.Memo, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.CreateMemo 失败，服务未注册")
	}
	v1, v2 := defaultService.CreateMemo(ctx, req)
	return v1, v2
}

func ListMemoRelations(ctx context.Context, req *model.ListMemoRelationsRequest) ([]*model.MemoRelation, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListMemoRelations 失败，服务未注册")
	}
	v1, v2 := defaultService.ListMemoRelations(ctx, req)
	return v1, v2
}

func ListReactions(ctx context.Context, req *model.ListReactionRequest) ([]*model.Reaction, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListReactions 失败，服务未注册")
	}
	v1, v2 := defaultService.ListReactions(ctx, req)
	return v1, v2
}
