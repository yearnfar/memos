package v1

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type InboxService struct {
	api.BaseService
	v1pb.UnimplementedInboxServiceServer
}

func (s *InboxService) ListInboxes(ctx context.Context, req *v1pb.ListInboxesRequest) (response *v1pb.ListInboxesResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}
	inboxes, err := memomod.ListInboxes(ctx, &model.ListInboxesRequest{ReceiverId: user.ID})
	if err != nil {
		err = errors.Errorf("failed to list inbox: %v", err)
		return
	}
	response = &v1pb.ListInboxesResponse{
		Inboxes: []*v1pb.Inbox{},
	}
	for _, inbox := range inboxes {
		inboxMessage := convertInboxFromStore(inbox)
		if inboxMessage.Type == v1pb.Inbox_TYPE_UNSPECIFIED {
			continue
		}
		response.Inboxes = append(response.Inboxes, inboxMessage)
	}
	return
}
