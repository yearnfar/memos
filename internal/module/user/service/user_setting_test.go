package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yearnfar/memos/internal/config"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
)

func TestService_UpsertAccessToken(t *testing.T) {
	config.Init("../../../../")

	ctx := context.Background()

	ctl := gomock.NewController(t)
	daoMock := usermod.NewMockDAO(ctl)
	daoMock.
		EXPECT().
		FindUserAccessTokens(ctx, 1).
		Return([]*model.AccessToken{}, nil)

	daoMock.
		EXPECT().
		UpsertUserSetting(ctx, gomock.Any()).
		Return(nil)

	userId := 1
	accessToken := "dsafadsfasdfadsfasdf"
	description := "登录"

	err := New(daoMock).UpsertAccessToken(ctx, userId, accessToken, description)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("done")
}
