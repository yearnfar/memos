package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateMemo(ctx context.Context, m *model.Memo) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindMemos(ctx context.Context, req *model.FindMemosRequest) (memos []*model.Memo, err error) {
	return
}
