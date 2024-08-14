//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package user

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
)

type DAO interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUser(ctx context.Context, req *model.GetUserRequest) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, update map[string]any) error

	UpsertUserSetting(ctx context.Context, m *model.UserSetting) (err error)
	GetUserAccessTokens(ctx context.Context, userId int) ([]*model.AccessToken, error)
}
