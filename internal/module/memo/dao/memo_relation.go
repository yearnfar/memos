package dao

import (
	"context"
	"errors"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) (list []*model.MemoRelation, err error) {
	conn := db.GetDB(ctx)
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}

func (dao *Dao) DeleteMemoRelations(ctx context.Context, req *model.DeleteMemoRelationsRequest) (err error) {
	if req.MemoID == 0 && req.Type == "" && req.RelatedMemoID == 0 {
		err = errors.New("no condition")
		return
	}
	conn := db.GetDB(ctx)
	if req.MemoID != 0 {
		conn.Where("memo_id=?", req.MemoID)
	}
	if req.Type != "" {
		conn.Where("type=?", req.Type)
	}
	if req.RelatedMemoID != 0 {
		conn.Where("related_memo_id=?", req.RelatedMemoID)
	}
	err = conn.Delete(&model.MemoRelation{}).Error
	return
}
