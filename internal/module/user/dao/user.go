package dao

import (
	"context"
	"strings"

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

func (Dao) FindUser(ctx context.Context, where []string, args []any, fields ...string) (*model.User, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var user model.User
	err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (Dao) UpdateUser(ctx context.Context, user *model.User, update map[string]any) (err error) {
	err = db.GetDB(ctx).Model(user).Updates(update).Error
	return
}

func (Dao) DeleteUserById(ctx context.Context, id int32) (err error) {
	err = db.GetDB(ctx).Model(&model.User{}).Delete("id=?", id).Error
	return
}

func (Dao) FindUsers(ctx context.Context, where []string, args []any, fields ...string) (list []*model.User, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}
