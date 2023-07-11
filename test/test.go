package test

import "github.com/cruvie/kk_etcd_go/kk_etcd"

var GlobalConfig Config

type Config struct {
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
)

func InitEtcd() {
	kk_etcd.InitEtcd(endpoints)
	kk_etcd.GetConfig(configKey, &GlobalConfig)
}
