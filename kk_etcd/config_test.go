package kk_etcd

import (
	"log/slog"
	"testing"
)

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

var GlobalConfig config

func TestGetConfig(t *testing.T) {

	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		configKey = "my_config"

		userName = "kk_etcd"
		password = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}
	err = GetConfig(configKey, &GlobalConfig)
	if err != nil {
		slog.Error("GetConfig failed", "err", err)
	}

}
func TestSetConfig(t *testing.T) {
	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}
	type args struct {
		configKey string
		config    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"my_config", "safdsfgdsg"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetConfig(tt.args.configKey, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("SetConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
