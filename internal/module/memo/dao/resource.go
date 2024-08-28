package dao

import (
	"context"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
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

func (dao *Dao) CreateResource(ctx context.Context, m *model.Resource) (err error) {
	err = db.GetDB(ctx).Create(m).Error
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

func (dao *Dao) SaveLocalFile(ctx context.Context, savePath string, blob []byte) (err error) {
	dir := filepath.Dir(savePath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		err = errors.Wrap(err, "Failed to create directory")
		return
	}
	dst, err := os.Create(savePath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file")
		return
	}
	defer dst.Close()
	if err = os.WriteFile(savePath, blob, 0644); err != nil {
		err = errors.Wrap(err, "Failed to write file")
		return
	}
	return
}
