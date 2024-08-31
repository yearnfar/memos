package service

import (
	"context"
	"slices"
	"time"

	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"

	"github.com/yearnfar/memos/internal/module/memo/model"
	"github.com/yearnfar/memos/internal/pkg/util"
)

func (s *Service) CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (memo *model.Memo, err error) {
	memoRelatedSetting, err := s.getWorkspaceMemoRelatedSetting(ctx)
	if err != nil {
		err = errors.Errorf("failed to get workspace memo related setting")
		return
	}
	if memoRelatedSetting.DisallowPublicVisibility && req.Visibility == model.Public {
		err = errors.Errorf("disable public memos system setting is enabled")
		return
	}
	if len(req.Content) > int(memoRelatedSetting.ContentLengthLimit) {
		err = errors.Errorf("content too long (max %d characters)", memoRelatedSetting.ContentLengthLimit)
		return
	}
	property, err := getMemoPropertyFromContent(req.Content)
	if err != nil {
		err = errors.Errorf("failed to get memo property: %v", err)
		return
	}
	memo = &model.Memo{
		UID:        shortuuid.New(),
		CreatorID:  int32(req.UserId),
		Content:    req.Content,
		Visibility: req.Visibility,
		RowStatus:  model.Normal,
		Payload: &model.MemoPayload{
			Property: property,
		},
	}
	err = s.dao.CreateMemo(ctx, memo)
	return
}

func (s *Service) CreateMemoComment(ctx context.Context, req *model.CreateMemoCommentRequest) (memo *model.Memo, err error) {
	relatedMemo, err := s.dao.FindMemo(ctx, &model.FindMemoRequest{Id: req.ID})
	if err != nil {
		err = errors.New("failed to get memo")
		return
	}
	memo, err = s.CreateMemo(ctx, req.Comment)
	if err != nil {
		err = errors.New("failed to create memo")
		return
	}
	err = s.dao.UpsertMemoRelation(ctx, &model.MemoRelation{
		MemoID:        memo.ID,
		RelatedMemoID: relatedMemo.ID,
		Type:          model.MemoRelationComment,
	})
	if err != nil {
		err = errors.New("failed to create memo relation")
		return
	}

	if memo.Visibility != model.Private && memo.CreatorID != relatedMemo.CreatorID {
		activity := &model.Activity{
			CreatorID: memo.CreatorID,
			Type:      model.ActivityTypeMemoComment,
			Level:     model.ActivityLevelInfo,
			Payload: &model.ActivityPayload{
				MemoComment: &model.ActivityMemoCommentPayload{
					MemoId:        memo.CreatorID,
					RelatedMemoId: relatedMemo.ID,
				},
			},
		}
		if err = s.dao.CreateActivity(ctx, activity); err != nil {
			err = errors.New("failed to create activity")
			return
		}
		err = s.dao.CreateInbox(ctx, &model.Inbox{
			SenderID:   memo.CreatorID,
			ReceiverID: relatedMemo.CreatorID,
			Status:     model.InboxStatusUnread,
			Message: &model.InboxMessage{
				Type:       model.InboxMsgTypeMemoComment,
				ActivityId: activity.ID,
			},
		})
		if err != nil {
			err = errors.New("failed to create inbox")
			return
		}
	}
	return
}

func getMemoPropertyFromContent(content string) (*model.MemoPayloadProperty, error) {
	nodes, err := parser.Parse(tokenizer.Tokenize(content))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse content")
	}

	property := &model.MemoPayloadProperty{}
	TraverseASTNodes(nodes, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.Tag:
			tag := n.Content
			if !slices.Contains(property.Tags, tag) {
				property.Tags = append(property.Tags, tag)
			}
		case *ast.Link, *ast.AutoLink:
			property.HasLink = true
		case *ast.TaskList:
			property.HasTaskList = true
			if !n.Complete {
				property.HasIncompleteTasks = true
			}
		case *ast.Code, *ast.CodeBlock:
			property.HasCode = true
		}
	})
	return property, nil
}

func TraverseASTNodes(nodes []ast.Node, fn func(ast.Node)) {
	for _, node := range nodes {
		fn(node)
		switch n := node.(type) {
		case *ast.Paragraph:
			TraverseASTNodes(n.Children, fn)
		case *ast.Heading:
			TraverseASTNodes(n.Children, fn)
		case *ast.Blockquote:
			TraverseASTNodes(n.Children, fn)
		case *ast.OrderedList:
			TraverseASTNodes(n.Children, fn)
		case *ast.UnorderedList:
			TraverseASTNodes(n.Children, fn)
		case *ast.TaskList:
			TraverseASTNodes(n.Children, fn)
		case *ast.Bold:
			TraverseASTNodes(n.Children, fn)
		}
	}
}

