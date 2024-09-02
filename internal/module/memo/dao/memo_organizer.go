package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) UpsertMemoOrganizer(ctx context.Context, m *model.MemoOrganizer) (err error) {
	err = db.GetDB(ctx).
		Where("memo_id=? AND user_id=?", m.MemoID, m.UserID).
		Assign("pinned", m.Pinned).
		FirstOrCreate(&m).Error
	return
}

func (dao *Dao) FindMemoOrganizers(ctx context.Context, where []string, args []any, fields ...string) (list []*model.MemoOrganizer, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindMemoOrganizer(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoOrganizer, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.MemoOrganizer
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
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
