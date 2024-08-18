//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import (
	"context"
	"time"

	"github.com/yearnfar/memos/internal/module/auth/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (resp *model.SignInResponse, err error)
	GenerateAccessToken(userId int, expirationTime time.Time, secret []byte) (string, error)
	Authenticate(ctx context.Context, accessToken, secret string) (user *usermodel.User, err error)
}
