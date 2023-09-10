package test

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"testing"
)

func InitEtcd() {
	config.InitConfig()
	kku_log.InitSlog(config.Config.DebugMode, nil, nil)
	kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password)
	_, list := service.KVList("")
	for _, pbKV := range list.ListKV {
		service.KVPut(key_prefix.Config+pbKV.Key, pbKV.Value)
	}
}

func TestName111(t *testing.T) {
	InitEtcd()
}
