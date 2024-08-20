//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type DAO interface {
	FindInboxes(ctx context.Context, req *model.FindInboxesRequest) ([]*model.Inbox, error)
	FindMemos(ctx context.Context, req *model.FindMemosRequest) ([]*model.Memo, error)
}
