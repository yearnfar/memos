package db

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yearnfar/memos/internal/config"
)

var dbConn *gorm.DB

func Init() {
	cfg := config.GetApp().Database
	var err error
	switch cfg.Type {
	case "sqlite":
		dsn := config.GetPath(cfg.DSN)
		dbConn, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	return dbConn.WithContext(ctx)
}
