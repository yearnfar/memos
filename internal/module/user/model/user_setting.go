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

type UserSettingValue struct {
	AccessTokens   []*AccessToken `json:"access_tokens,omitempty"`
	Locale         string         `json:"locale,omitempty"`
	Appearance     string         `json:"appearance,omitempty"`
	MemoVisibility string         `json:"memo_visibility,omitempty"`
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
	Audience    string
	KeyID       string
}

type UpdateUserSettingRequest struct {
	UpdateMasks    []string
	UserID         int32
	Locale         string
	Appearance     string
	MemoVisibility string
}
