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

const (
	// issuer is the issuer of the jwt token.
	Issuer = "memos"
	// Signing key section. For now, this is only used for signing, not for verifying since we only
	// have 1 version. But it will be used to maintain backward compatibility if we change the signing mechanism.
	KeyID = "v1"
	// AccessTokenAudienceName is the audience name of the access token.
	AccessTokenAudienceName = "user.access-token"
	AccessTokenDuration     = 7 * 24 * time.Hour

	// CookieExpDuration expires slightly earlier than the jwt expiration. Client would be logged out if the user
	// cookie expires, thus the client would always logout first before attempting to make a request with the expired jwt.
	CookieExpDuration = AccessTokenDuration - 1*time.Minute
	// AccessTokenCookieName is the cookie name of access token.
	AccessTokenCookieName = "memos.access-token"
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
	userId, ok := ctx.Value(userContextKey).(int)
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
	if err := grpc.SetHeader(ctx, metadata.New(map[string]string{
		"Set-Cookie": cookie,
	})); err != nil {
		return errors.Wrap(err, "failed to set grpc header")
	}
	return nil
}

func (s *BaseService) DoSignIn(ctx context.Context, username, password string) (err error) {
	resp, err := authmod.SignIn(ctx, &authmodel.SignInRequest{
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
	if err = grpc.SetHeader(ctx, metadata.New(map[string]string{
		"Set-Cookie": cookie,
	})); err != nil {
		err = errors.Errorf("failed to set grpc header, error: %v", err)
		return
	}
	return nil
}

func (s *BaseService) buildAccessTokenCookie(ctx context.Context, accessToken string, expireTime time.Time) (string, error) {
	attrs := []string{
		fmt.Sprintf("%s=%s", AccessTokenCookieName, accessToken),
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
	isHTTPS := strings.HasPrefix(origin, "https://")
	if isHTTPS {
		attrs = append(attrs, "SameSite=None")
		attrs = append(attrs, "Secure")
	} else {
		attrs = append(attrs, "SameSite=Strict")
	}
	return strings.Join(attrs, "; "), nil
}

func SetContext(ctx context.Context, userId int, accessToken string) context.Context {
	ctx = context.WithValue(ctx, userContextKey, userId)
	ctx = context.WithValue(ctx, accessTokenContextKey, accessToken)
	return ctx
}
