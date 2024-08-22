package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListReactions(ctx context.Context, req *model.ListReactionsRequest) (list []*model.Reaction, err error) {
	return s.dao.FindReactions(ctx, &model.FindReactionsRequest{
		Id:        req.Id,
		CreatorId: req.CreatorId,
		ContentId: req.ContentId,
	})
}
