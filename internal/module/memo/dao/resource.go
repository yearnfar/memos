package dao

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/yearnfar/gokit/fsutil"
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
		conn = conn.Where("memo_id=?", req.MemoID)
	}
	if req.UID != "" {
		conn = conn.Where("uid=?", req.UID)
	}
	if req.CreatorID != 0 {
		conn = conn.Where("creator_id=?", req.CreatorID)
	}
	if req.Filename != "" {
		conn = conn.Where("file_name=?", req.Filename)
	}
	if req.FilenameSearch != "" {
		conn = conn.Where("filename like ?", "%s"+req.FilenameSearch+"%")
	}
	if req.HasRelatedMemo {
		conn = conn.Where("memo_id is not null")
	}
	if req.StorageType != "" {
		conn = conn.Where("storage_type=?", req.StorageType)
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

func (dao *Dao) SaveLocalFile(ctx context.Context, fpath string, blob []byte) (err error) {
	dir := filepath.Dir(fpath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		err = errors.Wrap(err, "Failed to create directory")
		return
	}
	dst, err := os.Create(fpath)
	if err != nil {
		err = errors.Wrap(err, "Failed to create file")
		return
	}
	defer dst.Close()
	if err = os.WriteFile(fpath, blob, 0644); err != nil {
		err = errors.Wrap(err, "Failed to write file")
		return
	}
	return
}

func (dao *Dao) ReadLocalFile(ctx context.Context, fpath, name string) (blob []byte, err error) {
	file, err := os.Open(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.Errorf("file not found for resource: %s", name)
			return
		}
		errors.Errorf("failed to open the file: %v", err)
		return
	}
	defer file.Close()
	blob, err = io.ReadAll(file)
	if err != nil {
		err = errors.Errorf("failed to read the file: %v", err)
		return
	}
	return
}

func (dao *Dao) RemoveLocalFile(ctx context.Context, fpath string) error {
	if fsutil.IsFile(fpath) {
		return nil
	}
	err := os.Remove(fpath)
	return errors.Wrap(err, "failed to delete local file")
}
