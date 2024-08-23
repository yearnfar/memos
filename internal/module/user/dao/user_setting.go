package dao

import (
	"context"
	"encoding/json"

	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
	"gorm.io/gorm"
)

func (Dao) GetUserSetting(ctx context.Context, req *model.FindUserSettingRequest) (setting *model.UserSetting, err error) {
	conn := db.GetDB(ctx)
	if req.Id != 0 {
		conn = conn.Where("id=?", req.Id)
	}
	setting = &model.UserSetting{}
	err = conn.First(&setting).Error
	return
}

func (Dao) FindUserSettings(ctx context.Context, req *model.FindUserSettingsRequest) (list []*model.UserSetting, err error) {
	conn := db.GetDB(ctx)
	if req.UserId != 0 {
		conn = conn.Where("user_id=?", req.UserId)
	}
	err = conn.Find(&list).Error
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
		err = json.Unmarshal([]byte(setting.Value), &tokens)
		if err != nil {
			return
		}
	}
	return
}
