package test

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"testing"
)

func InitEtcd() {
	config.InitConfig()
	kku_stage.InitSlog(config.Config.DebugMode, nil, nil)
	kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password)
	_, list := service.KVList("")
	for _, pbKV := range list.ListKV {
		service.KVPut(kk_etcd_const.Config+pbKV.Key, pbKV.Value)
	}
}

func TestName111(t *testing.T) {
	InitEtcd()
}
