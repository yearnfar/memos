package dao

import (
	"context"
	"errors"
	"strings"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindMemoRelations(ctx context.Context, req *model.FindMemoRelationsRequest) (list []*model.MemoRelation, err error) {
	conn := db.GetDB(ctx)
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) UpsertMemoRelation(ctx context.Context, m *model.MemoRelation) (err error) {
	err = db.GetDB(ctx).Create(&m).Error
	return
}

func (dao *Dao) DeleteMemoRelations(ctx context.Context, req *model.DeleteMemoRelationsRequest) (err error) {
	var (
		where []string
		args  []any
	)
	if req.MemoID != 0 {
		where = append(where, "memo_id=?")
		args = append(args, req.MemoID)
	}
	if req.Type != "" {
		where = append(where, "type=?")
		args = append(args, req.Type)
	}
	if req.RelatedMemoID != 0 {
		where = append(where, "related_memo_id=?")
		args = append(args, req.RelatedMemoID)
	}
	if len(where) == 0 {
		err = errors.New("no where condition")
		return
	}
	err = db.GetDB(ctx).Model(&model.MemoRelation{}).Delete(strings.Join(where, " AND "), args...).Error
	return
}
