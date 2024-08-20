package v1

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/api"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type MemoService struct {
	api.BaseService
	v1pb.UnimplementedInboxServiceServer
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (response *v1pb.ListMemosResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}

	slog.Info("user", user)
	return
}
