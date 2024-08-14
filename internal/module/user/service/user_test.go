package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yearnfar/memos/internal/config"
	usermod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/model"
)

func TestService_CreateUser(t *testing.T) {
	config.Init("../../../../")

	ctx := context.Background()

	ctl := gomock.NewController(t)
	daoMock := usermod.NewMockDAO(ctl)
	daoMock.
		EXPECT().
		CreateUser(ctx, gomock.Any()).
		Return(nil)

	req := &model.CreateUserRequest{
		Username: "yearnfar",
		Role:     model.RoleHost,
		Email:    "yearnfar@gmail.com",
		Nickname: "yearnfar",
		Password: "123456",
	}
	user, err := New(daoMock).CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("user: %+v", user)
}
