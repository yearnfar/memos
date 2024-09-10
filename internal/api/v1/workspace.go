package v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/yearnfar/memos/internal/api"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type WorkspaceService struct {
	api.BaseService
	v1pb.UnimplementedWorkspaceServiceServer
}

func (s *WorkspaceService) GetWorkspaceProfile(ctx context.Context, _ *v1pb.GetWorkspaceProfileRequest) (*v1pb.WorkspaceProfile, error) {
	workspaceProfile := &v1pb.WorkspaceProfile{
		Version:      "",
		Mode:         "",
		Public:       true,
		PasswordAuth: true,
	}
	user, err := usermod.GetUser(ctx, &model.GetUserRequest{Role: model.RoleHost})
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "failed to get instance owner: %v", err)
	}
	if user != nil {
		workspaceProfile.Owner = fmt.Sprintf("%s%d", api.UserNamePrefix, user.ID)
	} else {
		// If owner is not found, set Public/PasswordAuth to true.
		workspaceProfile.Public = true
		workspaceProfile.PasswordAuth = true
	}
	return workspaceProfile, nil
}
