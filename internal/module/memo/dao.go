//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type DAO interface {
	FindInboxes(ctx context.Context, req *model.FindInboxesRequest) ([]*model.Inbox, error)
	CreateMemo(ctx context.Context, memo *model.Memo) error
	FindMemos(ctx context.Context, req *model.FindMemosRequest) ([]*model.Memo, error)
	FindMemo(ctx context.Context, req *model.FindMemoRequest) (*model.Memo, error)
	DeleteMemoById(ctx context.Context, id int32) error

	FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) ([]*model.MemoRelation, error)
	DeleteMemoRelations(ctx context.Context, req *model.DeleteMemoRelationsRequest) error
	UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error

	FindMemoOrganizers(ctx context.Context, req *model.FindMemoOrganizersRequest) ([]*model.MemoOrganizer, error)
	FindReactions(ctx context.Context, req *model.FindReactionsRequest) ([]*model.Reaction, error)
	CreateReaction(ctx context.Context, m *model.Reaction) error

	CreateResource(ctx context.Context, m *model.Resource) error
	FindResource(ctx context.Context, req *model.FindResourceRequest) (*model.Resource, error)
	FindResources(ctx context.Context, req *model.FindResourcesRequest) ([]*model.Resource, error)
	DeleteResourceById(ctx context.Context, id int32) error
	UpdateResource(ctx context.Context, m *model.Resource, update map[string]any) error

	SaveLocalFile(ctx context.Context, savePath string, blob []byte) error

	UpsertWorkspaceSetting(ctx context.Context, setting *model.WorkspaceSetting) error
	FindWorkspaceSettings(ctx context.Context, req *model.FindWorkspaceSettingsRequest) ([]*model.WorkspaceSetting, error)
}
