package v1

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	memomodel "github.com/yearnfar/memos/internal/module/memo/model"

	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type MemoService struct {
	api.BaseService
	v1pb.UnimplementedMemoServiceServer
}

func (s *MemoService) CreateMemo(ctx context.Context, request *v1pb.CreateMemoRequest) (response *v1pb.Memo, err error) {
	return
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (response *v1pb.ListMemosResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}

	list, err := memomod.ListMemos(ctx, &memomodel.ListMemosRequest{CreatorId: user.ID})
	if err != nil {
		return
	}

	slog.Info("list", list)
	return
}

func (s *MemoService) ListMemoTags(ctx context.Context, request *v1pb.ListMemoTagsRequest) (response *v1pb.ListMemoTagsResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}

	slog.Info("user", user)
	return
}

func (s *MemoService) ListMemoProperties(ctx context.Context, request *v1pb.ListMemoPropertiesRequest) (response *v1pb.ListMemoPropertiesResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}

	slog.Info("user", user)
	return
}
