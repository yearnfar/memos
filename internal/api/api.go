package api

import (
	"github.com/yearnfar/memos/internal/module/user"
	userSvc "github.com/yearnfar/memos/internal/module/user/service"
)

func Init() {
	user.Register(userSvc.Default())
}
