package service

import (
	"sync"

	"github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/dao"
)

type Service struct {
	dao memo.DAO

	workspaceSettingCache sync.Map // map[string]*storepb.WorkspaceSetting
	userCache             sync.Map // map[int]*User
	userSettingCache      sync.Map // map[string]*storepb.UserSetting
	idpCache              sync.Map // map[int]*storepb.IdentityProvider
}

func Default() *Service {
	return New(dao.New())
}

func New(dao memo.DAO) *Service {
	return &Service{
		dao: dao,
	}
}
