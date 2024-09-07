package dao

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (Dao) GetUserSetting(ctx context.Context, where []string, args []any, fields ...string) (*model.UserSetting, error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	var setting model.UserSetting
	if err := db.GetDB(ctx).Where(strings.Join(where, " and "), args...).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (Dao) FindUserSettings(ctx context.Context, where []string, args []any, fields ...string) (list []*model.UserSetting, err error) {
	if len(where) == 0 {
		where, args = []string{"1"}, []any{}
	}
	err = db.GetDB(ctx).Where(strings.Join(where, " and "), args...).Find(&list).Error
	return
}

func (dao *Dao) UpsertUserSetting(ctx context.Context, m *model.UserSetting) (err error) {
	err = db.GetDB(ctx).
		Where("user_id=? and key=?", m.UserId, m.Key).
		Assign(model.UserSetting{Value: m.Value}).
		FirstOrCreate(&m).Error
	return
}

func (dao *Dao) FindUserAccessTokens(ctx context.Context, userId int32) (tokens []*model.AccessToken, err error) {
	var setting model.UserSetting
	err = db.GetDB(ctx).Where("user_id=? and key=?", userId, model.UserSettingKeyAccessToken).First(&setting).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	} else if err != nil {
		return
	}
	if setting.Value != "" {
		var val model.UserSettingValue
		if err = json.Unmarshal([]byte(setting.Value), &val); err != nil {
			return
		}
		tokens = val.AccessTokens
	}
	return
}
