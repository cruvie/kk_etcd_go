package config

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

var Config config

type config struct {
	ServerAddr   string `yaml:"ServerAddr"`
	DebugMode    bool   `yaml:"DebugMode"`
	RootPassword string `yaml:"RootPassword"`
	Admin        struct {
		UserName string `yaml:"UserName"`
		Password string `yaml:"Password"`
	} `yaml:"Admin"`
	Etcd struct {
		Endpoint string `yaml:"Endpoint"`
	} `yaml:"Etcd"`
	JWT struct {
		ExpireTime int `yaml:"ExpireTime"`
	} `yaml:"JWT"`
}

func InitConfig() {

	workDir, _ := os.Getwd()
	//slog.Info("workDir", "dir", workDir)

	//docker
	data, err := os.ReadFile(workDir + "/kk_etcd_go/internal/config/config.yml")
	//data, err := os.ReadFile(workDir + "/internal/config/config.yml")
	if err != nil {
		slog.Info("unable to read config.yaml", "err", err)
		return
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		slog.Info("unable to unmarshal config.yaml", "err", err)
		return
	}
}
