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
	AI struct {
		Enable        bool   `yaml:"Enable"`
		LLMModel      string `yaml:"LLMModel"`
		OllamaService string `yaml:"OllamaService"`
		// WeaviateIpPort string `yaml:"WeaviateIpPort"`
		// EtcdDocPath    string `yaml:"EtcdDocPath"`
		// ReInitWeaviate bool   `yaml:"ReInitWeaviate"`
	} `yaml:"AI"`
	MCPService struct {
		Enable bool `yaml:"Enable"`
		Port   int  `yaml:"Port"`
	} `yaml:"MCPService"`
	/*

	   AI:
	     Enable: true
	     LLMModel: qwen2.5:0.5b
	     # default: http://127.0.0.1:11434
	     OllamaService: http://127.0.0.1:11434
	     # default 127.0.0.1:9035
	     WeaviateIpPort: 127.0.0.1:9035
	     EtcdDocPath: /Users/cruvie/cruvie-space/code-hub/temp/etcd-website/content/en/docs/v3.5
	     # if true, will re-init weaviate
	     ReInitWeaviate: false
	*/
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
