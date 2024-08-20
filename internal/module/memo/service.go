//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memo

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

type Service interface {
	ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error)
}
