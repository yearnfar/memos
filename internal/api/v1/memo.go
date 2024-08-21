package v1

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/pkg/errors"
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
	// listMemoRelationsResponse, err := s.ListMemoRelations(ctx, &v1pb.ListMemoRelationsRequest{Name: name})
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to list memo relations")
	// }

	// listMemoResourcesResponse, err := s.ListMemoResources(ctx, &v1pb.ListMemoResourcesRequest{Name: name})
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to list memo resources")
	// }

	// listMemoReactionsResponse, err := s.ListMemoReactions(ctx, &v1pb.ListMemoReactionsRequest{Name: name})
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to list memo reactions")
	// }

	// nodes, err := parser.Parse(tokenizer.Tokenize(memo.Content))
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to parse content")
	// }

	// snippet, err := getMemoContentSnippet(memo.Content)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to get memo content snippet")
	// }

	memoMessage := &v1pb.Memo{
		Name: name,
		Uid:  memo.UID,
		// RowStatus:   convertRowStatusFromStore(memo.RowStatus),
		Creator:     fmt.Sprintf("%s%d", api.UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(time.Unix(memo.CreatedTs, 0)),
		UpdateTime:  timestamppb.New(time.Unix(memo.UpdatedTs, 0)),
		DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:     memo.Content,
		// Snippet:     snippet,
		// Nodes:       convertFromASTNodes(nodes),
		Visibility: s.convertVisibilityFromStore(memo.Visibility),
		// Pinned:     memo.Pinned,
		// Relations:  listMemoRelationsResponse.Relations,
		// Resources:  listMemoResourcesResponse.Resources,
		// Reactions:  listMemoReactionsResponse.Reactions,
	}
	if memo.Payload != nil {
		memoMessage.Property = s.convertMemoPropertyFromStore(memo.Payload.Property)
	}
	// if memo.ParentID != nil {
	// 	parent := fmt.Sprintf("%s%d", api.MemoNamePrefix, *memo.ParentID)
	// 	memoMessage.Parent = &parent
	// }
	return memoMessage, nil
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
