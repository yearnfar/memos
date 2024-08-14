package model

type UserSetting struct {
	UserId int
	Key    UserSettingKey
	Value  string
}

func (UserSetting) TableName() string {
	return TableUserSetting
}

type AccessToken struct {
	Token       string
	Description string
}
