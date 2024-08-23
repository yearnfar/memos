package v1

import (
	"context"
	"fmt"
	"time"

	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	memomodel "github.com/yearnfar/memos/internal/module/memo/model"
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
)

func convertUserFromStore(user *usermodel.User) *v1pb.User {
	userpb := &v1pb.User{
		Name:        fmt.Sprintf("%s%d", api.UserNamePrefix, user.ID),
		Id:          int32(user.ID),
		RowStatus:   convertRowStatusFromStore(user.RowStatus),
		CreateTime:  timestamppb.New(time.Unix(user.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(user.UpdatedTs, 0)),
		Role:        convertUserRoleFromStore(user.Role),
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		AvatarUrl:   user.AvatarURL,
		Description: user.Description,
	}
	// Use the avatar URL instead of raw base64 image data to reduce the response size.
	if user.AvatarURL != "" {
		userpb.AvatarUrl = fmt.Sprintf("/file/%s/avatar", userpb.Name)
	}
	return userpb
}

func convertRowStatusFromStore(rowStatus usermodel.RowStatus) v1pb.RowStatus {
	switch rowStatus {
	case usermodel.Normal:
		return v1pb.RowStatus_ACTIVE
	case usermodel.Archived:
		return v1pb.RowStatus_ARCHIVED
	default:
		return v1pb.RowStatus_ROW_STATUS_UNSPECIFIED
	}
}

func convertUserRoleFromStore(role usermodel.Role) v1pb.User_Role {
	switch role {
	case usermodel.RoleHost:
		return v1pb.User_HOST
	case usermodel.RoleAdmin:
		return v1pb.User_ADMIN
	case usermodel.RoleUser:
		return v1pb.User_USER
	default:
		return v1pb.User_ROLE_UNSPECIFIED
	}
}

func convertInboxFromStore(inbox *memomodel.Inbox) *v1pb.Inbox {
	return &v1pb.Inbox{
		Name:       fmt.Sprintf("%s%d", api.InboxNamePrefix, inbox.ID),
		Sender:     fmt.Sprintf("%s%d", api.UserNamePrefix, inbox.SenderID),
		Receiver:   fmt.Sprintf("%s%d", api.UserNamePrefix, inbox.ReceiverID),
		Status:     convertInboxStatusFromStore(inbox.Status),
		CreateTime: timestamppb.New(time.Unix(inbox.CreatedTs, 0)),
		Type:       v1pb.Inbox_Type(inbox.Message.Type),
		ActivityId: &inbox.Message.ActivityId,
	}
}

func convertInboxStatusFromStore(status memomodel.InboxStatus) v1pb.Inbox_Status {
	switch status {
	case memomodel.InboxStatusUnread:
		return v1pb.Inbox_UNREAD
	case memomodel.InboxStatusArchived:
		return v1pb.Inbox_ARCHIVED
	default:
		return v1pb.Inbox_STATUS_UNSPECIFIED
	}
}

func convertInboxStatusToStore(status v1pb.Inbox_Status) memomodel.InboxStatus {
	switch status {
	case v1pb.Inbox_UNREAD:
		return memomodel.InboxStatusUnread
	case v1pb.Inbox_ARCHIVED:
		return memomodel.InboxStatusArchived
	default:
		return memomodel.InboxStatusUnread
	}
}

func convertMemoRelationFromStore(memoRelation *model.MemoRelation) *v1pb.MemoRelation {
	return &v1pb.MemoRelation{
		Memo:        fmt.Sprintf("%s%d", api.MemoNamePrefix, memoRelation.MemoID),
		RelatedMemo: fmt.Sprintf("%s%d", api.MemoNamePrefix, memoRelation.RelatedMemoID),
		Type:        convertMemoRelationTypeFromStore(memoRelation.Type),
	}
}

func convertMemoRelationTypeFromStore(relationType model.MemoRelationType) v1pb.MemoRelation_Type {
	switch relationType {
	case model.MemoRelationReference:
		return v1pb.MemoRelation_REFERENCE
	case model.MemoRelationComment:
		return v1pb.MemoRelation_COMMENT
	default:
		return v1pb.MemoRelation_TYPE_UNSPECIFIED
	}
}

func convertVisibilityFromStore(visibility model.Visibility) v1pb.Visibility {
	switch visibility {
	case model.Private:
		return v1pb.Visibility_PRIVATE
	case model.Protected:
		return v1pb.Visibility_PROTECTED
	case model.Public:
		return v1pb.Visibility_PUBLIC
	default:
		return v1pb.Visibility_VISIBILITY_UNSPECIFIED
	}
}

func convertMemoPropertyFromStore(property *model.MemoPayloadProperty) *v1pb.MemoProperty {
	if property == nil {
		return nil
	}
	return &v1pb.MemoProperty{
		Tags:               property.Tags,
		HasLink:            property.HasLink,
		HasTaskList:        property.HasTaskList,
		HasCode:            property.HasCode,
		HasIncompleteTasks: property.HasIncompleteTasks,
	}
}

func convertReactionFromStore(ctx context.Context, reaction *model.Reaction) (*v1pb.Reaction, error) {
	// creator, err := s.Store.GetUser(ctx, &model.FindUser{
	// 	ID: &reaction.CreatorID,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return &v1pb.Reaction{
		Id:           reaction.ID,
		Creator:      fmt.Sprintf("%s%d", api.UserNamePrefix, reaction.CreatorID),
		ContentId:    reaction.ContentID,
		ReactionType: v1pb.Reaction_Type(v1pb.Reaction_Type_value[string(reaction.ReactionType)]),
	}, nil
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
	if resource.StorageType == model.ResourceStorageTypeExternal || resource.StorageType == model.ResourceStorageTypeS3 {
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
