package dao

import "sync"

type Dao struct {
	workspaceSettingCache sync.Map // map[string]*storepb.WorkspaceSetting
	userCache             sync.Map // map[int]*User
	userSettingCache      sync.Map // map[string]*storepb.UserSetting
	idpCache              sync.Map // map[int]*storepb.IdentityProvider
}

func New() *Dao {
	return &Dao{}
}
