package service

import (
	"context"
	"slices"

	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"

	"github.com/yearnfar/memos/internal/module/memo/model"
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

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (list []*model.Memo, err error) {
	return s.dao.FindMemos(ctx, &model.FindMemosRequest{})
}