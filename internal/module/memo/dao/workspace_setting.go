package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) UpsertWorkspaceSetting(ctx context.Context, m *model.WorkspaceSetting) (err error) {
	err = db.GetDB(ctx).
		Where("name=?", m.Name).
		Assign(model.WorkspaceSetting{Value: m.Value, Description: m.Description}).
		FirstOrCreate(&m).Error
	return
}

func (dao *Dao) FindWorkspaceSettings(ctx context.Context, where []string, args []any, fields ...string) (list []*model.WorkspaceSetting, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindWorkspaceSetting(ctx context.Context, where []string, args []any, fields ...string) (*model.WorkspaceSetting, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.WorkspaceSetting
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) DeleteWorkspaceSetting(ctx context.Context, name string) error {
	err := db.GetDB(ctx).Model(&model.WorkspaceSetting{}).Delete("name=?", name).Error
	return err
}
