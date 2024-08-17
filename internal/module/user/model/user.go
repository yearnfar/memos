package model

type User struct {
	ID int

	// Standard fields
	RowStatus RowStatus
	CreatedTs int64 `gorm:"autoCreateTime"`
	UpdatedTs int64 `gorm:"autoUpdateTime"`

	// Domain specific fields
	Username     string
	Role         Role
	Email        string
	Nickname     string
	PasswordHash string
	AvatarURL    string
	Description  string
}

func (User) TableName() string {
	return TableUser
}

type GetUserRequest struct {
	Id       int
	Username string
}

type CreateUserRequest struct {
	Username string
	Role     Role
	Email    string
	Nickname string
	Password string
}

type UpdateUserRequest struct {
	UpdateMasks []string
	UserId      int
	Username    string
	Role        Role
	RowStatus   RowStatus
	Email       string
	AvatarURL   string
	Nickname    string
	Password    string
	Description string
}

type SignUpRequest struct {
	Username string
	Password string
}
type ListUsersRequest struct {
	Role Role
}
