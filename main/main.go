package main

import (
	"context"

	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_pprof"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/api_etcd"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/etcd_ai"
	"github.com/cruvie/kk_etcd_go/internal/service_housekeeper"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	startTime := time.Now()

	config.InitConfig()

	stage := kk_stage.NewStage(context.Background(), consts.ServiceName).SetStartTime(startTime)
	{
		config.InitLog(stage)
		defer func() {
			config.LogCfg.Close()
		}()
	}
	{
		//init pprof
		pprof := kk_pprof.ConfigPprofWeb{
			Port:        2444,
			EnableBlock: true,
			EnableMutex: true,
		}
		pprof.Init(stage)
		pprof.Print()
		defer pprof.Close()
	}

	internal_client.InitEtcd(stage)

	etcdCfg := clientv3.Config{
		Endpoints:   config.Config.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    consts.UserRoot,
		Password:    config.Config.RootPassword,
	}
	closeFunc, err := kk_etcd.InitClient(&kk_etcd.InitClientConfig{
		Config:    etcdCfg,
		DebugMode: config.Config.DebugMode,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		err := closeFunc()
		if err != nil {
			slog.Error("close etcd client failed", "err", err)
		}
	}()

	service_housekeeper.InitKubernetesClient(etcdCfg)
	defer func() {
		service_housekeeper.CloseKubernetesClient()
	}()
	service_housekeeper.InitServiceHub()

	kkServer := kk_server.NewKKServer(5*time.Second, stage)
	kkServer.Add("ApiGrpc", 0, api_etcd.NewGrpcServer(internal_client.GlobalStage))
	//kkServer.Add("ApiMCP", 0, api_etcd.ApiMCP())
	//kkServer.Add("etcd_maintain", 0, service_hub.NewEtcdMaintain())
	kkServer.Add("etcd_ai", 0, etcd_ai.EtcdAIService())
	kkServer.ServeAndWait()
}
