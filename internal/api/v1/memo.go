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
	usermodel "github.com/yearnfar/memos/internal/module/user/model"
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
		relationsList = append(relationsList, convertMemoRelationFromStore(relation))
	}

	resources, err := memomod.ListResources(ctx, &model.ListResourcesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list memo relations")
	}
	var resourcesList []*v1pb.Resource
	for _, resource := range resources {
		resourcesList = append(resourcesList, convertResourceFromStore(ctx, resource))
	}

	reactions, err := memomod.ListReactions(ctx, &model.ListReactionsRequest{Id: 1})
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
		RowStatus:   convertRowStatusFromStore(usermodel.RowStatus(memo.RowStatus)),
		Creator:     fmt.Sprintf("%s%d", api.UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(time.Unix(memo.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(memo.UpdatedTs, 0)),
		DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:     memo.Content,
		Snippet:     snippet,
		// Nodes:       convertFromASTNodes(nodes),
		Visibility: convertVisibilityFromStore(memo.Visibility),
		// Pinned:     memo.Pinned,
		Relations: relationsList,
		Resources: resourcesList,
		Reactions: reactionList,
	}
	if memo.Payload != nil {
		memoMessage.Property = convertMemoPropertyFromStore(memo.Payload.Property)
	}
	// if memo.ParentID != 0 {
	// 	parent := fmt.Sprintf("%s%d", api.MemoNamePrefix, *memo.ParentID)
	// 	memoMessage.Parent = &parent
	// }
	return memoMessage, nil
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (response *v1pb.ListMemosResponse, err error) {
	user, err := s.GetCurrentUser(ctx)
	if err != nil {
		err = errors.Errorf("failed to get current user: %v", err)
		return
	}

	list, err := memomod.ListMemos(ctx, &model.ListMemosRequest{CreatorId: user.ID})
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
