package config

import (
	syslog "log"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v8"
	"github.com/pkg/errors"
	"github.com/subosito/gotenv"
	"github.com/yearnfar/gokit/fsutil"
)

const (
	DevMode  = "dev"  // 开发环境
	TestMode = "test" // 测试环境
	ProdMode = "prod" // 生产环境
)

// 版本信息，在编译时自动生成
var (
	Version   = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

var UpTime time.Time // 系统启动时间

// App 配置结构体
type App struct {
	Logger   Logger   `toml:"logger" envPrefix:"LOGGER_"`     // 日志配置
	Server   Server   `toml:"server" envPrefix:"SERVER_"`     // http server配置
	Database Database `toml:"database" envPrefix:"DATABASE_"` // database configuration
	JWT      JWT      `toml:"jwt"  envPrefix:"JWT_"`          // jwt配置`
}

type Logger struct {
	Level string `toml:"level" evn:"LEVEL"`
}

type Server struct {
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	Secret string `toml:"secret"`
}

type JWT struct {
	Key string `toml:"key" env:"KEY"`
}

type Database struct {
	Type string `toml:"type" env:"TYPE"`
	DSN  string `toml:"dsn" env:"DSN"`
}

var app *App       // 项目配置
var appPath string // 项目运行目录

// Init 初始化配置
func Init(runDir string, cfgFiles ...string) {
	var err error
	if runDir == "" {
		runDir, err = os.Getwd()
		if err != nil {
			syslog.Fatal("获取运行地址失败", err)
		}
	} else if runDir != "" && !filepath.IsAbs(runDir) {
		runDir, err = filepath.Abs(runDir)
		if err != nil {
			syslog.Fatal("获取绝对地址失败", err)
		}
	}
	appPath = runDir
	app = &App{}
	var cfgFile string
	if len(cfgFiles) > 0 {
		cfgFile = cfgFiles[0]
	}
	// syslog.Printf("app_path: %s\n", appPath)
	err = initConfig(app, cfgFile)
	if err != nil {
		syslog.Fatal("初始化配置失败", err)
		return
	}
}

// GetApp 获取配置实例
func GetApp() *App {
	if app == nil {
		syslog.Fatal("应用未实例化")
	}
	return app
}

func initConfig(app *App, cfgFile string) (err error) {
	if cfgFile != "" {
		cfgFile, err = filepath.Abs(cfgFile)
		if err != nil {
			return
		} else if !fsutil.IsFile(cfgFile) {
			err = errors.Errorf("不存在配置文件%s", cfgFile)
			return
		}
	} else {
		if defCfgFile := filepath.Join(appPath, "configs", "config.toml"); fsutil.IsFile(defCfgFile) {
			cfgFile = defCfgFile
		}
	}
	// syslog.Printf("config_path: %s\n", cfgFile)
	if cfgFile != "" {
		if _, err := toml.DecodeFile(cfgFile, app); err != nil {
			return errors.Wrap(err, "解析配置文件失败")
		}
	}
	if envFile := filepath.Join(appPath, ".env"); fsutil.IsFile(envFile) {
		if err := gotenv.Load(envFile); err != nil {
			return err
		}
	}
	if err := env.ParseWithOptions(app, env.Options{Prefix: "MOSS_"}); err != nil {
		return err
	}
	return nil
}

// GetPath 输入项目中的相对地址，获得绝对地址
func GetPath(elem ...string) string {
	rp := appPath
	if len(elem) == 0 {
		return rp
	}
	return filepath.Join(rp, filepath.Join(elem...))
}
