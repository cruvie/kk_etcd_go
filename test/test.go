package test

import "github.com/cruvie/kk_etcd_go/kk_etcd_client"

var Config config

type config struct {
	ServerAddr string `yaml:"ServerAddr"`
	Postgres   struct {
		Dsn string `yaml:"Dsn"`
	} `yaml:"Postgres"`
	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
	} `yaml:"Redis"`
	MinIO struct {
		AccessEndpoint string `yaml:"AccessEndpoint"`
	} `yaml:"MinIO"`
}

var (
	endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
	configKey = "my_config"

	userName = "admin"
	password = "admin"
)

func InitEtcd() {
	kk_etcd_client.InitEtcd(endpoints, userName, password)
	kk_etcd_client.GetConfig(configKey, &Config)
}
