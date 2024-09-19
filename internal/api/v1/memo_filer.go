package v1

import (
	"github.com/google/cel-go/cel"
	"github.com/pkg/errors"
	"github.com/yearnfar/memos/internal/module/memo/model"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

// MemoFilterCELAttributes are the CEL attributes.
var MemoFilterCELAttributes = []cel.EnvOption{
	cel.Variable("content_search", cel.ListType(cel.StringType)),
	cel.Variable("visibilities", cel.ListType(cel.StringType)),
	cel.Variable("tag_search", cel.ListType(cel.StringType)),
	cel.Variable("order_by_pinned", cel.BoolType),
	cel.Variable("order_by_time_asc", cel.BoolType),
	cel.Variable("display_time_before", cel.IntType),
	cel.Variable("display_time_after", cel.IntType),
	cel.Variable("creator", cel.StringType),
	cel.Variable("uid", cel.StringType),
	cel.Variable("row_status", cel.StringType),
	cel.Variable("random", cel.BoolType),
	cel.Variable("limit", cel.IntType),
	cel.Variable("include_comments", cel.BoolType),
	cel.Variable("has_link", cel.BoolType),
	cel.Variable("has_task_list", cel.BoolType),
	cel.Variable("has_code", cel.BoolType),
	cel.Variable("has_incomplete_tasks", cel.BoolType),
}

type MemoFilter struct {
	ContentSearch      []string
	Visibilities       []model.Visibility
	TagSearch          []string
	OrderByPinned      bool
	OrderByTimeAsc     bool
	DisplayTimeBefore  int64
	DisplayTimeAfter   int64
	Creator            string
	RowStatus          model.RowStatus
	Random             bool
	Limit              int
	IncludeComments    bool
	HasLink            bool
	HasTaskList        bool
	HasCode            bool
	HasIncompleteTasks bool
}

func parseMemoFilter(expression string) (*MemoFilter, error) {
	e, err := cel.NewEnv(MemoFilterCELAttributes...)
	if err != nil {
		return nil, err
	}
	ast, issues := e.Compile(expression)
	if issues != nil {
		return nil, errors.Errorf("found issue %v", issues)
	}
	filter := &MemoFilter{}
	expr, err := cel.AstToParsedExpr(ast)
	if err != nil {
		return nil, err
	}
	callExpr := expr.GetExpr().GetCallExpr()
	findMemoField(callExpr, filter)
	return filter, nil
}

func findMemoField(callExpr *expr.Expr_Call, filter *MemoFilter) {
	if len(callExpr.Args) == 2 {
		idExpr := callExpr.Args[0].GetIdentExpr()
		if idExpr != nil {
			if idExpr.Name == "content_search" {
				contentSearch := []string{}
				for _, expr := range callExpr.Args[1].GetListExpr().GetElements() {
					value := expr.GetConstExpr().GetStringValue()
					contentSearch = append(contentSearch, value)
				}
				filter.ContentSearch = contentSearch
			} else if idExpr.Name == "visibilities" {
				visibilities := []model.Visibility{}
				for _, expr := range callExpr.Args[1].GetListExpr().GetElements() {
					value := expr.GetConstExpr().GetStringValue()
					visibilities = append(visibilities, model.Visibility(value))
				}
				filter.Visibilities = visibilities
			} else if idExpr.Name == "tag_search" {
				tagSearch := []string{}
				for _, expr := range callExpr.Args[1].GetListExpr().GetElements() {
					value := expr.GetConstExpr().GetStringValue()
					tagSearch = append(tagSearch, value)
				}
				filter.TagSearch = tagSearch
			} else if idExpr.Name == "order_by_pinned" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.OrderByPinned = value
			} else if idExpr.Name == "order_by_time_asc" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.OrderByTimeAsc = value
			} else if idExpr.Name == "display_time_before" {
				displayTimeBefore := callExpr.Args[1].GetConstExpr().GetInt64Value()
				filter.DisplayTimeBefore = displayTimeBefore
			} else if idExpr.Name == "display_time_after" {
				displayTimeAfter := callExpr.Args[1].GetConstExpr().GetInt64Value()
				filter.DisplayTimeAfter = displayTimeAfter
			} else if idExpr.Name == "creator" {
				creator := callExpr.Args[1].GetConstExpr().GetStringValue()
				filter.Creator = creator
			} else if idExpr.Name == "row_status" {
				rowStatus := model.RowStatus(callExpr.Args[1].GetConstExpr().GetStringValue())
				filter.RowStatus = rowStatus
			} else if idExpr.Name == "random" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.Random = value
			} else if idExpr.Name == "limit" {
				limit := int(callExpr.Args[1].GetConstExpr().GetInt64Value())
				filter.Limit = limit
			} else if idExpr.Name == "include_comments" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.IncludeComments = value
			} else if idExpr.Name == "has_link" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.HasLink = value
			} else if idExpr.Name == "has_task_list" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.HasTaskList = value
			} else if idExpr.Name == "has_code" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.HasCode = value
			} else if idExpr.Name == "has_incomplete_tasks" {
				value := callExpr.Args[1].GetConstExpr().GetBoolValue()
				filter.HasIncompleteTasks = value
			}
			return
		}
	}
	for _, arg := range callExpr.Args {
		callExpr := arg.GetCallExpr()
		if callExpr != nil {
			findMemoField(callExpr, filter)
		}
	}
}