func (s *Service) UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (memo *model.Memo, err error) {
	memo, err = s.dao.FindMemo(ctx, &model.FindMemoRequest{Id: req.ID})
	if err != nil {
		return
	}
	memoRelatedSetting, err := s.getWorkspaceMemoRelatedSetting(ctx)
	if err != nil {
		return nil, errors.New("failed to get workspace memo related setting")
	}
	update := map[string]any{}
	for _, path := range req.UpdateMasks {
		if path == "content" {
			if len(req.Content) > int(memoRelatedSetting.ContentLengthLimit) {
				return nil, errors.Errorf("content too long (max %d characters)", memoRelatedSetting.ContentLengthLimit)
			}
			update["content"] = req.Content
			property, err := getMemoPropertyFromContent(req.Content)
			if err != nil {
				return nil, errors.Errorf("failed to get memo property: %v", err)
			}
			payload := memo.Payload
			payload.Property = property
			update["payload"] = payload
		} else if path == "uid" {
			if !util.UIDMatcher.MatchString(req.UID) {
				return nil, errors.New("invalid resource name")
			}
			update["mid"] = req.UID
		} else if path == "visibility" {
			if memoRelatedSetting.DisallowPublicVisibility && req.Visibility == model.Public {
				return nil, errors.New("disable public memos system setting is enabled")
			}
			update["visibility"] = req.Visibility
		} else if path == "row_status" {
			update["row_status"] = req.RowStatus
		} else if path == "create_time" {
			update["created_ts"] = req.CreatedTime
		} else if path == "display_time" {
			if memoRelatedSetting.DisplayWithUpdateTime {
				update["updated_ts"] = req.DisplayTime
			} else {
				update["created_ts"] = req.DisplayTime
			}
		} else if path == "pinned" {
			// if _, err := s.Store.UpsertMemoOrganizer(ctx, &store.MemoOrganizer{
			// 	MemoID: id,
			// 	UserID: user.ID,
			// 	Pinned: request.Memo.Pinned,
			// }); err != nil {
			// 	return nil, status.Errorf(codes.Internal, "failed to upsert memo organizer")
			// }
		}
	}
	err = s.dao.UpdateMemo(ctx, memo, update)
	return
}

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (list []*model.Memo, err error) {
	list, err = s.dao.FindMemos(ctx, &model.FindMemoRequest{
		CreatorId:       req.CreatorId,
		ExcludeComments: req.ExcludeComments,
	})
	if err != nil {
		return
	}

	return
}

func (s *Service) GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.Memo, error) {
	memo, err := s.dao.FindMemo(ctx, &model.FindMemoRequest{
		Id:  req.Id,
		UID: req.UID,
	})
	if err != nil {
		return nil, err
	}
	if req.CurrentUserId != 0 {
		if memo.Visibility == model.Private && memo.CreatorID != req.CurrentUserId {
			return nil, errors.New("permission denied")
		}
	}
	return memo, nil
}

func (s *Service) DeleteMemo(ctx context.Context, req *model.DeleteMemoRequest) (err error) {
	memo, err := s.dao.FindMemo(ctx, &model.FindMemoRequest{Id: req.Id})
	if err != nil {
		return
	}
	if memo.CreatorID != req.CurrentUserId {
		err = errors.New("permission denied")
		return
	}

	// TO-DO
	// 删除webhook

	if err = s.dao.DeleteMemoById(ctx, req.Id); err != nil {
		return
	}
	err = s.dao.DeleteMemoRelations(ctx, &model.DeleteMemoRelationsRequest{MemoID: req.Id})
	if err != nil {
		return
	}
	return
}

func (s *Service) UpsertReaction(ctx context.Context, req *model.UpsertReactionRequest) (reaction *model.Reaction, err error) {
	reaction = &model.Reaction{
		CreatorID:    req.CreatorID,
		ContentID:    req.ContentID,
		ReactionType: req.ReactionType,
	}
	err = s.dao.CreateReaction(ctx, reaction)
	return
}

func (s *Service) SetMemoResources(ctx context.Context, req *model.SetMemoResourcesRequest) (err error) {
	resources, err := s.dao.FindResources(ctx, &model.FindResourceRequest{
		MemoID: req.MemoID,
	})
	if err != nil {
		err = errors.New("failed to list resources")
		return
	}

	// Delete resources that are not in the request.
	for _, resource := range resources {
		found := false
		for _, requestResource := range req.Resources {
			if resource.UID == requestResource.Uid {
				found = true
				break
			}
		}
		if !found {
			if err = s.dao.DeleteResourceById(ctx, resource.ID); err != nil {
				err = errors.New("failed to delete resource")
				return
			}
		}
	}

	slices.Reverse(req.Resources)

	// Update resources' memo_id in the request.
	for index, resource := range req.Resources {
		var res *model.Resource
		res, err = s.dao.FindResource(ctx, &model.FindResourceRequest{ID: resource.ID})
		if err != nil {
			continue
		}
		if err = s.dao.UpdateResource(ctx, res, map[string]any{
			"memo_id":    req.MemoID,
			"updated_ts": time.Now().Unix() + int64(index),
		}); err != nil {
			err = errors.Errorf("failed to update resource: %v", err)
			return
		}
	}
	return
}
