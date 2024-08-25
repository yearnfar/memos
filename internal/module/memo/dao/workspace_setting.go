package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (dao *Dao) FindWorkspaceSettings(ctx context.Context, req *model.FindWorkspaceSettingsRequest) (list []*model.WorkspaceSetting, err error) {
	conn := db.GetDB(ctx)
	if req.Name != "" {
		conn = conn.Where("name=?", req.Name)
	}
	err = conn.Find(&list).Error
	return
}

func (dao *Dao) UpsertWorkspaceSetting(ctx context.Context, m *model.WorkspaceSetting) (err error) {
	err = db.GetDB(ctx).
		Where("name=?", m.Name).
		Assign(model.WorkspaceSetting{Value: m.Value, Description: m.Description}).
		FirstOrCreate(&m).Error
	return
}
