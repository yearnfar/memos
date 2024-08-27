package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindReactions(ctx context.Context, req *model.FindReactionsRequest) (list []*model.Reaction, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn.Where("id=?", req.Id)
	}
	if req.CreatorId != 0 {
		conn.Where("creator_id=?", req.CreatorId)
	}
	if req.ContentId != "" {
		conn.Where("content_id=?", req.ContentId)
	}
	err = conn.Find(&list).Error
	return
}
func (dao *Dao) CreateReaction(ctx context.Context, m *model.Reaction) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}
