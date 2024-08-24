package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	authmod "github.com/yearnfar/memos/internal/module/auth"
	"github.com/yearnfar/memos/internal/module/user/model"
)

func (s *Service) GetAccessTokens(ctx context.Context, userId int32) (tokens []*model.AccessToken, err error) {
	return s.dao.FindUserAccessTokens(ctx, userId)
}

func (s *Service) CreateUserAccessToken(ctx context.Context, req *model.CreateUserAccessTokenRequest) (token *model.AccessToken, err error) {
	authToken, err := authmod.GenerateAccessToken(ctx, req.UserID, req.ExpiresAt)
	if err != nil {
		err = errors.Errorf("failed to generate access token: %v", err)
		return
	}
	token = &model.AccessToken{
		Token:       authToken.Token,
		Description: req.Description,
	}
	err = s.UpsertAccessToken(ctx, req.UserID, token)
	return
}

func (s *Service) UpsertAccessToken(ctx context.Context, userId int32, token *model.AccessToken) (err error) {
	tokens, err := s.dao.FindUserAccessTokens(ctx, userId)
	if err != nil {
		err = errors.Wrap(err, "failed to get user access tokens")
		return
	}
	tokens = append(tokens, token)
	data, err := json.Marshal(tokens)
	if err != nil {
		return
	}
	err = s.dao.UpsertUserSetting(ctx, &model.UserSetting{
		UserId: userId,
		Key:    model.UserSettingKeyAccessToken,
		Value:  string(data),
	})
	if err != nil {
		err = errors.Wrap(err, "failed to upsert user setting")
		return
	}
	return
}

func (s *Service) DeleteAccessToken(ctx context.Context, userId int32, accessToken string) (err error) {
	tokens, err := s.dao.FindUserAccessTokens(ctx, userId)
	if err != nil {
		err = errors.Wrap(err, "failed to get user access tokens")
		return
	}
	var newTokens []*model.AccessToken
	for _, token := range tokens {
		if accessToken != token.Token {
			newTokens = append(newTokens, token)
		}
	}
	data, err := json.Marshal(newTokens)
	if err != nil {
		return
	}
	err = s.dao.UpsertUserSetting(ctx, &model.UserSetting{
		UserId: userId,
		Key:    model.UserSettingKeyAccessToken,
		Value:  string(data),
	})
	if err != nil {
		err = errors.Wrap(err, "failed to upsert user setting")
		return
	}
	return
}

func (s *Service) GetUserSettings(ctx context.Context, userId int32) ([]*model.UserSetting, error) {
	return s.dao.FindUserSettings(ctx, &model.FindUserSettingsRequest{UserId: userId})
}
