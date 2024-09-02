package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateActivity(ctx context.Context, m *model.Activity) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindActivities(ctx context.Context, where []string, args []any, fields ...string) (list []*model.Activity, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindActivity(ctx context.Context, where []string, args []any, fields ...string) (*model.Activity, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.Activity
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
