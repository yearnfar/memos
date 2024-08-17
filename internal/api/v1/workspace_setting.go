package v1

import (
	"context"

	"github.com/yearnfar/memos/internal/api"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type WorkspaceSettingService struct {
	api.BaseService
	v1pb.UnimplementedWorkspaceSettingServiceServer
}

func (s *WorkspaceSettingService) GetWorkspaceSetting(ctx context.Context, request *v1pb.GetWorkspaceSettingRequest) (workspaceSetting *v1pb.WorkspaceSetting, err error) {

	return
}

func (s *WorkspaceSettingService) SetWorkspaceSetting(ctx context.Context, request *v1pb.SetWorkspaceSettingRequest) (workspaceSetting *v1pb.WorkspaceSetting, err error) {

	return
}
