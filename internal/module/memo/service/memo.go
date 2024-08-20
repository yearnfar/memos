package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (list []*model.Memo, err error) {
	return s.dao.FindMemos(ctx, &model.FindMemosRequest{})
}
