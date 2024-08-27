package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindResource(ctx context.Context, req *model.FindResourceRequest) (*model.Resource, error) {
	conn := db.GetDB(ctx)
	if req.ID != 0 {
		conn = conn.Where("id=?", req.ID)
	}
	var m model.Resource
	if err := conn.First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) FindResources(ctx context.Context, req *model.FindResourcesRequest) (list []*model.Resource, err error) {
	conn := db.GetDB(ctx)
	if req.ID != 0 {
		conn = conn.Where("id=?", req.ID)
	}
	if req.MemoID != 0 {
		conn = conn.Where("memo_id", req.MemoID)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) DeleteResourceById(ctx context.Context, id int32) (err error) {
	err = db.GetDB(ctx).Model(&model.Resource{}).Delete("id=?", id).Error
	return
}

func (dao *Dao) UpdateResource(ctx context.Context, m *model.Resource, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(m).Updates(update).Error
	return
}
