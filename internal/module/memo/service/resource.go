package service

import (
	"context"

	"github.com/lithammer/shortuuid/v4"
	"github.com/yearnfar/memos/internal/module/memo/model"
)

func (s *Service) CreateResource(ctx context.Context, req *model.CreateResourceRequest) (resource *model.Resource, err error) {
	resource = &model.Resource{
		UID:       shortuuid.New(),
		CreatorID: req.UserId,
		Filename:  req.Filename,
		Type:      req.Type,
	}
	err = s.dao.CreateResource(ctx, resource)
	return
}

func (s *Service) ListResources(ctx context.Context, req *model.ListResourcesRequest) (list []*model.Resource, err error) {
	return
}
