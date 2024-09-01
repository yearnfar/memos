package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (dao *Dao) CreateMemo(ctx context.Context, m *model.Memo) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindMemos(ctx context.Context, req *model.FindMemoRequest) (list []*model.MemoInfo, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn = conn.Where("id=?", req.Id)
	}
	if req.UID != "" {
		conn = conn.Where("uid=?", req.UID)
	}
	if req.CreatorId != 0 {
		conn = conn.Where("creator_id=?", req.CreatorId)
	}
	if req.ExcludeComments {
		conn = conn.Where("parent_id IS NULL")
	}

	err = conn.
		Select(`m.*, related_memo_id AS parent_id, IFNULL(mo.pinned, 0) AS pinned`).
		Table("memo m").
		Joins("LEFT JOIN memo_organizer mo ON m.id = mo.memo_id AND m.creator_id = mo.user_id").
		Joins("LEFT JOIN memo_relation mr ON m.id = mr.memo_id AND mr.type = ?", model.MemoRelationComment).
		Find(&list).Error
	return
}

func (dao *Dao) FindMemo(ctx context.Context, req *model.FindMemoRequest) (*model.MemoInfo, error) {
	list, err := dao.FindMemos(ctx, req)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return list[0], nil
}

func (dao *Dao) FindMemoById(ctx context.Context, id int32) (*model.Memo, error) {
	var m model.Memo
	if err := db.GetDB(ctx).Where("id=?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) UpdateMemo(ctx context.Context, m *model.Memo, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(m).Updates(update).Error
	return
}

func (dao *Dao) DeleteMemoById(ctx context.Context, id int32) (err error) {
	err = db.GetDB(ctx).Model(&model.Memo{}).Delete("id=?", id).Error
	return
}
