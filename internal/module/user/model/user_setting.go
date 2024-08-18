package model

type UserSetting struct {
	UserId int            `json:"user_id"`
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
