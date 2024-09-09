package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/yearnfar/memos/internal/api"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type AuthService struct {
	api.BaseService
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) GetAuthStatus(ctx context.Context, _ *v1pb.GetAuthStatusRequest) (userInfo *v1pb.User, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Unauthenticated, "failed to get current user: %v", err)
		return
	}
	if user == nil {
		if err = s.ClearAccessTokenCookie(ctx); err != nil {
			err = status.Errorf(codes.Internal, "failed to set grpc header: %v", err)
		} else {
			err = status.Errorf(codes.Unauthenticated, "user not found")
		}
		return
	}
	return convertUserFromStore(user), nil
}

func (s *AuthService) SignIn(ctx context.Context, request *v1pb.SignInRequest) (userInfo *v1pb.User, err error) {
	if err = s.DoSignIn(ctx, request.Username, request.Password); err != nil {
		return
	}
	user, err := usermod.GetUser(ctx, &model.GetUserRequest{Username: request.Username})
	if err != nil {
		return
	}
	return convertUserFromStore(user), nil
}

func (s *AuthService) SignInWithSSO(ctx context.Context, request *v1pb.SignInWithSSORequest) (*v1pb.User, error) {
	return nil, nil
}

func (s *AuthService) SignUp(ctx context.Context, request *v1pb.SignUpRequest) (userInfo *v1pb.User, err error) {
	req := &usermodel.SignUpRequest{
		Username: request.Username,
		Password: request.Password,
	}
	user, err := usermod.SignUp(ctx, req)
	if err != nil {
		return
	}
	if err = s.DoSignIn(ctx, req.Username, req.Password); err != nil {
		return
	}
	return convertUserFromStore(user), nil
}

func (s *AuthService) SignOut(ctx context.Context, request *v1pb.SignOutRequest) (*emptypb.Empty, error) {
	return nil, nil
}
