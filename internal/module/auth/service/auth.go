package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermdl "github.com/yearnfar/memos/internal/module/user/model"
)

// SignIn 登录
func (s *Service) SignIn(ctx context.Context, req *model.SignInRequest) (resp *model.SignInResponse, err error) {
	user, err := usermod.GetUser(ctx, &usermdl.GetUserRequest{Username: req.Username})
	if err != nil {
		err = errors.Errorf("failed to find user by username %s", req.Username)
		return
	}
	if user.RowStatus == usermdl.Archived {
		err = errors.Errorf("user has been archived with username %s", req.Username)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received.
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		err = errors.New("unmatched email and password")
		return
	}
	expireTime := time.Now().Add(model.AccessTokenDuration)
	if req.NeverExpire {
		// Set the expire time to 100 years.
		expireTime = time.Now().Add(100 * 365 * 24 * time.Hour)
	}
	tokenStr, err := s.doSignIn(ctx, user, req.Audience, req.KeyID, expireTime)
	if err != nil {
		err = errors.Errorf("failed to sign in, err: %s", err)
		return
	}
	resp = &model.SignInResponse{
		AccessToken: tokenStr,
		ExpireTime:  expireTime,
	}
	return
}

func (s *Service) doSignIn(ctx context.Context, user *usermdl.User, audience, keyId string, expireTime time.Time) (tokenStr string, err error) {
	accessToken, err := s.GenerateAccessToken(ctx, user.ID, audience, keyId, expireTime)
	if err != nil {
		err = errors.Errorf("failed to generate tokens, err: %s", err)
		return
	}
	token := &usermdl.AccessToken{
		Token:       accessToken.Token,
		Description: "user login",
	}
	if err = usermod.UpsertAccessToken(ctx, user.ID, token); err != nil {
		err = errors.Errorf("failed to upsert access token to store, err: %s", err)
		return
	}
	return accessToken.Token, nil
}

func (s *Service) SignOut(ctx context.Context, req *model.SignOutRequest) (err error) {
	err = usermod.DeleteAccessToken(ctx, req.UserId, req.AccessToken)
	return
}
