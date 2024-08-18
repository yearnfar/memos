package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
)

// GenerateAccessToken generates an access token.
func (s *Service) GenerateAccessToken(userID int, expirationTime time.Time, secret []byte) (string, error) {
	return s.generateToken(userID, model.AccessTokenAudienceName, expirationTime, secret)
}

// generateToken generates a jwt token.
func (s *Service) generateToken(userID int, audience string, expirationTime time.Time, secret []byte) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Issuer:   model.Issuer,
		Audience: jwt.ClaimStrings{audience},
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Subject:  fmt.Sprint(userID),
	}
	if !expirationTime.IsZero() {
		registeredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	token.Header["kid"] = model.KeyID

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (in *Service) Authenticate(ctx context.Context, accessToken, secret string) (user *usermodel.User, err error) {
	if accessToken == "" {
		err = errors.New("access token not found")
		return
	}
	claims := &jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(secret), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		err = errors.New("Invalid or expired access token")
		return
	}
	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return
	}
	user, err = usermod.GetUserById(ctx, userId)
	if err != nil {
		return
	}
	if user.RowStatus == usermodel.Archived {
		err = errors.Errorf("user %q is archived", user.ID)
		return
	}
	tokens, err := usermod.GetAccessTokens(ctx, user.ID)
	if err != nil {
		return
	}
	var valid bool
	for _, token := range tokens {
		if accessToken == token.Token {
			valid = true
			break
		}
	}
	if !valid {
		err = errors.New("invalid access token")
		return
	}
	return
}
