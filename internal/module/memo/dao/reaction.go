package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindReactions(ctx context.Context, where []string, args []any, fields ...string) (list []*model.Reaction, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindReaction(ctx context.Context, where []string, args []any, fields ...string) (*model.Reaction, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.Reaction
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) CreateReaction(ctx context.Context, m *model.Reaction) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}

func (dao *Dao) DeleteReaction(ctx context.Context, id int32) error {
	err := db.GetDB(ctx).Model(&model.Reaction{}).Delete("id=?", id).Error
	return err
}
