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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yearnfar/memos/internal/api"
	memomod "github.com/yearnfar/memos/internal/module/memo"
	"github.com/yearnfar/memos/internal/module/memo/model"
	v1pb "github.com/yearnfar/memos/internal/proto/api/v1"
)

type MemoService struct {
	api.BaseService
	v1pb.UnimplementedMemoServiceServer
}

func (s *MemoService) CreateMemo(ctx context.Context, request *v1pb.CreateMemoRequest) (response *v1pb.Memo, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	memo, err := memomod.CreateMemo(ctx, &model.CreateMemoRequest{
		UserId:     user.ID,
		Content:    request.Content,
		Visibility: model.Visibility(request.Visibility.String()),
	})
	if err != nil {
		return
	}
	return s.convertMemoFromStore(ctx, memo)
}

func (s *MemoService) convertMemoFromStore(ctx context.Context, memo *model.MemoInfo) (*v1pb.Memo, error) {
	displayTs := memo.CreatedTs
	// workspaceMemoRelatedSetting, err := s.Store.GetWorkspaceMemoRelatedSetting(ctx)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to get workspace memo related setting")
	// }
	// if workspaceMemoRelatedSetting.DisplayWithUpdateTime {
	// 	displayTs = memo.UpdatedTs
	// }

	name := fmt.Sprintf("%s%d", api.MemoNamePrefix, memo.ID)
	relations, err := memomod.ListMemoRelations(ctx, &model.ListMemoRelationsRequest{MemoID: memo.ID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo relations")
	}
	var relationsList []*v1pb.MemoRelation
	for _, relation := range relations {
		relationsList = append(relationsList, convertMemoRelationFromStore(relation))
	}

	resources, err := memomod.ListResources(ctx, &model.ListResourcesRequest{MemoID: memo.ID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo relations")
	}
	var resourcesList []*v1pb.Resource
	for _, resource := range resources {
		resourcesList = append(resourcesList, convertResourceFromStore(ctx, resource))
	}

	reactions, err := memomod.ListReactions(ctx, &model.ListReactionsRequest{ContentId: name})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo reactions")
	}
	var reactionList []*v1pb.Reaction
	for _, reaction := range reactions {
		item, _ := convertReactionFromStore(ctx, reaction)
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
		Nodes:       convertFromASTNodes(nodes),
		Visibility:  convertVisibilityFromStore(memo.Visibility),
		Pinned:      memo.Pinned,
		Relations:   relationsList,
		Resources:   resourcesList,
		Reactions:   reactionList,
	}
	if memo.Payload != nil {
		memoMessage.Property = convertMemoPropertyFromStore(memo.Payload.Property)
	}
	if memo.ParentID != 0 {
		parent := fmt.Sprintf("%s%d", api.MemoNamePrefix, memo.ParentID)
		memoMessage.Parent = &parent
	}
	return memoMessage, nil
}

func (s *MemoService) SetMemoResources(ctx context.Context, request *v1pb.SetMemoResourcesRequest) (response *emptypb.Empty, err error) {
	memoID, err := api.ExtractMemoIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
	}
	var resources []*model.MemoResource
	for _, res := range request.Resources {
		resID, _ := api.ExtractResourceIDFromName(res.Name)
		resources = append(resources, &model.MemoResource{
			ID:           resID,
			Uid:          res.Uid,
			CreateTime:   res.CreateTime.AsTime().Unix(),
			Filename:     res.Filename,
			Content:      res.Content,
			ExternalLink: res.ExternalLink,
			Type:         res.Type,
			Size:         res.Size,
		})
	}
	err = memomod.SetMemoResources(ctx, &model.SetMemoResourcesRequest{
		MemoID:    memoID,
		Resources: resources,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "set memo resource fail: %v", err)
	}
	return
}

func (s *MemoService) SetMemoRelations(ctx context.Context, request *v1pb.SetMemoRelationsRequest) (response *emptypb.Empty, err error) {
	memoID, err := api.ExtractMemoIDFromName(request.Name)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
		return
	}
	var list []*model.MemoRelation
	for _, relation := range request.Relations {
		var relatedMemoID int32
		relatedMemoID, err = api.ExtractMemoIDFromName(relation.RelatedMemo)
		if err != nil {
			err = status.Errorf(codes.Internal, "set memo relations fail: %v", err)
			return
		}
		item := &model.MemoRelation{
			MemoID:        memoID,
			RelatedMemoID: relatedMemoID,
			Type:          convertMemoRelationTypeToStore(relation.Type),
		}
		list = append(list, item)
	}
	err = memomod.SetMemoRelations(ctx, &model.SetMemoRelationsRequest{
		MemoID:    memoID,
		Relations: list,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "set memo relations fail: %v", err)
		return
	}
	return
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (response *v1pb.ListMemosResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}

	memos, err := memomod.ListMemos(ctx, &model.ListMemosRequest{
		CreatorId:       user.ID,
		ExcludeComments: true,
	})
	if err != nil {
		return
	}

	var list []*v1pb.Memo
	for _, memo := range memos {
		var item *v1pb.Memo
		item, err = s.convertMemoFromStore(ctx, memo)
		if err != nil {
			return
		}
		list = append(list, item)
	}

	response = &v1pb.ListMemosResponse{
		Memos:         list,
		NextPageToken: "",
	}
	return
}

