package v1

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/pkg/errors"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/renderer"
	"github.com/yearnfar/gokit/strutil"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	memomodel "github.com/yearnfar/memos/internal/module/memo/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type MemoService struct {
	api.BaseService
	v1pb.UnimplementedMemoServiceServer
}

func (s *MemoService) CreateMemo(ctx context.Context, request *v1pb.CreateMemoRequest) (response *v1pb.Memo, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}
	memo, err := memomod.CreateMemo(ctx, &memomodel.CreateMemoRequest{
		UserId:     user.ID,
		Content:    request.Content,
		Visibility: model.Visibility(request.Visibility.String()),
	})
	if err != nil {
		return
	}
	return s.convertMemoFromStore(ctx, memo)
}

func (s *MemoService) convertMemoFromStore(ctx context.Context, memo *model.Memo) (*v1pb.Memo, error) {
	displayTs := memo.CreatedTs
	// workspaceMemoRelatedSetting, err := s.Store.GetWorkspaceMemoRelatedSetting(ctx)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to get workspace memo related setting")
	// }
	// if workspaceMemoRelatedSetting.DisplayWithUpdateTime {
	// 	displayTs = memo.UpdatedTs
	// }

	name := fmt.Sprintf("%s%d", api.MemoNamePrefix, memo.ID)
	relations, err := memomod.ListMemoRelations(ctx, &model.ListMemoRelationsRequest{Id: 0})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo relations")
	}
	var relationsList []*v1pb.MemoRelation
	for _, relation := range relations {
		relationsList = append(relationsList, s.convertMemoRelationFromStore(relation))
	}

	resources, err := memomod.ListResources(ctx, &model.ListResourcesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo relations")
	}
	var resourcesList []*v1pb.Resource
	for _, resource := range resources {
		resourcesList = append(resourcesList, s.convertResourceFromStore(ctx, resource))
	}

	reactions, err := memomod.ListReactions(ctx, &model.ListReactionsRequest{Id: 1})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo reactions")
	}
	var reactionList []*v1pb.Reaction
	for _, reaction := range reactions {
		item, _ := s.convertReactionFromStore(ctx, reaction)
		reactionList = append(reactionList, item)
	}

	nodes, err := parser.Parse(tokenizer.Tokenize(memo.Content))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse content")
	}

	snippet := renderer.NewStringRenderer().Render(nodes)
	if strutil.Len(snippet) > 100 {
		snippet = strutil.Substr(snippet, 0, 100, "...")
	}

	memoMessage := &v1pb.Memo{
		Name:        name,
		Uid:         memo.UID,
		RowStatus:   s.convertRowStatusFromStore(memo.RowStatus),
		Creator:     fmt.Sprintf("%s%d", api.UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(time.Unix(memo.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(memo.UpdatedTs, 0)),
		DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:     memo.Content,
		Snippet:     snippet,
		// Nodes:       convertFromASTNodes(nodes),
		Visibility: s.convertVisibilityFromStore(memo.Visibility),
		// Pinned:     memo.Pinned,
		Relations: relationsList,
		Resources: resourcesList,
		Reactions: reactionList,
	}
	if memo.Payload != nil {
		memoMessage.Property = s.convertMemoPropertyFromStore(memo.Payload.Property)
	}
	// if memo.ParentID != 0 {
	// 	parent := fmt.Sprintf("%s%d", api.MemoNamePrefix, *memo.ParentID)
	// 	memoMessage.Parent = &parent
	// }
	return memoMessage, nil
}

func (s *MemoService) convertResourceFromStore(ctx context.Context, resource *model.Resource) *v1pb.Resource {
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

func (s *MemoService) convertReactionFromStore(ctx context.Context, reaction *model.Reaction) (*v1pb.Reaction, error) {
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

func (s *MemoService) convertMemoRelationFromStore(memoRelation *model.MemoRelation) *v1pb.MemoRelation {
	return &v1pb.MemoRelation{
		Memo:        fmt.Sprintf("%s%d", api.MemoNamePrefix, memoRelation.MemoID),
		RelatedMemo: fmt.Sprintf("%s%d", api.MemoNamePrefix, memoRelation.RelatedMemoID),
		Type:        s.convertMemoRelationTypeFromStore(memoRelation.Type),
	}
}

func (s *MemoService) convertMemoRelationTypeFromStore(relationType model.MemoRelationType) v1pb.MemoRelation_Type {
	switch relationType {
	case model.MemoRelationReference:
		return v1pb.MemoRelation_REFERENCE
	case model.MemoRelationComment:
		return v1pb.MemoRelation_COMMENT
	default:
		return v1pb.MemoRelation_TYPE_UNSPECIFIED
	}
}

func (s *MemoService) convertRowStatusFromStore(rowStatus model.RowStatus) v1pb.RowStatus {
	switch rowStatus {
	case model.Normal:
		return v1pb.RowStatus_ACTIVE
	case model.Archived:
		return v1pb.RowStatus_ARCHIVED
	default:
		return v1pb.RowStatus_ROW_STATUS_UNSPECIFIED
	}
}

func (s *MemoService) convertVisibilityFromStore(visibility model.Visibility) v1pb.Visibility {
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

func (s *MemoService) convertMemoPropertyFromStore(property *model.MemoPayloadProperty) *v1pb.MemoProperty {
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
