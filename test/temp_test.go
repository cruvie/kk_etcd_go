package test

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"testing"
)

func InitEtcd() {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	config.InitConfig()
	kku_stage.InitSlog(config.Config.DebugMode, nil, nil)
	err := kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password)
	if err != nil {
		return
	}
	_, list := service.KVList(stage, "")
	for _, pbKV := range list.ListKV {
		service.KVPut(stage, kk_etcd_const.Config+pbKV.Key, pbKV.Value)
	}
}

func TestName111(t *testing.T) {
	InitEtcd()
}
