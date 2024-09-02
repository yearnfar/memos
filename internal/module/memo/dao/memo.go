package dao

import (
	"context"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) CreateMemo(ctx context.Context, m *model.Memo) error {
	return db.GetDB(ctx).Create(m).Error
}

func (dao *Dao) FindMemos(ctx context.Context, where []string, args []any, fields ...string) (list []*model.MemoInfo, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	if len(fields) == 0 {
		fields = []string{`m.*, related_memo_id AS parent_id, IFNULL(mo.pinned, 0) AS pinned`}
	}
	err = db.GetDB(ctx).
		Select(fields[0]).
		Table("memo m").
		Joins("LEFT JOIN memo_organizer mo ON m.id = mo.memo_id AND m.creator_id = mo.user_id").
		Joins("LEFT JOIN memo_relation mr ON m.id = mr.memo_id AND mr.type = ?", model.MemoRelationComment).
		Where(strings.Join(where, " and "), args...).
		Find(&list).Error
	return
}

func (dao *Dao) FindMemo(ctx context.Context, where []string, args []any, fields ...string) (*model.MemoInfo, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var m model.MemoInfo
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *Dao) FindMemoByID(ctx context.Context, id int32, fields ...string) (*model.MemoInfo, error) {
	if len(fields) == 0 {
		fields = []string{`m.*, related_memo_id AS parent_id, IFNULL(mo.pinned, 0) AS pinned`}
	}
	var m model.MemoInfo
	if err := db.GetDB(ctx).
		Select(fields[0]).
		Table("memo m").
		Joins("LEFT JOIN memo_organizer mo ON m.id = mo.memo_id AND m.creator_id = mo.user_id").
		Joins("LEFT JOIN memo_relation mr ON m.id = mr.memo_id AND mr.type = ?", model.MemoRelationComment).
		Where("m.id=?", id).First(&m).Error; err != nil {
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
