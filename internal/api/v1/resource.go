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

	var memoId int32
	if request.Resource.Memo != nil {
		memoId, err = api.ExtractMemoIDFromName(*request.Resource.Memo)
		if err != nil {
			return
		}
	}
	resource, err := memomod.CreateResource(ctx, &model.CreateResourceRequest{
		UserId:       user.ID,
		Name:         request.Resource.Name,
		Uid:          request.Resource.Uid,
		CreateTime:   request.Resource.CreateTime.AsTime().Unix(),
		Filename:     request.Resource.Filename,
		Content:      request.Resource.Content,
		ExternalLink: request.Resource.ExternalLink,
		Type:         request.Resource.Type,
		Size:         request.Resource.Size,
		MemoID:       memoId,
	})
	if err != nil {
		return
	}
	return convertResourceFromStore(ctx, resource), nil
}
