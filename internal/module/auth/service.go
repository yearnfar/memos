//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import (
	"context"
	"time"

	"github.com/yearnfar/memos/internal/module/auth/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (*model.SignInResponse, error)
	GenerateAccessToken(ctx context.Context, userId int32, expirationTime time.Time) (*model.AccessToken, error)
	Authenticate(ctx context.Context, tokenStr string) (*model.AccessToken, error)
}
