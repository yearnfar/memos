package v1

import (
	"context"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type ResourceService struct {
	api.BaseService
	v1pb.UnimplementedResourceServiceServer
}

func (s *ResourceService) CreateResource(ctx context.Context, request *v1pb.CreateResourceRequest) (response *v1pb.Resource, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
	}

	resource, err := memomod.CreateResource(ctx, &model.CreateResourceRequest{
		UserId: user.ID,
	})
	if err != nil {
		return
	}

	return convertResourceFromStore(ctx, resource), nil
}
