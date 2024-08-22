package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindResources(ctx context.Context, req *model.FindResourcesRequest) (list []*model.Resource, err error) {
	conn := db.GetDB(ctx)
	err = conn.Find(&list).Error
	return
}
