package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	authmod "github.com/yearnfar/memos/internal/module/auth"
	authmdl "github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type UserService struct {
	api.BaseService
	v1pb.UnimplementedUserServiceServer
}

func (s *UserService) ListUsers(ctx context.Context, _ *v1pb.ListUsersRequest) (*v1pb.ListUsersResponse, error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if currentUser.Role != model.RoleHost && currentUser.Role != model.RoleAdmin {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	users, err := usermod.ListUsers(ctx, &model.ListUsersRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}

	response := &v1pb.ListUsersResponse{
		Users: []*v1pb.User{},
	}
	for _, user := range users {
		response.Users = append(response.Users, convertUserFromStore(user))
	}
	return response, nil
}

func (s *UserService) CreateUser(ctx context.Context, request *v1pb.CreateUserRequest) (*v1pb.User, error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if currentUser.Role != model.RoleHost {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}
	user, err := usermod.CreateUser(ctx, &model.CreateUserRequest{
		Username: request.User.Username,
		Nickname: request.User.Nickname,
		Role:     model.Role(request.User.Role.String()),
		Email:    request.User.Email,
		Password: request.User.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create user fail: %v", err)
	}
	return convertUserFromStore(user), nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1pb.UpdateUserRequest) (userInfo *v1pb.User, err error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
		return
	}
	userID, err := api.ExtractUserIDFromName(req.User.Name)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
		return
	}
	if currentUser.ID != userID && currentUser.Role != model.RoleAdmin && currentUser.Role != model.RoleHost {
		err = status.Errorf(codes.PermissionDenied, "permission denied")
		return
	}
	user, err := usermod.UpdateUser(ctx, &model.UpdateUserRequest{
		UpdateMasks: req.UpdateMask.Paths,
		UserId:      userID,
		Username:    req.User.Username,
		Role:        model.Role(req.User.Role.String()),
		RowStatus:   convertRowStatusToStore(req.User.RowStatus),
		Email:       req.User.Email,
		AvatarURL:   req.User.AvatarUrl,
		Nickname:    req.User.Nickname,
		Password:    req.User.Password,
		Description: req.User.Description,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "update user error: %v", err)
		return
	}
	userInfo = convertUserFromStore(user)
	return
}
func (s *UserService) DeleteUser(ctx context.Context, request *v1pb.DeleteUserRequest) (response *emptypb.Empty, err error) {
	userID, err := api.ExtractUserIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
	}
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if currentUser.ID != userID && currentUser.Role != model.RoleAdmin && currentUser.Role != model.RoleHost {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if err = usermod.DeleteUserById(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return
}

func (s *UserService) GetUserSetting(ctx context.Context, _ *v1pb.GetUserSettingRequest) (response *v1pb.UserSetting, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	userSettings, err := usermod.GetUserSettings(ctx, user.ID)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to list user settings: %v", err)
		return
	}
	userSettingMessage := s.getDefaultUserSetting()
	for _, setting := range userSettings {
		var item *model.UserSettingValue
		err = json.Unmarshal([]byte(setting.Value), &item)
		if err != nil {
			return
		}
		if setting.Key == model.UserSettingKeyLocale {
			userSettingMessage.Locale = item.Locale
		} else if setting.Key == model.UserSettingKeyAppearance {
			userSettingMessage.Appearance = item.Appearance
		} else if setting.Key == model.UserSettingKeyMemoVisibility {
			userSettingMessage.MemoVisibility = item.MemoVisibility
		}
	}
	return userSettingMessage, nil
}

func (s *UserService) getDefaultUserSetting() *v1pb.UserSetting {
	return &v1pb.UserSetting{
		Locale:         "en",
		Appearance:     "system",
		MemoVisibility: "PRIVATE",
	}
}

