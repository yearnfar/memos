package dao

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/yearnfar/gokit/datetime"

	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/internal/module/user/model"
	"github.com/yearnfar/memos/internal/pkg/db"
)

func TestDao_CreateUser(t *testing.T) {
	config.Init("../../../../")
	db.Init()

	user := &model.User{
		RowStatus: model.Normal,
		Role:      model.RoleUser,
		Username:  fmt.Sprintf("test_%s", datetime.Strftime(time.Now(), "%-Y%m%d_%H%M%S")),
	}
	err := New().CreateUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("done")
}

func TestDao_GetUserById(t *testing.T) {
	config.Init("../../../../")
	db.Init()

	ctx := context.Background()
	user, err := New().FindUserById(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)
}

func TestDao_GetUserByUsername(t *testing.T) {
	config.Init("../../../../")
	db.Init()

	ctx := context.Background()
	user, err := New().FindUserByUsername(ctx, "yearnfar")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)
}

func TestDao_GetUser(t *testing.T) {
	config.Init("../../../../")
	db.Init()

	ctx := context.Background()
	user, err := New().FindUser(ctx, []string{"username=?"}, []any{"yearnfar"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)
}
