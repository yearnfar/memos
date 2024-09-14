package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	authmod "github.com/yearnfar/memos/internal/module/auth"
	authmodel "github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
)

// ContextKey is the key type of context value.
type ContextKey int

const (
	// The key name used to store username in the context
	// user id is extracted from the jwt token subject field.
	userContextKey ContextKey = iota
	accessTokenContextKey
)

type BaseService struct {
}

func (s *BaseService) GetCurrentUser(ctx context.Context) (userInfo *usermodel.User, err error) {
	userId, ok := ctx.Value(userContextKey).(int32)
	if !ok {
		return nil, nil
	}
	userInfo, err = usermod.GetUserById(ctx, userId)
	return
}

func (s *BaseService) ClearAccessTokenCookie(ctx context.Context) error {
	cookie, err := s.buildAccessTokenCookie(ctx, "", time.Time{})
	if err != nil {
		return errors.Wrap(err, "failed to build access token cookie")
	}
	if err := grpc.SetHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie})); err != nil {
		return errors.Wrap(err, "failed to set grpc header")
	}
	return nil
}

func (s *BaseService) DoSignIn(ctx context.Context, username, password string) (err error) {
	resp, err := authmod.SignIn(ctx, &authmodel.SignInRequest{
		Audience: authmodel.AccessTokenAudienceName,
		KeyID:    authmodel.KeyID,
		Username: username,
		Password: password,
	})
	if err != nil {
		return
	}
	cookie, err := s.buildAccessTokenCookie(ctx, resp.AccessToken, resp.ExpireTime)
	if err != nil {
		err = errors.Errorf("failed to build access token cookie, err: %s", err)
		return
	}
	if err = grpc.SetHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie})); err != nil {
		err = errors.Errorf("failed to set grpc header, error: %v", err)
		return
	}
	return nil
}

func (s *BaseService) buildAccessTokenCookie(ctx context.Context, accessToken string, expireTime time.Time) (string, error) {
	attrs := []string{
		fmt.Sprintf("%s=%s", authmodel.AccessTokenCookieName, accessToken),
		"Path=/",
		"HttpOnly",
	}
	if expireTime.IsZero() {
		attrs = append(attrs, "Expires=Thu, 01 Jan 1970 00:00:00 GMT")
	} else {
		attrs = append(attrs, "Expires="+expireTime.Format(time.RFC1123))
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("failed to get metadata from context")
	}
	var origin string
	for _, v := range md.Get("origin") {
		origin = v
	}
	if strings.HasPrefix(origin, "https://") {
		attrs = append(attrs, "SameSite=None")
		attrs = append(attrs, "Secure")
	} else {
		attrs = append(attrs, "SameSite=Strict")
	}
	return strings.Join(attrs, "; "), nil
}
