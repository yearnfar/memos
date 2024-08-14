package service

import (
	"github.com/yearnfar/memos/internal/module/auth"
	"github.com/yearnfar/memos/internal/module/memo/dao"
)

type Service struct {
	dao auth.DAO
}

func Default() *Service {
	return New(dao.New())
}

func New(d auth.DAO) *Service {
	return &Service{dao: d}
}
