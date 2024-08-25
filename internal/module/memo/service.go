//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type Service interface {
	ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error)
	GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.Memo, error)
	ListMemos(ctx context.Context, req *model.ListMemosRequest) ([]*model.Memo, error)
	CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (*model.Memo, error)

	SetMemoResources(ctx context.Context, req *model.SetMemoResourcesRequest) error

	ListMemoRelations(ctx context.Context, req *model.ListMemoRelationsRequest) ([]*model.MemoRelation, error)

	ListReactions(ctx context.Context, req *model.ListReactionsRequest) ([]*model.Reaction, error)

	ListResources(ctx context.Context, req *model.ListResourcesRequest) ([]*model.Resource, error)

	SetWorkspaceSetting(ctx context.Context, req *model.SetWorkspaceSettingRequest) (*model.WorkspaceSettingCache, error)
	GetWorkspaceSetting(ctx context.Context, req *model.GetWorkspaceSettingRequest) (*model.WorkspaceSettingCache, error)
}
