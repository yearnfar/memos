package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListReactions(ctx context.Context, req *model.ListReactionsRequest) (list []*model.Reaction, err error) {
	return s.dao.FindReactions(ctx, []string{"id", "creator_id", "content_id"}, []any{req.Id, req.CreatorId, req.ContentId})
}
