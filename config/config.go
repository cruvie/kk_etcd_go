package config

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

var Config config

type config struct {
	ServerAddr string `yaml:"ServerAddr"`
	DebugMode  bool   `yaml:"DebugMode"`
	Admin      struct {
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
	//slog.Info("workDir:", workDir)

	//docker
	//data, err := os.ReadFile(workDir + "/kk_etcd_go/config/config.yml")
	data, err := os.ReadFile(workDir + "/config/config.yml")
	if err != nil {
		slog.Info("unable to read config.yaml: ", err)
		return
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		slog.Info("unable to unmarshal config.yaml: ", err)
		return
	}
}
