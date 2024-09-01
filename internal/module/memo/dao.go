//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type DAO interface {
	FindInboxes(ctx context.Context, req *model.FindInboxRequest) ([]*model.Inbox, error)
	CreateInbox(ctx context.Context, inbox *model.Inbox) error

	CreateActivity(ctx context.Context, memo *model.Activity) error

	CreateMemo(ctx context.Context, memo *model.Memo) error
	FindMemos(ctx context.Context, req *model.FindMemoRequest) ([]*model.MemoInfo, error)
	FindMemo(ctx context.Context, req *model.FindMemoRequest) (*model.MemoInfo, error)
	UpdateMemo(ctx context.Context, memo *model.Memo, update map[string]any) error
	DeleteMemoById(ctx context.Context, id int32) error

	FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) ([]*model.MemoRelation, error)
	DeleteMemoRelations(ctx context.Context, req *model.DeleteMemoRelationsRequest) error
	UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error

	FindMemoOrganizers(ctx context.Context, req *model.FindMemoOrganizersRequest) ([]*model.MemoOrganizer, error)
	UpsertMemoOrganizer(ctx context.Context, m *model.MemoOrganizer) error

	FindReactions(ctx context.Context, req *model.FindReactionsRequest) ([]*model.Reaction, error)
	CreateReaction(ctx context.Context, m *model.Reaction) error

	CreateResource(ctx context.Context, m *model.Resource) error
	FindResource(ctx context.Context, req *model.FindResourceRequest) (*model.Resource, error)
	FindResources(ctx context.Context, req *model.FindResourceRequest) ([]*model.Resource, error)
	DeleteResourceById(ctx context.Context, id int32) error
	UpdateResource(ctx context.Context, m *model.Resource, update map[string]any) error

	SaveLocalFile(ctx context.Context, fpath string, blob []byte) error
	ReadLocalFile(ctx context.Context, fpath, name string) ([]byte, error)
	RemoveLocalFile(ctx context.Context, fpath string) error

	UpsertWorkspaceSetting(ctx context.Context, setting *model.WorkspaceSetting) error
	FindWorkspaceSettings(ctx context.Context, req *model.FindWorkspaceSettingsRequest) ([]*model.WorkspaceSetting, error)
}