func (s *UserService) UpdateUserSetting(ctx context.Context, request *v1pb.UpdateUserSettingRequest) (response *v1pb.UserSetting, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	err = usermod.UpdateUserSetting(ctx, &model.UpdateUserSettingRequest{
		UpdateMasks:    request.UpdateMask.Paths,
		UserID:         user.ID,
		Locale:         request.Setting.Locale,
		Appearance:     request.Setting.Appearance,
		MemoVisibility: request.Setting.MemoVisibility,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get update user setting: %v", err)
		return
	}
	return s.GetUserSetting(ctx, nil)
}

func (s *UserService) CreateUserAccessToken(ctx context.Context, request *v1pb.CreateUserAccessTokenRequest) (response *v1pb.UserAccessToken, err error) {
	userID, err := api.ExtractUserIDFromName(request.Name)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
		return
	}
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	if currentUser == nil || currentUser.ID != userID {
		err = status.Errorf(codes.PermissionDenied, "permission denied")
		return
	}
	accessToken, err := usermod.CreateUserAccessToken(ctx, &model.CreateUserAccessTokenRequest{
		UserID:      userID,
		Description: request.Description,
		ExpiresAt:   request.ExpiresAt.AsTime(),
		Audience:    authmdl.AccessTokenAudienceName,
		KeyID:       authmdl.KeyID,
	})
	if err != nil {
		return
	}
	authToken, err := authmod.Authenticate(ctx, accessToken.Token, authmdl.KeyID)
	if err != nil {
		return
	}
	response = &v1pb.UserAccessToken{
		AccessToken: authToken.Token,
		Description: request.Description,
		IssuedAt:    timestamppb.New(time.Unix(authToken.IssuedAt, 0)),
		ExpiresAt:   timestamppb.New(time.Unix(authToken.ExpiresAt, 0)),
	}
	return
}

func (s *UserService) ListUserAccessTokens(ctx context.Context, request *v1pb.ListUserAccessTokensRequest) (response *v1pb.ListUserAccessTokensResponse, err error) {
	userID, err := api.ExtractUserIDFromName(request.Name)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
		return
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	if user == nil || user.ID != userID {
		err = status.Errorf(codes.PermissionDenied, "permission denied")
		return
	}
	accessTokens, err := usermod.GetAccessTokens(ctx, userID)
	if err != nil {
		return
	}
	var userAccessTokens []*v1pb.UserAccessToken
	for _, accessToken := range accessTokens {
		authToken, err2 := authmod.Authenticate(ctx, accessToken.Token, authmdl.KeyID)
		if err2 != nil {
			continue
		}
		userAccessTokens = append(userAccessTokens, &v1pb.UserAccessToken{
			AccessToken: authToken.Token,
			Description: accessToken.Description,
			IssuedAt:    timestamppb.New(time.Unix(authToken.IssuedAt, 0)),
			ExpiresAt:   timestamppb.New(time.Unix(authToken.ExpiresAt, 0)),
		})
	}
	// Sort by issued time in descending order.
	slices.SortFunc(userAccessTokens, func(i, j *v1pb.UserAccessToken) int {
		return int(i.IssuedAt.Seconds - j.IssuedAt.Seconds)
	})
	response = &v1pb.ListUserAccessTokensResponse{
		AccessTokens: userAccessTokens,
	}
	return
}

func (s *UserService) DeleteUserAccessToken(ctx context.Context, request *v1pb.DeleteUserAccessTokenRequest) (response *emptypb.Empty, err error) {
	userID, err := api.ExtractUserIDFromName(request.Name)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid user name: %v", err)
		return
	}
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	if currentUser == nil || currentUser.ID != userID {
		err = status.Errorf(codes.PermissionDenied, "permission denied")
		return
	}
	err = usermod.DeleteAccessToken(ctx, userID, request.AccessToken)
	return
}

func convertUserFromStore(user *model.User) *v1pb.User {
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

func convertRowStatusToStore(rowStatus v1pb.RowStatus) model.RowStatus {
	switch rowStatus {
	case v1pb.RowStatus_ACTIVE:
		return model.Normal
	case v1pb.RowStatus_ARCHIVED:
		return model.Archived
	default:
		return model.Normal
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
