package config

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
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
		Enable       bool   `yaml:"Enable"`
		LLMModel     string `yaml:"LLMModel"`
		OllamaServer string `yaml:"OllamaServer"`
		//WeaviateIpPort string `yaml:"WeaviateIpPort"`
		//EtcdDocPath    string `yaml:"EtcdDocPath"`
		//ReInitWeaviate bool   `yaml:"ReInitWeaviate"`
	} `yaml:"AI"`
	/*

	   AI:
	     Enable: true
	     LLMModel: qwen2.5:0.5b
	     # default: http://127.0.0.1:11434
	     OllamaServer: http://127.0.0.1:11434
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
	//slog.Info("workDir", "dir", workDir)

	//docker
	//data, err := os.ReadFile(workDir + "/kk_etcd_go/internal/config/config.yml")
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
