package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) UpsertMemoOrganizer(ctx context.Context, m *model.MemoOrganizer) (err error) {
	err = db.GetDB(ctx).
		Where("pinned=?", m.Pinned).
		Assign(model.MemoOrganizer{MemoID: m.MemoID, UserID: m.UserID}).
		FirstOrCreate(&m).Error
	return
}

func (dao *Dao) FindMemoOrganizers(ctx context.Context, req *model.FindMemoOrganizersRequest) (list []*model.MemoOrganizer, err error) {
	conn := db.GetDB(ctx)
	if req.MemoID != 0 {
		conn = conn.Where("memo_id=?", req.MemoID)
	}
	if req.UserID != 0 {
		conn = conn.Where("user_id=?", req.UserID)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) DeleteMemoOrganizer(ctx context.Context, req *model.DeleteMemoOrganizersRequest) error {
	conn := db.GetDB(ctx)
	if req.MemoID != 0 {
		conn = conn.Where("memo_id=?", req.MemoID)
	}
	if req.UserID != 0 {
		conn = conn.Where("user_id=?", req.UserID)
	}
	err := conn.Delete(&model.MemoOrganizer{}).Error
	return err
}
