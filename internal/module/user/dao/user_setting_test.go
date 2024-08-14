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

func TestDao_UpsertUserSetting(t *testing.T) {
	config.Init("../../../../")
	db.Init()

	setting := &model.UserSetting{
		UserId: 2,
		Key:    model.UserSettingKeyAccessToken,
		Value:  fmt.Sprintf(`{"time": "%s"}`, datetime.Strftime(time.Now(), "%Y-%m-%d %H:%M:%S")),
	}
	err := New().UpsertUserSetting(context.Background(), setting)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("done")
}
