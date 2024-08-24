//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
)

type Service interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error)
	SignUp(ctx context.Context, req *model.SignUpRequest) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserById(ctx context.Context, id int32) (*model.User, error)
	UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (*model.User, error)
	ListUsers(ctx context.Context, req *model.ListUsersRequest) ([]*model.User, error)

	UpsertAccessToken(ctx context.Context, userId int32, token *model.AccessToken) error
	DeleteAccessToken(ctx context.Context, userId int32, accessToken string) error
	GetAccessTokens(ctx context.Context, userId int32) ([]*model.AccessToken, error)
	CreateUserAccessToken(ctx context.Context, req *model.CreateUserAccessTokenRequest) (*model.AccessToken, error)

	GetUserSettings(ctx context.Context, userId int32) ([]*model.UserSetting, error)
	UpdateUserSetting(ctx context.Context, req *model.UpdateUserSettingRequest) error
}
