package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (dao *Dao) FindReactions(ctx context.Context, req *model.FindReactionsRequest) (list []*model.Reaction, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn = conn.Where("id=?", req.Id)
	}
	if req.CreatorId != 0 {
		conn = conn.Where("creator_id=?", req.CreatorId)
	}
	if req.ContentId != "" {
		conn = conn.Where("content_id=?", req.ContentId)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) FindReaction(ctx context.Context, req *model.FindReactionsRequest) (*model.Reaction, error) {
	list, err := dao.FindReactions(ctx, req)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (dao *Dao) CreateReaction(ctx context.Context, m *model.Reaction) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}

func (dao *Dao) DeleteReaction(ctx context.Context, id int32) error {
	err := db.GetDB(ctx).Model(&model.Reaction{}).Delete("id=?", id).Error
	return err
}
