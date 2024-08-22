package service

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) ListMemoRelations(ctx context.Context, req *model.ListMemoRelationsRequest) (list []*model.MemoRelation, err error) {
	list, err = s.dao.FindMemoRelations(ctx, &model.FindMemoRelationsRequest{MemoId: req.Id})
	if err != nil {
		return nil, err
	}
	tempList, err := s.dao.FindMemoRelations(ctx, &model.FindMemoRelationsRequest{
		RelatedMemoId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	list = append(list, tempList...)
	return
}
