package interceptor

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	authmod "github.com/yearnfar/memos/internal/module/auth"
	"github.com/yearnfar/memos/internal/module/user/model"
)

// ContextKey is the key type of context value.
type ContextKey int

const (
	// The key name used to store username in the context
	// user id is extracted from the jwt token subject field.
	usernameContextKey ContextKey = iota
	accessTokenContextKey
)

// GRPCAuthInterceptor is the auth interceptor for gRPC server.
type GRPCAuthInterceptor struct {
	secret string
}

// NewGRPCAuthInterceptor returns a new API auth interceptor.
func NewGRPCAuthInterceptor(secret string) *GRPCAuthInterceptor {
	return &GRPCAuthInterceptor{
		secret: secret,
	}
}

// AuthenticationInterceptor is the unary interceptor for gRPC API.
func (in *GRPCAuthInterceptor) AuthenticationInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "failed to parse metadata from incoming context")
	}
	accessToken, err := getTokenFromMetadata(md)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	userId, err := authmod.Authenticate(ctx, accessToken, "")
	if err != nil {
		return nil, err
	}

	username, err := in.authenticate(ctx, accessToken)
	if err != nil {
		if isUnauthorizeAllowedMethod(serverInfo.FullMethod) {
			return handler(ctx, request)
		}
		return nil, err
	}

	if isOnlyForAdminAllowedMethod(serverInfo.FullMethod) && user.Role != store.RoleHost && user.Role != store.RoleAdmin {
		return nil, errors.Errorf("user %q is not admin", username)
	}

	ctx = context.WithValue(ctx, usernameContextKey, username)
	ctx = context.WithValue(ctx, accessTokenContextKey, accessToken)
	return handler(ctx, request)
}

func getTokenFromMetadata(md metadata.MD) (string, error) {
	// Check the HTTP request header first.
	authorizationHeaders := md.Get("Authorization")
	if len(md.Get("Authorization")) > 0 {
		authHeaderParts := strings.Fields(authorizationHeaders[0])
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return "", errors.New("authorization header format must be Bearer {token}")
		}
		return authHeaderParts[1], nil
	}
	// Check the cookie header.
	var accessToken string
	for _, t := range append(md.Get("grpcgateway-cookie"), md.Get("cookie")...) {
		header := http.Header{}
		header.Add("Cookie", t)
		request := http.Request{Header: header}
		if v, _ := request.Cookie(AccessTokenCookieName); v != nil {
			accessToken = v.Value
		}
	}
	return accessToken, nil
}

func validateAccessToken(token string, userAccessTokens []*model.AccessToken) bool {
	for _, userAccessToken := range userAccessTokens {
		if token == userAccessToken.Token {
			return true
		}
	}
	return false
}
