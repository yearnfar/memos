package service

import (
	"github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/dao"
)

type Service struct {
	dao memo.DAO
}

func Default() *Service {
	return New(dao.New())
}

func New(dao memo.DAO) *Service {
	return &Service{
		dao: dao,
	}
}
