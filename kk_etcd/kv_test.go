package kk_etcd

import (
	"log/slog"
	"testing"
)

type configTest struct {
	ServerAddr string
	Postgres   struct {
		Dsn  string
		Port int
	}
	Redis struct {
		Addr     string
		Password string
	}
}

func TestPutYaml(t *testing.T) {
	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password, true)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}

	type args struct {
		configKey string
		config    configTest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{configKey: "my_config", config: configTest{
			ServerAddr: "http://example.com",
			Postgres: struct {
				Dsn  string
				Port int
			}{
				Dsn:  "postgres://username:password@localhost/dbname",
				Port: 5432,
			},
			Redis: struct {
				Addr     string
				Password string
			}{
				Addr:     "localhost:6379",
				Password: "secret",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PutYaml(tt.args.configKey, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("PutYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetYaml(t *testing.T) {
	var GlobalConfig configTest
	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password, true)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}
	err = GetYaml("my_config", &GlobalConfig)
	if err != nil {
		slog.Error("GetYaml failed", "err", err)
	}
	t.Logf("GlobalConfig: %+v", GlobalConfig)

}

func TestPutJson(t *testing.T) {
	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password, true)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}

	type args struct {
		configKey string
		config    configTest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{configKey: "my_config_json", config: configTest{
			ServerAddr: "http://example.com",
			Postgres: struct {
				Dsn  string
				Port int
			}{
				Dsn:  "postgres://username:password@localhost/dbname",
				Port: 5432,
			},
			Redis: struct {
				Addr     string
				Password string
			}{
				Addr:     "localhost:6379",
				Password: "secret",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PutJson(tt.args.configKey, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("PutYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetJson(t *testing.T) {
	var GlobalConfig configTest
	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)

	err := InitEtcd(endpoints, userName, password, true)
	if err != nil {
		slog.Error("InitEtcd failed", "err", err)
	}
	err = GetJson("my_config_json", &GlobalConfig)
	if err != nil {
		slog.Error("GetYaml failed", "err", err)
	}
	t.Logf("GlobalConfig: %+v", GlobalConfig)
}
