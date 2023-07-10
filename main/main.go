package main

import (
	"gitee.com/cruvie/kk_go_kit/kk_config/kkcfg_log"
	"kk_etcd_go/api_etcd"
	"kk_etcd_go/config"
	"kk_etcd_go/kk_etcd"
	_ "kk_etcd_go/main/docs"
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
	defer mainClose()
	kkcfg_log.InitLog()
	config.InitConfig()

	kk_etcd.InitEtcd([]string{config.GlobalConfig.Etcd.Endpoint})
	api_etcd.ApiEtcd()
}
func mainClose() {
}
