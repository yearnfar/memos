package v1

import (
	"context"

	"github.com/yearnfar/memos/internal/api"

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
		// InstanceUrl:  "",
	}
	// owner, err := s.GetInstanceOwner(ctx)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to get instance owner: %v", err)
	// }
	// if owner != nil {
	// workspaceProfile.Owner = owner.Name
	// } else {
	// If owner is not found, set Public/PasswordAuth to true.
	workspaceProfile.Owner = "yearnfar"
	workspaceProfile.Public = true
	workspaceProfile.PasswordAuth = true
	// }
	return workspaceProfile, nil
}

// func (s *WorkspaceService) GetInstanceOwner(ctx context.Context) (*v1pb.User, error) {
// 	if ownerCache != nil {
// 		return ownerCache, nil
// 	}

// 	hostUserType := store.RoleHost
// 	user, err := s.Store.GetUser(ctx, &store.FindUser{
// 		Role: &hostUserType,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "failed to find owner")
// 	}
// 	if user == nil {
// 		return nil, nil
// 	}

// 	ownerCache = convertUserFromStore(user)
// 	return ownerCache, nil
// }
