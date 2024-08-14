package model

// RowStatus is the status for a row.
type RowStatus string

const (
	// Normal is the status for a normal row.
	Normal RowStatus = "NORMAL"
	// Archived is the status for an archived row.
	Archived RowStatus = "ARCHIVED"
)

func (r RowStatus) String() string {
	return string(r)
}

// Role is the type of a role.
type Role string

const (
	// RoleHost is the HOST role.
	RoleHost Role = "HOST"
	// RoleAdmin is the ADMIN role.
	RoleAdmin Role = "ADMIN"
	// RoleUser is the USER role.
	RoleUser Role = "USER"
)

func (r Role) String() string {
	return string(r)
}

type UserSettingKey string

const (
	UserSettingKeyUnspecified    UserSettingKey = "USER_SETTING_KEY_UNSPECIFIED"
	UserSettingKeyAccessToken    UserSettingKey = "ACCESS_TOKENS"
	UserSettingKeyLocale         UserSettingKey = "LOCALE"
	UserSettingKeyAppearance     UserSettingKey = "APPEARANCE"
	UserSettingKeyMemoVisibility UserSettingKey = "MEMO_VISIBILITY"
)

const (
	TableUser        = "user"
	TableUserSetting = "user_setting"
)
