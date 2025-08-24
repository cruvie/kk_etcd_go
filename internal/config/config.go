package config

import (
	"log/slog"
	"os"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	Port         int    `yaml:"Port"`
	DebugMode    bool   `yaml:"DebugMode"`
	RootPassword string `yaml:"RootPassword"`
	Etcd         struct {
		Endpoints []string `yaml:"Endpoints"`
	} `yaml:"Etcd"`
	JWT struct {
		Key        string `yaml:"Key"`
		ExpireTime int    `yaml:"ExpireTime"`
	} `yaml:"JWT"`
}

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// slog.Info("workDir", "dir", workDir)

	// docker
	// data, err := os.ReadFile(workDir + "/kk_etcd_go/internal/config/config.yml")
	data, err := os.ReadFile(workDir + "/internal/config/config.yml")
	if err != nil {
		slog.Error("unable to read config.yaml", "err", err)
		panic(err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		slog.Error("unable to unmarshal config.yaml", "err", err)
		panic(err)
	}
}

var LogCfg kk_stage.ConfigLog

func InitLog(stage *kk_stage.Stage) {
	LogCfg = kk_stage.ConfigLog{
		DebugMode:  Config.DebugMode,
		Lumberjack: kk_stage.DefaultLogConfig(time.Now(), consts.ServiceName),
		StartTime:  stage.StartTime,
	}
	LogCfg.Init()
}