func (s *MemoService) GetMemo(ctx context.Context, request *v1pb.GetMemoRequest) (response *v1pb.Memo, err error) {
	id, err := api.ExtractMemoIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
	}
	memo, err := memomod.GetMemo(ctx, &model.GetMemoRequest{Id: id})
	if err != nil {
		err = status.Errorf(codes.Internal, "get memo failed: %v", err)
		return
	}
	response, err = s.convertMemoFromStore(ctx, memo)
	return
}

func (s *MemoService) GetMemoByUid(ctx context.Context, request *v1pb.GetMemoByUidRequest) (response *v1pb.Memo, err error) {
	memo, err := memomod.GetMemo(ctx, &model.GetMemoRequest{UID: request.Uid})
	if err != nil {
		err = status.Errorf(codes.Internal, "get memo failed: %v", err)
		return
	}
	response, err = s.convertMemoFromStore(ctx, memo)
	return
}

func (s *MemoService) UpsertMemoReaction(ctx context.Context, request *v1pb.UpsertMemoReactionRequest) (response *v1pb.Reaction, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	reaction, err := memomod.UpsertReaction(ctx, &model.UpsertReactionRequest{
		CreatorID:    user.ID,
		ContentID:    request.Reaction.ContentId,
		ReactionType: model.ReactionType(v1pb.Reaction_Type_name[int32(request.Reaction.ReactionType)]),
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to upsert reaction")
		return
	}

	response, err = convertReactionFromStore(ctx, reaction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert reaction")
	}
	return response, nil
}

func (s *MemoService) UpdateMemo(ctx context.Context, request *v1pb.UpdateMemoRequest) (response *v1pb.Memo, err error) {
	id, err := api.ExtractMemoIDFromName(request.Memo.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}

	memo, err := memomod.UpdateMemo(ctx, &model.UpdateMemoRequest{
		UpdateMasks: request.UpdateMask.Paths,
		UserId:      user.ID,
		ID:          id,
		UID:         request.Memo.Uid,
		Content:     request.Memo.Content,
		RowStatus:   s.convertRowStatusToStore(request.Memo.RowStatus),
		Visibility:  model.Visibility(request.Memo.Visibility.String()),
		Pinned:      request.Memo.Pinned,
		UpdatedTime: request.Memo.UpdateTime.AsTime().Unix(),
		CreatedTime: request.Memo.CreateTime.AsTime().Unix(),
		DisplayTime: request.Memo.DisplayTime.AsTime().Unix(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed update user: %v", err.Error())
	}
	response, err = s.convertMemoFromStore(ctx, memo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert memo")
	}
	return
}

func (s *MemoService) DeleteMemo(ctx context.Context, request *v1pb.DeleteMemoRequest) (response *emptypb.Empty, err error) {
	id, err := api.ExtractMemoIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}
	err = memomod.DeleteMemo(ctx, &model.DeleteMemoRequest{
		Id:            id,
		CurrentUserId: user.ID,
	})
	return
}

func (s *MemoService) CreateMemoComment(ctx context.Context, request *v1pb.CreateMemoCommentRequest) (response *v1pb.Memo, err error) {
	id, err := api.ExtractMemoIDFromName(request.Name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid memo name: %v", err)
	}
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user")
		return
	}

	req := &model.CreateMemoCommentRequest{
		ID: id,
		Comment: &model.CreateMemoRequest{
			UserId:     user.ID,
			Content:    request.Comment.Content,
			Visibility: model.Visibility(request.Comment.Visibility.String()),
		},
	}
	memo, err := memomod.CreateMemoComment(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create memo comment")
	}

	response, err = s.convertMemoFromStore(ctx, memo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert memo")
	}
	return
}

func (s *MemoService) ListMemoTags(ctx context.Context, request *v1pb.ListMemoTagsRequest) (response *v1pb.ListMemoTagsResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user: %v", err)
		return
	}

	slog.Info("user", user)
	return
}

func (s *MemoService) ListMemoProperties(ctx context.Context, request *v1pb.ListMemoPropertiesRequest) (response *v1pb.ListMemoPropertiesResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get current user")
		return
	}

	slog.Info("user", user)
	return
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

func (s *MemoService) convertRowStatusToStore(rowStatus v1pb.RowStatus) model.RowStatus {
	switch rowStatus {
	case v1pb.RowStatus_ACTIVE:
		return model.Normal
	case v1pb.RowStatus_ARCHIVED:
		return model.Archived
	default:
		return model.Normal
	}
}

func convertMemoRelationTypeToStore(relationType v1pb.MemoRelation_Type) model.MemoRelationType {
	switch relationType {
	case v1pb.MemoRelation_REFERENCE:
		return model.MemoRelationReference
	case v1pb.MemoRelation_COMMENT:
		return model.MemoRelationComment
	default:
		return model.MemoRelationReference
	}
}
