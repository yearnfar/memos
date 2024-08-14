package dao

import (
	"context"
	"encoding/json"

	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (dao *Dao) UpsertUserSetting(ctx context.Context, m *model.UserSetting) (err error) {
	err = db.GetDB(ctx).
		Where("user_id=? and key=?", m.UserId, m.Key).
		Assign(model.UserSetting{Value: m.Value}).
		FirstOrCreate(&m).Error
	return
}

func (dao *Dao) GetUserAccessTokens(ctx context.Context, userId int) (tokens []*model.AccessToken, err error) {
	var setting model.UserSetting
	err = db.GetDB(ctx).Where("user_id=? and key=?", userId, model.UserSettingKeyAccessToken).First(&setting).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		return
	} else if err != nil {
		return
	}
	if setting.Value != "" {
		err = json.Unmarshal([]byte(setting.Value), &tokens)
		if err != nil {
			return
		}
	}
	return
}
