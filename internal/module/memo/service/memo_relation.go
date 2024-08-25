package service

import (
	"context"

	"github.com/pkg/errors"
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

func (s *Service) SetMemoRelations(ctx context.Context, req *model.SetMemoRelationsRequest) (err error) {
	referenceType := model.MemoRelationReference
	if err = s.dao.DeleteMemoRelations(ctx, &model.DeleteMemoRelationsRequest{
		MemoID: req.MemoID,
		Type:   referenceType,
	}); err != nil {
		err = errors.New("failed to delete memo relation")
		return
	}

	for _, relation := range req.Relations {
		if req.MemoID == relation.RelatedMemoID {
			continue
		}
		if relation.Type == model.MemoRelationComment {
			continue
		}
		if err = s.dao.UpsertMemoRelation(ctx, &model.MemoRelation{
			MemoID:        req.MemoID,
			RelatedMemoID: relation.RelatedMemoID,
			Type:          relation.Type,
		}); err != nil {
			err = errors.New("failed to upsert memo relation")
			return
		}
	}
	return
}
