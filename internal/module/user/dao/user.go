package dao

import (
	"context"

	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func (Dao) CreateUser(ctx context.Context, user *model.User) error {
	return db.GetDB(ctx).Create(user).Error
}

func (Dao) GetUserById(ctx context.Context, id int) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB(ctx).Where("id=?", id).First(&user).Error
	return
}

func (Dao) GetUserByUsername(ctx context.Context, username string) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB(ctx).Where("username=?", username).First(&user).Error
	return
}

func (Dao) GetUser(ctx context.Context, req *model.GetUserRequest) (user *model.User, err error) {
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
