package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateWebhook(ctx context.Context, m *model.Webhook) error {
	err := db.GetDB(ctx).Create(m).Error
	return err
}

func (dao *Dao) FindWebhooks(ctx context.Context, req *model.FindWebhookRequest) (list []*model.Webhook, err error) {
	conn := db.GetDB(ctx)
	if req.ID != 0 {
		conn = conn.Where("id=?", req.ID)
	}
	if req.CreatorID != 0 {
		conn = conn.Where("creator_id=?", req.CreatorID)
	}
	err = conn.Find(&list).Error
	return list, err
}

func (dao *Dao) FindWebhook(ctx context.Context, find *model.FindWebhookRequest) (*model.Webhook, error) {
	list, err := dao.FindWebhooks(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (dao *Dao) UpdateWebhook(ctx context.Context, m *model.Webhook, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(m).Updates(update).Error
	return
}

func (dao *Dao) DeleteWebhookById(ctx context.Context, id int32) error {
	err := db.GetDB(ctx).Model(&model.Webhook{}).Delete("id=?", id).Error
	return err
}
