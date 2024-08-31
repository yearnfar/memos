package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateActivity(ctx context.Context, m *model.Activity) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindActivities(ctx context.Context, req *model.FindActivityRequest) (list []*model.Memo, err error) {
	conn := db.GetDB(ctx)
	if req.ID != 0 {
		conn = conn.Where("id=?", req.ID)
	}
	if req.Type != "" {
		conn = conn.Where("type=?", req.Type)
	}
	err = conn.Find(&list).Error
	return
}
