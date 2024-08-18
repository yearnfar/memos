package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
)

// SignIn 登录
func (s *Service) SignIn(ctx context.Context, req *model.SignInRequest) (resp *model.SignInResponse, err error) {
	user, err := usermod.GetUserByUsername(ctx, req.Username)
	if err != nil {
		err = errors.Errorf("failed to find user by username %s", req.Username)
		return
	}
	if user.RowStatus == usermodel.Archived {
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
	accessToken, err := s.doSignIn(ctx, user, expireTime)
	if err != nil {
		err = errors.Errorf("failed to sign in, err: %s", err)
		return
	}
	resp = &model.SignInResponse{
		AccessToken: accessToken,
		ExpireTime:  expireTime,
	}
	return
}

func (s *Service) doSignIn(ctx context.Context, user *usermodel.User, expireTime time.Time) (accessToken string, err error) {
	cfg := config.GetApp().JWT
	accessToken, err = s.GenerateAccessToken(user.ID, expireTime, []byte(cfg.Key))
	if err != nil {
		err = errors.Errorf("failed to generate tokens, err: %s", err)
		return
	}
	if err = usermod.UpsertAccessToken(ctx, user.ID, accessToken, "user login"); err != nil {
		err = errors.Errorf("failed to upsert access token to store, err: %s", err)
		return
	}
	return
}

func (s *Service) SignOut(ctx context.Context, req *model.SignOutRequest) (err error) {
	err = usermod.DeleteAccessToken(ctx, req.UserId, req.AccessToken)
	return
}
