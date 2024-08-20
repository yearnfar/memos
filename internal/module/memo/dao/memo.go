package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (dao *Dao) FindMemos(ctx context.Context, req *model.FindMemosRequest) (memos []*model.Memo, err error) {
	return
}
