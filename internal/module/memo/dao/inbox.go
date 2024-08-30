package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindInboxes(ctx context.Context, req *model.FindInboxesRequest) (inboxes []*model.Inbox, err error) {
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
	inboxes = []*model.Inbox{}
	err = conn.Find(&inboxes).Error
	return
}

func (dao *Dao) CreateInbox(ctx context.Context, m *model.Inbox) error {
	return db.GetDB(ctx).Create(m).Error
}
