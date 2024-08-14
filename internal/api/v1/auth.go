package v1

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	"github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type AuthService struct {
	api.BaseService
}

func (s *AuthService) GetAuthStatus(ctx context.Context, _ *v1pb.GetAuthStatusRequest) (*v1pb.User, error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get current user: %v", err)
	}
	if user == nil {
		// Set the cookie header to expire access token.
		if err := s.ClearAccessTokenCookie(ctx); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to set grpc header: %v", err)
		}
		return nil, status.Errorf(codes.Unauthenticated, "user not found")
	}
	return convertUserFromStore(user), nil
}

func (s *AuthService) SignIn(ctx context.Context, request *v1pb.SignInRequest) (*v1pb.User, error) {

	return convertUserFromStore(user), nil
}

func convertUserFromStore(user *user.UserInfo) *v1pb.User {
	userpb := &v1pb.User{
		Name:        fmt.Sprintf("%s%d", api.UserNamePrefix, user.ID),
		Id:          user.ID,
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

func convertRowStatusFromStore(rowStatus model.RowStatus) v1pb.RowStatus {
	switch rowStatus {
	case model.Normal:
		return v1pb.RowStatus_ACTIVE
	case model.Archived:
		return v1pb.RowStatus_ARCHIVED
	default:
		return v1pb.RowStatus_ROW_STATUS_UNSPECIFIED
	}
}

func convertUserRoleFromStore(role model.Role) v1pb.User_Role {
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
