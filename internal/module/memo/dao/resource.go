package dao

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/yearnfar/gokit/fsutil"
	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateResource(ctx context.Context, m *model.Resource) (err error) {
	err = db.GetDB(ctx).Create(m).Error
	return
}

func (dao *Dao) FindResources(ctx context.Context, where []string, args []any, fields ...string) (list []*model.Resource, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindResource(ctx context.Context, where []string, args []any, fields ...string) (*model.Resource, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.Resource
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) FindResourceByID(ctx context.Context, id int32, fields ...string) (*model.Resource, error) {
	var m model.Resource
	if err := db.GetDB(ctx).Where("id=?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
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
