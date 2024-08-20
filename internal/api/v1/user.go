package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/api"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
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
