package model

import "time"

type SignInRequest struct {
	Username    string
	Password    string
	Audience    string
	KeyID       string
	NeverExpire bool
}

type SignInResponse struct {
	AccessToken string
	ExpireTime  time.Time
}

type SignOutRequest struct {
	UserId      int32
	AccessToken string
}
