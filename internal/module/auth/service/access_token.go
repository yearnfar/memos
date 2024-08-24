package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/internal/module/auth/model"
)

// GenerateAccessToken generates an access token.
func (s *Service) GenerateAccessToken(_ context.Context, userID int32, expirationTime time.Time) (*model.AccessToken, error) {
	cfg := config.GetApp().JWT
	tokenStr, issuedAt, err := s.generateToken(userID, model.AccessTokenAudienceName, expirationTime, []byte(cfg.Key))
	if err != nil {
		return nil, err
	}
	return &model.AccessToken{
		UserId:    userID,
		Token:     tokenStr,
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  issuedAt.Unix(),
	}, nil
}

// generateToken generates a jwt token.
func (s *Service) generateToken(userID int32, audience string, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	issuedAt := time.Now()
	registeredClaims := jwt.RegisteredClaims{
		Issuer:   model.Issuer,
		Audience: jwt.ClaimStrings{audience},
		IssuedAt: jwt.NewNumericDate(issuedAt),
		Subject:  fmt.Sprint(userID),
	}
	if !expirationTime.IsZero() {
		registeredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	token.Header["kid"] = model.KeyID

	// Create the JWT string.
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenStr, issuedAt, nil
}

func (in *Service) Authenticate(ctx context.Context, tokenStr string) (accessToken *model.AccessToken, err error) {
	if tokenStr == "" {
		err = errors.New("access token not found")
		return
	}
	cfg := config.GetApp().JWT
	claims := &jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(cfg.Key), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		err = errors.New("Invalid or expired access token")
		return
	}
	userId, err := strconv.ParseInt(claims.Subject, 10, 32)
	if err != nil {
		return
	}
	accessToken = &model.AccessToken{
		UserId: int32(userId),
		Token:  tokenStr,
	}
	if claims.IssuedAt != nil {
		accessToken.IssuedAt = claims.ExpiresAt.Unix()
	}
	if claims.ExpiresAt != nil {
		accessToken.ExpiresAt = claims.ExpiresAt.Unix()
	}
	return
}
