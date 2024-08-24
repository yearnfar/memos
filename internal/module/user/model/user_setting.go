package model

import "time"

type UserSetting struct {
	UserId int32          `json:"user_id"`
	Key    UserSettingKey `json:"key"`
	Value  string         `json:"value"`
}

func (UserSetting) TableName() string {
	return TableUserSetting
}

type AccessToken struct {
	Token       string `json:"token"`
	Description string `json:"description"`
}

type FindUserSettingRequest struct {
	Id int32
}

type FindUserSettingsRequest struct {
	UserId int32
}

type CreateUserAccessTokenRequest struct {
	UserID      int32
	Description string
	ExpiresAt   time.Time
}
