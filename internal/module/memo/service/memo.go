package service

import (
	"context"
	"encoding/json"
	"fmt"
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

func (s *Service) CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (memoInfo *model.MemoInfo, err error) {
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
	memo := &model.Memo{
		UID:        shortuuid.New(),
		CreatorID:  int32(req.UserId),
		Content:    req.Content,
		Visibility: req.Visibility,
		RowStatus:  model.Normal,
		Payload: &model.MemoPayload{
			Property: property,
		},
	}
	if err = s.dao.CreateMemo(ctx, memo); err != nil {
		return
	}
	memoInfo, err = s.dao.FindMemoByID(ctx, memo.ID)
	return
}

func (s *Service) CreateMemoComment(ctx context.Context, req *model.CreateMemoCommentRequest) (memo *model.MemoInfo, err error) {
	relatedMemo, err := s.dao.FindMemoByID(ctx, req.ID)
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

func (s *Service) UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (memoInfo *model.MemoInfo, err error) {
	memoInfo, err = s.dao.FindMemoByID(ctx, req.ID)
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
			property, err := getMemoPropertyFromContent(req.Content)
			if err != nil {
				return nil, errors.Errorf("failed to get memo property: %v", err)
			}
			update["content"] = req.Content
			payload, _ := json.Marshal(&model.MemoPayload{Property: property})
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
			if err := s.dao.UpsertMemoOrganizer(ctx, &model.MemoOrganizer{
				MemoID: memoInfo.ID,
				UserID: req.UserId,
				Pinned: req.Pinned,
			}); err != nil {
				return nil, errors.New("failed to upsert memo organizer")
			}
		}
	}
	if len(update) > 0 {
		err = s.dao.UpdateMemo(ctx, &memoInfo.Memo, update)
		if err != nil {
			return
		}
	}
	return s.dao.FindMemoByID(ctx, req.ID)
}

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (list []*model.MemoInfo, err error) {
	where := []string{"creator_id=?"}
	args := []any{req.CreatorID}
	if req.ExcludeComments {
		where = append(where, "mr.memo_id is null")
	}
	if v := req.PayloadFind; v != nil {
		if v.Raw != "" {
			where, args = append(where, "`memo`.`payload` = ?"), append(args, v.Raw)
		}
		if len(v.TagSearch) != 0 {
			for _, tag := range v.TagSearch {
				where, args = append(where, "JSON_CONTAINS(JSON_EXTRACT(`memo`.`payload`, '$.property.tags'), ?)"), append(args, fmt.Sprintf(`"%s"`, tag))
			}
		}
		if v.HasLink {
			where = append(where, "JSON_EXTRACT(`memo`.`payload`, '$.property.hasLink') IS TRUE")
		}
		if v.HasTaskList {
			where = append(where, "JSON_EXTRACT(`memo`.`payload`, '$.property.hasTaskList') IS TRUE")
		}
		if v.HasCode {
			where = append(where, "JSON_EXTRACT(`memo`.`payload`, '$.property.hasCode') IS TRUE")
		}
		if v.HasIncompleteTasks {
			where = append(where, "JSON_EXTRACT(`memo`.`payload`, '$.property.hasIncompleteTasks') IS TRUE")
		}
	}
	list, err = s.dao.FindMemos(ctx, where, args)
	if err != nil {
		return
	}
	return
}

func (s *Service) GetMemo(ctx context.Context, req *model.GetMemoRequest) (*model.MemoInfo, error) {
	memo, err := s.dao.FindMemo(ctx, []string{"id=?", "uid=?"}, []any{req.Id, req.UID})
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
	memo, err := s.dao.FindMemoByID(ctx, req.Id)
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
	err = s.dao.DeleteMemoRelations(ctx, []string{"memo_id=?"}, []any{req.Id})
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
	resources, err := s.dao.FindResources(ctx, []string{"memo_id=?"}, []any{req.MemoID})
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
		res, err = s.dao.FindResourceByID(ctx, resource.ID)
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
