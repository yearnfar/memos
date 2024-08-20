package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/module/user/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CreateUserAccessToken(ctx context.Context, request *v1pb.CreateUserAccessTokenRequest) (*v1pb.UserAccessToken, error) {
	return nil, nil
}

func (s *Service) DeleteUserAccessToken(ctx context.Context, request *v1pb.DeleteUserAccessTokenRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Service) UpsertAccessTokenToStore(ctx context.Context, user *model.User, accessToken, description string) error {
	return nil
}

func (s *Service) ListUserAccessTokens(ctx context.Context, request *v1pb.ListUserAccessTokensRequest) (*v1pb.ListUserAccessTokensResponse, error) {
	return nil, nil
}

func (s *Service) GetAccessTokens(ctx context.Context, userId int) (tokens []*model.AccessToken, err error) {
	return s.dao.FindUserAccessTokens(ctx, userId)
}

func (s *Service) UpsertAccessToken(ctx context.Context, userId int, accessToken, description string) (err error) {
	tokens, err := s.dao.FindUserAccessTokens(ctx, userId)
	if err != nil {
		err = errors.Wrap(err, "failed to get user access tokens")
		return
	}
	tokens = append(tokens, &model.AccessToken{Token: accessToken, Description: description})
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

func (s *Service) DeleteAccessToken(ctx context.Context, userId int, accessToken string) (err error) {
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

func (s *Service) GetUserSettings(ctx context.Context, userId int) ([]*model.UserSetting, error) {
	return s.dao.FindUserSettings(ctx, &model.FindUserSettingsRequest{UserId: userId})
}
