package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (dao *Dao) CreateInbox(ctx context.Context, m *model.Inbox) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindInboxes(ctx context.Context, req *model.FindInboxRequest) (list []*model.Inbox, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn.Where("id=?", req.Id)
	}
	if req.SenderId != 0 {
		conn.Where("sender_id=?", req.SenderId)
	}
	if req.ReceiverId != 0 {
		conn.Where("receiver_id=?", req.ReceiverId)
	}
	if req.Status != "" {
		conn.Where("status=?", req.Status)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) FindInbox(ctx context.Context, req *model.FindInboxRequest) (*model.Inbox, error) {
	list, err := dao.FindInboxes(ctx, req)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (dao *Dao) UpdateInbox(ctx context.Context, m *model.Inbox, update map[string]any) error {
	return db.GetDB(ctx).Model(m).Updates(update).Error
}

func (dao *Dao) DeleteInboxById(ctx context.Context, id int32) error {
	return db.GetDB(ctx).Model(&model.Memo{}).Delete("id=?", id).Error
}
