package api

import (
	"github.com/yearnfar/memos/internal/module/auth"
	authSvc "github.com/yearnfar/memos/internal/module/auth/service"
	"github.com/yearnfar/memos/internal/module/memo"
	memoSvc "github.com/yearnfar/memos/internal/module/memo/service"
	"github.com/yearnfar/memos/internal/module/user"
	userSvc "github.com/yearnfar/memos/internal/module/user/service"
)

func Init() {
	user.Register(userSvc.Default())
	auth.Register(authSvc.Default())
	memo.Register(memoSvc.Default())
}
