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
	data, err := json.Marshal(model.UserSettingValue{AccessTokens: append(tokens, token)})
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
	var val *model.UserSettingValue
	for _, token := range tokens {
		if accessToken != token.Token {
			val.AccessTokens = append(val.AccessTokens, token)
		}
	}
	data, err := json.Marshal(val)
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

func (s *Service) UpdateUserSetting(ctx context.Context, req *model.UpdateUserSettingRequest) (err error) {
	if len(req.UpdateMasks) == 0 {
		err = errors.New("update mask is empty")
		return
	}

	var key model.UserSettingKey
	var value model.UserSettingValue
	for _, field := range req.UpdateMasks {
		if field == "locale" {
			key = model.UserSettingKeyLocale
			value = model.UserSettingValue{Locale: req.Locale}
		} else if field == "appearance" {
			key = model.UserSettingKeyAppearance
			value = model.UserSettingValue{Appearance: req.Appearance}
		} else if field == "memo_visibility" {
			key = model.UserSettingKeyMemoVisibility
			value = model.UserSettingValue{MemoVisibility: req.MemoVisibility}
		} else {
			err = errors.Errorf("invalid update path: %s", field)
			return
		}
	}
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	if err = s.dao.UpsertUserSetting(ctx, &model.UserSetting{
		UserId: req.UserID,
		Key:    key,
		Value:  string(data),
	}); err != nil {
		err = errors.Errorf("failed to upsert user setting: %v", err)
		return
	}
	return
}
