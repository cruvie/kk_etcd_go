package main

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_pprof"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/api_etcd"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/etcd_ai"
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	_ "github.com/cruvie/kk_etcd_go/swagger"
	"log/slog"
	"time"
)

//	@title			kk_etcd_go API
//	@version		1.4.1
//	@description	kk_etcd_go terms.
//	@termsOfService	http://kk_etcd_go/terms/

//	@contact.name	kk_etcd_go
//	@contact.url	http://kk_etcd_go/support
//	@contact.email	kk_etcd_go@qq.com

//	@license.name	Secret
//	@license.url	http://kk_etcd_go/licenses/LICENSE-2.0.html

// @host		localhost:2333
// @BasePath	/
func main() {
	startTime := time.Now()

	config.InitConfig()

	stage := kk_stage.NewStage(context.Background(), consts.ServerName, config.Config.DebugMode).SetStartTime(startTime)

	{
		//init log
		configLog := kk_log.ConfigLog{
			DebugMode:  config.Config.DebugMode,
			Lumberjack: kk_log.DefaultLogConfig(time.Now(), consts.ServerName),
			StartTime:  stage.StartTime,
		}
		configLog.Init()
		defer configLog.Close()
	}
	{
		//init pprof
		pprof := kk_pprof.ConfigPprofWeb{
			Port:        2444,
			EnableBlock: true,
			EnableMutex: true,
		}
		pprof.Init(stage)
	}

	internal_client.InitEtcd()

	closeFunc, err := kk_etcd.InitClient(&kk_etcd.InitClientConfig{
		Endpoints: config.Config.Etcd.Endpoints,
		UserName:  consts.UserRoot,
		Password:  config.Config.RootPassword,
		DebugMode: internal_client.GlobalStage.DebugMode})
	if err != nil {
		panic(err)
	}
	defer func() {
		err := closeFunc()
		if err != nil {
			slog.Error("close etcd client failed", "err", err)
		}
	}()

	server_hub.InitServiceHub()

	kkServer := kk_server.NewKKServer(5*time.Second, stage)
	kkServer.Add("ApiHttp", 0, api_etcd.ApiHttp(internal_client.GlobalStage))
	kkServer.Add("ApiMCP", 0, api_etcd.ApiMCP())
	kkServer.Add("etcd_maintain", 0, server_hub.NewEtcdMaintain())
	kkServer.Add("etcd_ai", 0, etcd_ai.EtcdAIServer())
	kkServer.ServeAndWait()
}
