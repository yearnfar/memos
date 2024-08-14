package service

import (
	mod "github.com/yearnfar/memos/internal/module/user"
	"github.com/yearnfar/memos/internal/module/user/dao"
)

type Service struct {
	dao mod.DAO
}

func Default() *Service {
	return New(dao.New())
}

func New(dao mod.DAO) *Service {
	return &Service{
		dao: dao,
	}
}
