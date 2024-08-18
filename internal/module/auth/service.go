//go:generate service-export ./$GOFILE
//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package auth

import (
	"context"
	"time"

	"github.com/yearnfar/memos/internal/module/auth/model"
)

type Service interface {
	SignIn(ctx context.Context, req *model.SignInRequest) (resp *model.SignInResponse, err error)
	GenerateAccessToken(username string, userId int, expirationTime time.Time, secret []byte) (string, error)
}
