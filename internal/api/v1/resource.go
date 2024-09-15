package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

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
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
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

	return &httpbody.HttpBody{
		ContentType: contentType,
		Data:        rb.Blob,
	}, nil
}

func (s *ResourceService) ListResources(ctx context.Context, _ *v1pb.ListResourcesRequest) (response *v1pb.ListResourcesResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	list, err := memomod.ListResources(ctx, &model.ListResourcesRequest{CreatorID: user.ID})
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to list resources: %v", err)
		return
	}

	var resources []*v1pb.Resource
	for _, item := range list {
		resources = append(resources, convertResourceFromStore(ctx, item))
	}
	response = &v1pb.ListResourcesResponse{Resources: resources}
	return
}

func (s *ResourceService) DeleteResource(ctx context.Context, request *v1pb.DeleteResourceRequest) (*emptypb.Empty, error) {
	id, err := api.ExtractResourceIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid resource id: %v", err)
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get current user: %v", err)
	}
	// Delete the resource from the database.
	if err := memomod.DeleteResource(ctx, &model.DeleteResourceRequest{
		ID:     id,
		UserID: user.ID,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete resource: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func convertResourceFromStore(ctx context.Context, resource *model.Resource) *v1pb.Resource {
	resourceMessage := &v1pb.Resource{
		Name:       fmt.Sprintf("%s%d", api.ResourceNamePrefix, resource.ID),
		Uid:        resource.UID,
		CreateTime: timestamppb.New(time.Unix(resource.CreatedTs, 0)),
		Filename:   resource.Filename,
		Type:       resource.Type,
		Size:       resource.Size,
	}
	if resource.StorageType == model.ResourceStorageTypeExternal ||
		resource.StorageType == model.ResourceStorageTypeS3 {
		resourceMessage.ExternalLink = resource.Reference
	}
	if resource.MemoID != 0 {
		memo, _ := memomod.GetMemo(ctx, &model.GetMemoRequest{Id: resource.MemoID})
		if memo != nil {
			memoName := fmt.Sprintf("%s%d", api.MemoNamePrefix, memo.ID)
			resourceMessage.Memo = &memoName
		}
	}
	return resourceMessage
}
