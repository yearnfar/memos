package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/internal/module/auth/model"
	usermod "github.com/yearnfar/memos/internal/module/user"
	usermdl "github.com/yearnfar/memos/internal/module/user/model"
)

func TestService_SignIn(t *testing.T) {
	config.Init("../../../../")

	ctx := context.Background()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
		return
	}
	ctl := gomock.NewController(t)
	userSvc := usermod.NewMockService(ctl)
	userSvc.
		EXPECT().
		GetUser(ctx, &usermdl.GetUserRequest{Username: "yearnfar"}).
		Return(&usermdl.User{
			Username:     "yearnfar",
			PasswordHash: string(passwordHash)}, nil)

	userSvc.
		EXPECT().
		UpsertAccessToken(ctx, gomock.Any(), gomock.Any()).
		Return(nil)

	usermod.Register(userSvc)

	req := &model.SignInRequest{
		Username:    "yearnfar",
		Password:    "123456",
		NeverExpire: false,
	}

	resp, err := New(nil).SignIn(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sign_in: %v", resp)
}
