package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/api"
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
	user, err := usermod.UpdateUser(ctx, &model.UpdateUserRequest{
		UpdateMasks: req.UpdateMask.Paths,
		UserId:      currentUser.ID,
		Username:    req.User.Name,
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
