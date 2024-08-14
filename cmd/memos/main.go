package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/yearnfar/memos/internal/api"
	"github.com/yearnfar/memos/internal/config"
	"github.com/yearnfar/memos/server"
)

// 版本信息，在编译时自动生成
var (
	Version   = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func main() {
	app := &cli.App{
		Name:    "memos app",
		Usage:   "memos -c ",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "d",
				Aliases: []string{"dir"},
				Value:   ".",
				Usage:   "程序运行目录",
			},
			&cli.StringFlag{
				Name:    "c",
				Aliases: []string{"config"},
				Usage:   "配置文件地址",
				Value:   "./configs/config.toml",
				EnvVars: []string{"MEMOS_CONFIG_FILE"},
			},
		},
		Before: func(c *cli.Context) error {
			config.Version = Version
			config.BuildTime = BuildTime
			config.GitCommit = GitCommit

			runPath := c.String("d")
			configFile := c.String("c")
			config.Init(runPath, configFile)

			api.Init()
			return nil
		},
		Action: func(c *cli.Context) error {
			config.UpTime = time.Now()

			ctx, cancel := context.WithCancel(context.Background())
			srv := server.NewService(ctx)
			if err := srv.Start(ctx); err != nil {
				cancel()
				return err
			}

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				<-quit
				srv.Shutdown(ctx)
				cancel()
			}()

			<-ctx.Done()
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
