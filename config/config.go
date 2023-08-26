package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var GlobalConfig Config

type Config struct {
	ServerAddr string `yaml:"ServerAddr"`
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
	//log.Println("workDir:", workDir)

	//docker
	data, err := os.ReadFile(workDir + "/kk_etcd_go/config/config.yml")
	//data, err := os.ReadFile(workDir + "/config/config.yml")
	if err != nil {
		log.Println("unable to read config.yaml: ", err)
		return
	}

	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		log.Println("unable to unmarshal config.yaml: ", err)
		return
	}
}
