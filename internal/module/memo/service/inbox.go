package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListInboxes(ctx context.Context, req *model.ListInboxesRequest) ([]*model.Inbox, error) {
	return s.dao.FindInboxes(ctx, &model.FindInboxRequest{ReceiverId: req.ReceiverId})
}
