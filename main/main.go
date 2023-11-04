package main

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/api_etcd"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	_ "github.com/cruvie/kk_etcd_go/main/docs"
	"log/slog"
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
// @BasePath	/User
func main() {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("main panic", "error", r)
		}
	}()
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	defer mainClose(stage)

	config.InitConfig()

	kku_stage.InitSlog(config.Config.DebugMode, nil, nil)
	kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password)
	api_etcd.ApiEtcd(stage)
}
func mainClose(stage *kku_stage.Stage) {
}
