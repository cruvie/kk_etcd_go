package main

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/api_etcd"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	_ "github.com/cruvie/kk_etcd_go/main/docs"
	"time"
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
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName())

	config.InitConfig()

	kk_stage.InitSlog(config.Config.DebugMode, nil, nil)

	kk_jwt.InitJwt(config.Config.JWT.Key, time.Duration(config.Config.JWT.ExpireTime)*time.Hour)

	err := kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password)
	if err != nil {
		panic(err)
	}

	api_etcd.ApiEtcd(stage)
}
