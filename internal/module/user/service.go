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
	GetUserById(ctx context.Context, id int) (*model.User, error)

	UpsertAccessToken(ctx context.Context, userId int, accessToken, description string) error
	DeleteAccessToken(ctx context.Context, userId int, accessToken string) error
	GetAccessTokens(ctx context.Context, userId int) ([]*model.AccessToken, error)
}
