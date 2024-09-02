package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error) {
	return s.dao.FindInboxes(ctx, []string{"receiver_id=?"}, []any{req.ReceiverId})
}
