//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type DAO interface {
	CreateInbox(ctx context.Context, inbox *model.Inbox) error
	FindInboxes(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Inbox, error)
	FindInbox(ctx context.Context, where []string, args []any, fields ...string) (*model.Inbox, error)

	CreateActivity(ctx context.Context, memo *model.Activity) error
	FindActivities(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Activity, error)
	FindActivity(ctx context.Context, where []string, args []any, fields ...string) (*model.Activity, error)

	CreateMemo(ctx context.Context, memo *model.Memo) error
	FindMemos(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoInfo, error)
	FindMemo(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoInfo, error)
	FindMemoByID(ctx context.Context, id int32, fields ...string) (*model.MemoInfo, error)
	UpdateMemo(ctx context.Context, memo *model.Memo, update map[string]any) error
	DeleteMemoById(ctx context.Context, id int32) error

	UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) error
	FindMemoRelations(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoRelation, error)
	FindMemoRelation(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoRelation, error)
	DeleteMemoRelations(ctx context.Context, where []string, args []any) error

	FindMemoOrganizers(ctx context.Context, where []string, args []any, fields ...string) ([]*model.MemoOrganizer, error)
	FindMemoOrganizer(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoOrganizer, error)
	UpsertMemoOrganizer(ctx context.Context, m *model.MemoOrganizer) error

	CreateReaction(ctx context.Context, m *model.Reaction) error
	FindReactions(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Reaction, error)
	FindReaction(ctx context.Context, where []string, args []any, fields ...string) (*model.Reaction, error)

	CreateResource(ctx context.Context, m *model.Resource) error
	FindResourceByID(ctx context.Context, id int32, fields ...string) (*model.Resource, error)
	FindResource(ctx context.Context, where []string, args []any, fields ...string) (*model.Resource, error)
	FindResources(ctx context.Context, where []string, args []any, fields ...string) ([]*model.Resource, error)
	DeleteResourceById(ctx context.Context, id int32) error
	UpdateResource(ctx context.Context, m *model.Resource, update map[string]any) error

	SaveLocalFile(ctx context.Context, fpath string, blob []byte) error
	ReadLocalFile(ctx context.Context, fpath, name string) ([]byte, error)
	RemoveLocalFile(ctx context.Context, fpath string) error

	UpsertWorkspaceSetting(ctx context.Context, setting *model.WorkspaceSetting) error
	FindWorkspaceSettings(ctx context.Context, where []string, args []any, fields ...string) ([]*model.WorkspaceSetting, error)
	FindWorkspaceSetting(ctx context.Context, where []string, args []any, fields ...string) (*model.WorkspaceSetting, error)
}
