package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"log/slog"
	"sync"
	"testing"
	"time"
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

		userName = "admin"
		password = "admin"
	)

	InitEtcd(endpoints, userName, password)
	GetConfig(configKey, &GlobalConfig)

}

func TestRegisterService1(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.11:8484", "ssss")
	RegisterGrpcService(ctx, "128.2.2.11:84844", "ssss")
	w.Wait()
	cancelFunc()
}
func TestRegisterService2(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.3:8484", "ssss")
	w.Wait()
	cancelFunc()
}
func TestRegisterService3(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.3:8484", "haha")
	w.Wait()
	cancelFunc()
}

func TestGetServiceList(t *testing.T) {

	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 5)
		list, _ := ServerList(key_prefix.ServiceHttp)
		slog.Info("list", "list", list)
	}
}
func TestGetServiceList2(t *testing.T) {

	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 5)
		list, _ := ServerList(key_prefix.ServiceGrpc)
		slog.Info("list", "list", list)
	}
}
