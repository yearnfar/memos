package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (dao *Dao) FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) ([]*model.MemoRelation, error) {
	return nil, nil
}
