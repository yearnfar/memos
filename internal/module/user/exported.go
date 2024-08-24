// Code generated by service-export(1.0.0). DO NOT EDIT.
// source: service.go

package user

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
)

var defaultService Service

func Register(s Service) {
	defaultService = s
}

func CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.CreateUser 失败，服务未注册")
	}
	v1, v2 := defaultService.CreateUser(ctx, req)
	return v1, v2
}

func SignUp(ctx context.Context, req *model.SignUpRequest) (*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.SignUp 失败，服务未注册")
	}
	v1, v2 := defaultService.SignUp(ctx, req)
	return v1, v2
}

func GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.GetUserByUsername 失败，服务未注册")
	}
	v1, v2 := defaultService.GetUserByUsername(ctx, username)
	return v1, v2
}

func GetUserById(ctx context.Context, id int32) (*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.GetUserById 失败，服务未注册")
	}
	v1, v2 := defaultService.GetUserById(ctx, id)
	return v1, v2
}

func UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.UpdateUser 失败，服务未注册")
	}
	v1, v2 := defaultService.UpdateUser(ctx, req)
	return v1, v2
}

func ListUsers(ctx context.Context, req *model.ListUsersRequest) ([]*model.User, error) {
	if defaultService == nil {
		panic("调用模块方法: user.ListUsers 失败，服务未注册")
	}
	v1, v2 := defaultService.ListUsers(ctx, req)
	return v1, v2
}

func UpsertAccessToken(ctx context.Context, userId int32, token *model.AccessToken) error {
	if defaultService == nil {
		panic("调用模块方法: user.UpsertAccessToken 失败，服务未注册")
	}
	v1 := defaultService.UpsertAccessToken(ctx, userId, token)
	return v1
}

func DeleteAccessToken(ctx context.Context, userId int32, accessToken string) error {
	if defaultService == nil {
		panic("调用模块方法: user.DeleteAccessToken 失败，服务未注册")
	}
	v1 := defaultService.DeleteAccessToken(ctx, userId, accessToken)
	return v1
}

func GetAccessTokens(ctx context.Context, userId int32) ([]*model.AccessToken, error) {
	if defaultService == nil {
		panic("调用模块方法: user.GetAccessTokens 失败，服务未注册")
	}
	v1, v2 := defaultService.GetAccessTokens(ctx, userId)
	return v1, v2
}

func CreateUserAccessToken(ctx context.Context, req *model.CreateUserAccessTokenRequest) (*model.AccessToken, error) {
	if defaultService == nil {
		panic("调用模块方法: user.CreateUserAccessToken 失败，服务未注册")
	}
	v1, v2 := defaultService.CreateUserAccessToken(ctx, req)
	return v1, v2
}

func GetUserSettings(ctx context.Context, userId int32) ([]*model.UserSetting, error) {
	if defaultService == nil {
		panic("调用模块方法: user.GetUserSettings 失败，服务未注册")
	}
	v1, v2 := defaultService.GetUserSettings(ctx, userId)
	return v1, v2
}

func UpdateUserSetting(ctx context.Context, req *model.UpdateUserSettingRequest) error {
	if defaultService == nil {
		panic("调用模块方法: user.UpdateUserSetting 失败，服务未注册")
	}
	v1 := defaultService.UpdateUserSetting(ctx, req)
	return v1
}
