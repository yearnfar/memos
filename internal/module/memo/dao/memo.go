package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateMemo(ctx context.Context, m *model.Memo) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindMemos(ctx context.Context, req *model.FindMemosRequest) (list []*model.Memo, err error) {
	conn := db.GetDB(ctx)
	if req.CreatorId != 0 {
		conn = conn.Where("creator_id=?", req.CreatorId)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) FindMemo(ctx context.Context, req *model.FindMemoRequest) (memo *model.Memo, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn = conn.Where("id=?", req.Id)
	}
	memo = &model.Memo{}
	err = conn.First(&memo).Error
	return
}

func (dao *Dao) UpdateMemo(ctx context.Context, memo *model.Memo, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(memo).Updates(update).Error
	return
}

func (dao *Dao) DeleteMemoById(ctx context.Context, id int32) (err error) {
	err = db.GetDB(ctx).Model(&model.Memo{}).Delete("id=?", id).Error
	return
}
