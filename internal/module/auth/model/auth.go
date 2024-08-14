package model

import "time"

type SignInRequest struct {
	Username    string
	Password    string
	NeverExpire bool
}

type SignInResponse struct {
	AccessToken string
	ExpireTime  time.Time
}

type SignOutRequest struct {
	UserId      int
	AccessToken string
}
