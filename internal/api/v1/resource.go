package v1

import (
	"context"
	"strings"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
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

func (s *ResourceService) GetResourceBinary(ctx context.Context, request *v1pb.GetResourceBinaryRequest) (*httpbody.HttpBody, error) {
	id, err := api.ExtractResourceIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid resource id: %v", err)
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
	}

	rb, err := memomod.GetResourceBinary(ctx, &model.GetResourceBinaryRequest{UserId: user.ID, Id: id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get resource: %v", err)
	}

	contentType := rb.Resource.Type
	if strings.HasPrefix(contentType, "text/") {
		contentType += "; charset=utf-8"
	}

	httpBody := &httpbody.HttpBody{
		ContentType: contentType,
		Data:        rb.Blob,
	}
	return httpBody, nil
}
