package kk_etcd_test

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"log"
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
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()

	mgrKV := kk_etcd.NewMgrKV()
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
			if err := mgrKV.PutYaml(tt.args.configKey, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("MgrKV.PutYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetYaml(t *testing.T) {
	var GlobalConfig configTest
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrKV := kk_etcd.NewMgrKV()
	err := mgrKV.GetYaml("my_config", &GlobalConfig)
	if err != nil {
		t.Errorf("MgrKV.GetYaml failed: %v", err)
	}
	t.Logf("GlobalConfig: %+v", GlobalConfig)
}

func TestPutJson(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()

	mgrKV := kk_etcd.NewMgrKV()
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
			if err := mgrKV.PutJson(tt.args.configKey, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("MgrKV.PutJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetJson(t *testing.T) {
	var GlobalConfig configTest
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrKV := kk_etcd.NewMgrKV()
	err := mgrKV.GetJson("my_config_json", &GlobalConfig)
	if err != nil {
		t.Errorf("MgrKV.GetJson failed: %v", err)
	}
	t.Logf("GlobalConfig: %+v", GlobalConfig)
}
