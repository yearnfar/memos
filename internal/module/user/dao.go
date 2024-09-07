//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
)

type DAO interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserById(ctx context.Context, id int32) (*model.User, error)
	DeleteUserById(ctx context.Context, userId int32) error
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
	FindUser(ctx context.Context, where []string, args []any, fields ...string) (*model.User, error)
	FindUsers(ctx context.Context, where []string, args []any, fields ...string) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, update map[string]any) error

	FindUserSettings(ctx context.Context, where []string, args []any, fields ...string) ([]*model.UserSetting, error)
	UpsertUserSetting(ctx context.Context, m *model.UserSetting) (err error)
	FindUserAccessTokens(ctx context.Context, userId int32) ([]*model.AccessToken, error)
}
