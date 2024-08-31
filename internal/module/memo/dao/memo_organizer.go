package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

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
