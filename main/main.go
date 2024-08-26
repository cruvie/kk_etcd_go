package main

import (
	"gitee.com/cruvie/kk_go_kit/kk_config_interface"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/api_etcd"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	_ "github.com/cruvie/kk_etcd_go/swagger"
)

//	@title			kk_etcd_go API
//	@version		1.0
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
	config.InitConfig()
	global_stage.InitGlobalStage(config.Config.DebugMode)
	configLog := kk_log.ConfigLog{
		Lumberjack: kk_log.DefaultLogConfig(kk_etcd_const.ServerName),
	}
	init := &kk_config_interface.InitArgs{
		DebugMode:   global_stage.GlobalStage.DebugMode,
		ServiceName: global_stage.GlobalStage.ServiceName,
	}
	configLog.Init(init)
	defer configLog.Close()

	internal_client.InitEtcd()
	internal_client.InitInternalClient()

	api_etcd.ApiEtcd(global_stage.GlobalStage)
}
