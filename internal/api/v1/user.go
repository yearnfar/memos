package v1

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/api"
	authmod "github.com/yearnfar/memos/internal/module/auth"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	api.BaseService
	v1pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUserSetting(ctx context.Context, _ *v1pb.GetUserSettingRequest) (*v1pb.UserSetting, error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, errors.Errorf("failed to get current user: %v", err)
	}

	userSettings, err := usermod.GetUserSettings(ctx, user.ID)
	if err != nil {
		return nil, errors.Errorf("failed to list user settings: %v", err)
	}
	userSettingMessage := s.getDefaultUserSetting()
	for _, setting := range userSettings {
		if setting.Key == usermodel.UserSettingKeyLocale {
			userSettingMessage.Locale = string(setting.Key)
		} else if setting.Key == usermodel.UserSettingKeyAppearance {
			userSettingMessage.Appearance = string(setting.Key)
		} else if setting.Key == usermodel.UserSettingKeyMemoVisibility {
			userSettingMessage.MemoVisibility = string(setting.Key)
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

func (s *UserService) UpdateUser(ctx context.Context, req *v1pb.UpdateUserRequest) (userInfo *v1pb.User, err error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get user: %v", err)
		return
	}
	userID, err := api.ExtractUserIDFromName(req.User.Name)
	if err != nil {
		err = errors.Errorf("invalid user name: %v", err)
		return
	}
	if currentUser.ID != userID && currentUser.Role != model.RoleAdmin && currentUser.Role != model.RoleHost {
		err = errors.New("permission denied")
		return
	}
	user, err := usermod.UpdateUser(ctx, &model.UpdateUserRequest{
		UpdateMasks: req.UpdateMask.Paths,
		UserId:      userID,
		Username:    req.User.Username,
		Role:        usermodel.Role(req.User.Role.String()),
		RowStatus:   usermodel.RowStatus(req.User.RowStatus.String()),
		Email:       req.User.Email,
		AvatarURL:   req.User.AvatarUrl,
		Nickname:    req.User.Nickname,
		Password:    req.User.Password,
		Description: req.User.Description,
	})
	if err != nil {
		return
	}
	userInfo = s.convertUserFromStore(user)
	return
}

func (s *UserService) CreateUserAccessToken(ctx context.Context, request *v1pb.CreateUserAccessTokenRequest) (response *v1pb.UserAccessToken, err error) {
	userID, err := api.ExtractUserIDFromName(request.Name)
	if err != nil {
		err = errors.Errorf("invalid user name: %v", err)
		return
	}
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get user: %v", err)
		return
	}
	if currentUser == nil || currentUser.ID != userID {
		err = errors.New("permission denied")
		return
	}
	accessToken, err := usermod.CreateUserAccessToken(ctx, &usermodel.CreateUserAccessTokenRequest{
		UserID:      userID,
		Description: request.Description,
		ExpiresAt:   request.ExpiresAt.AsTime(),
	})
	if err != nil {
		return
	}
	authToken, err := authmod.Authenticate(ctx, accessToken.Token)
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
		err = errors.Errorf("invalid user name: %v", err)
		return
	}
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}
	if currentUser == nil || currentUser.ID != userID {
		err = errors.New("permission denied")
		return
	}
	accessTokens, err := usermod.GetAccessTokens(ctx, userID)
	if err != nil {
		return
	}
	var userAccessTokens []*v1pb.UserAccessToken
	for _, accessToken := range accessTokens {
		authToken, err2 := authmod.Authenticate(ctx, accessToken.Token)
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

func (s *UserService) convertUserFromStore(user *model.User) *v1pb.User {
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
