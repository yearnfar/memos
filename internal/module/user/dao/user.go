package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (Dao) CreateUser(ctx context.Context, user *model.User) error {
	return db.GetDB(ctx).Create(user).Error
}

func (Dao) FindUserById(ctx context.Context, id int32) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB(ctx).Where("id=?", id).First(&user).Error
	return
}

func (Dao) FindUserByUsername(ctx context.Context, username string) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB(ctx).Where("username=?", username).First(&user).Error
	return
}

func (Dao) FindUser(ctx context.Context, req *model.FindUserRequest) (user *model.User, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn = conn.Where("id=?", req.Id)
	}
	if req.Username != "" {
		conn = conn.Where("username=?", req.Username)
	}
	user = &model.User{}
	err = conn.First(&user).Error
	return
}

func (Dao) UpdateUser(ctx context.Context, user *model.User, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(user).Updates(update).Error
	return
}

func (Dao) DeleteUserById(ctx context.Context, id int32) (err error) {
	err = db.GetDB(ctx).Model(&model.User{}).Delete("id=?", id).Error
	return
}

func (Dao) FindUsers(ctx context.Context, req *model.FindUsersRequest) (list []*model.User, err error) {
	conn := db.GetDB(ctx)
	if req.Role != "" {
		conn = conn.Where("role=?", req.Role)
	}
	err = conn.Find(&list).Error
	return
}
