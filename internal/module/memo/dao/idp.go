package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateIdentityProvider(ctx context.Context, m *model.IdentityProvider) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindIdentityProviders(ctx context.Context, where []string, args []any, fields ...string) (list []*model.IdentityProvider, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindIdentityProvider(ctx context.Context, where []string, args []any, fields ...string) (*model.IdentityProvider, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.IdentityProvider
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) UpdateIdentityProvider(ctx context.Context, m *model.IdentityProvider, update map[string]any) error {
	return db.GetDB(ctx).Model(m).Updates(update).Error
}

func (dao *Dao) DeleteIdentityProviderById(ctx context.Context, id int32) error {
	return db.GetDB(ctx).Model(&model.Memo{}).Delete("id=?", id).Error
}
