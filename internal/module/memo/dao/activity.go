package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateActivity(ctx context.Context, m *model.Activity) error {
	return db.GetDB(ctx).Create(m).Error
}
