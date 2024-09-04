package dao

import (
	"context"
	"errors"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindMemoRelations(ctx context.Context, where []string, args []any, fields ...string) (list []*model.MemoRelation, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindMemoRelation(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoRelation, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.MemoRelation
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}

func (dao *Dao) DeleteMemoRelations(ctx context.Context, where []string, args []any) (err error) {
	if len(where) == 0 {
		return errors.New("where can not be empty")
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Delete(&model.MemoRelation{}).Error
	return
}
