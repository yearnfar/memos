package v1

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	"github.com/yearnfar/memos/internal/module/user/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type AuthService struct {
	api.BaseService
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) GetAuthStatus(ctx context.Context, _ *v1pb.GetAuthStatusRequest) (*v1pb.User, error) {
	return nil, nil
}

func (s *AuthService) SignIn(ctx context.Context, request *v1pb.SignInRequest) (*v1pb.User, error) {
	return nil, nil
}

func (s *AuthService) SignInWithSSO(ctx context.Context, request *v1pb.SignInWithSSORequest) (*v1pb.User, error) {
	return nil, nil
}

func (s *AuthService) SignUp(ctx context.Context, req *v1pb.SignUpRequest) (*v1pb.User, error) {
	return nil, nil
}

func (s *AuthService) SignOut(ctx context.Context, request *v1pb.SignOutRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func convertUserFromStore(user *usermodel.User) *v1pb.User {
	userpb := &v1pb.User{
		Name:        fmt.Sprintf("%s%d", api.UserNamePrefix, user.ID),
		Id:          int32(user.ID),
		RowStatus:   convertRowStatusFromStore(user.RowStatus),
		CreateTime:  timestamppb.New(time.Unix(user.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(user.UpdatedTs, 0)),
		Role:        convertUserRoleFromStore(user.Role),
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		AvatarUrl:   user.AvatarURL,
		Description: user.Description,
	}
	// Use the avatar URL instead of raw base64 image data to reduce the response size.
	if user.AvatarURL != "" {
		userpb.AvatarUrl = fmt.Sprintf("/file/%s/avatar", userpb.Name)
	}
	return userpb
}

func convertRowStatusFromStore(rowStatus usermodel.RowStatus) v1pb.RowStatus {
	switch rowStatus {
	case model.Normal:
		return v1pb.RowStatus_ACTIVE
	case model.Archived:
		return v1pb.RowStatus_ARCHIVED
	default:
		return v1pb.RowStatus_ROW_STATUS_UNSPECIFIED
	}
}

func convertUserRoleFromStore(role usermodel.Role) v1pb.User_Role {
	switch role {
	case model.RoleHost:
		return v1pb.User_HOST
	case model.RoleAdmin:
		return v1pb.User_ADMIN
	case model.RoleUser:
		return v1pb.User_USER
	default:
		return v1pb.User_ROLE_UNSPECIFIED
	}
}
