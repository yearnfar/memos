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

func GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.Memo, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.GetMemo 失败，服务未注册")
	}
	v1, v2 := defaultService.GetMemo(ctx, req)
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

func DeleteMemo(ctx context.Context, req *model.DeleteMemoRequest) error {
	if defaultService == nil {
		panic("调用模块方法: memo.DeleteMemo 失败，服务未注册")
	}
	v1 := defaultService.DeleteMemo(ctx, req)
	return v1
}

func UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (*model.Memo, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.UpdateMemo 失败，服务未注册")
	}
	v1, v2 := defaultService.UpdateMemo(ctx, req)
	return v1, v2
}

func CreateMemoComment(ctx context.Context, req *model.CreateMemoCommentRequest) (*model.Memo, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.CreateMemoComment 失败，服务未注册")
	}
	v1, v2 := defaultService.CreateMemoComment(ctx, req)
	return v1, v2
}

func CreateResource(ctx context.Context, req *model.CreateResourceRequest) (*model.Resource, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.CreateResource 失败，服务未注册")
	}
	v1, v2 := defaultService.CreateResource(ctx, req)
	return v1, v2
}

func GetResource(ctx context.Context, req *model.GetResourceRequest) (*model.Resource, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.GetResource 失败，服务未注册")
	}
	v1, v2 := defaultService.GetResource(ctx, req)
	return v1, v2
}

func ListResources(ctx context.Context, req *model.ListResourcesRequest) ([]*model.Resource, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListResources 失败，服务未注册")
	}
	v1, v2 := defaultService.ListResources(ctx, req)
	return v1, v2
}

func DeleteResource(ctx context.Context, req *model.DeleteResourceRequest) error {
	if defaultService == nil {
		panic("调用模块方法: memo.DeleteResource 失败，服务未注册")
	}
	v1 := defaultService.DeleteResource(ctx, req)
	return v1
}

func GetResourceBinary(ctx context.Context, req *model.GetResourceBinaryRequest) (rb *model.ResourceBinary, err error) {
	if defaultService == nil {
		panic("调用模块方法: memo.GetResourceBinary 失败，服务未注册")
	}
	v1, v2 := defaultService.GetResourceBinary(ctx, req)
	return v1, v2
}

func SetMemoResources(ctx context.Context, req *model.SetMemoResourcesRequest) error {
	if defaultService == nil {
		panic("调用模块方法: memo.SetMemoResources 失败，服务未注册")
	}
	v1 := defaultService.SetMemoResources(ctx, req)
	return v1
}

func ListMemoRelations(ctx context.Context, req *model.ListMemoRelationsRequest) ([]*model.MemoRelation, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListMemoRelations 失败，服务未注册")
	}
	v1, v2 := defaultService.ListMemoRelations(ctx, req)
	return v1, v2
}

func SetMemoRelations(ctx context.Context, req *model.SetMemoRelationsRequest) error {
	if defaultService == nil {
		panic("调用模块方法: memo.SetMemoRelations 失败，服务未注册")
	}
	v1 := defaultService.SetMemoRelations(ctx, req)
	return v1
}

func ListReactions(ctx context.Context, req *model.ListReactionsRequest) ([]*model.Reaction, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.ListReactions 失败，服务未注册")
	}
	v1, v2 := defaultService.ListReactions(ctx, req)
	return v1, v2
}

func UpsertReaction(ctx context.Context, req *model.UpsertReactionRequest) (*model.Reaction, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.UpsertReaction 失败，服务未注册")
	}
	v1, v2 := defaultService.UpsertReaction(ctx, req)
	return v1, v2
}

func SetWorkspaceSetting(ctx context.Context, req *model.SetWorkspaceSettingRequest) (*model.WorkspaceSettingCache, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.SetWorkspaceSetting 失败，服务未注册")
	}
	v1, v2 := defaultService.SetWorkspaceSetting(ctx, req)
	return v1, v2
}

func GetWorkspaceSetting(ctx context.Context, req *model.GetWorkspaceSettingRequest) (*model.WorkspaceSettingCache, error) {
	if defaultService == nil {
		panic("调用模块方法: memo.GetWorkspaceSetting 失败，服务未注册")
	}
	v1, v2 := defaultService.GetWorkspaceSetting(ctx, req)
	return v1, v2
}
