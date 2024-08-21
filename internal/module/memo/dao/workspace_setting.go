package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (dao *Dao) FindWorkspaceSettings(ctx context.Context, req *model.FindWorkspaceSettingsRequest) (settings []*model.WorkspaceSetting, err error) {
	return
}
