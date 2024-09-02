package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateWebhook(ctx context.Context, m *model.Webhook) error {
	err := db.GetDB(ctx).Create(m).Error
	return err
}

func (dao *Dao) FindWebhooks(ctx context.Context, where []string, args []any, fields ...string) (list []*model.Webhook, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) FindWebhook(ctx context.Context, where []string, args []any, fields ...string) (*model.Webhook, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.Webhook
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) UpdateWebhook(ctx context.Context, m *model.Webhook, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(m).Updates(update).Error
	return
}

func (dao *Dao) DeleteWebhookById(ctx context.Context, id int32) error {
	err := db.GetDB(ctx).Model(&model.Webhook{}).Delete("id=?", id).Error
	return err
}
