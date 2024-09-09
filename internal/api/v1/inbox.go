package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

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
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
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

func convertInboxFromStore(inbox *model.Inbox) *v1pb.Inbox {
	return &v1pb.Inbox{
		Name:       fmt.Sprintf("%s%d", api.InboxNamePrefix, inbox.ID),
		Sender:     fmt.Sprintf("%s%d", api.UserNamePrefix, inbox.SenderID),
		Receiver:   fmt.Sprintf("%s%d", api.UserNamePrefix, inbox.ReceiverID),
		Status:     convertInboxStatusFromStore(inbox.Status),
		CreateTime: timestamppb.New(time.Unix(inbox.CreatedTs, 0)),
		Type:       v1pb.Inbox_Type(v1pb.Inbox_Type_value[string(inbox.Message.Type)]),
		ActivityId: &inbox.Message.ActivityId,
	}
}

func convertInboxStatusFromStore(status model.InboxStatus) v1pb.Inbox_Status {
	switch status {
	case model.InboxStatusUnread:
		return v1pb.Inbox_UNREAD
	case model.InboxStatusArchived:
		return v1pb.Inbox_ARCHIVED
	default:
		return v1pb.Inbox_STATUS_UNSPECIFIED
	}
}

func convertInboxStatusToStore(status v1pb.Inbox_Status) model.InboxStatus {
	switch status {
	case v1pb.Inbox_UNREAD:
		return model.InboxStatusUnread
	case v1pb.Inbox_ARCHIVED:
		return model.InboxStatusArchived
	default:
		return model.InboxStatusUnread
	}
}
