package test

import (
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"testing"
)

func InitEtcd() {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName())
	config.InitConfig()
	kk_stage.InitSlog(config.Config.DebugMode, nil, nil)
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

func TestName(t *testing.T) {
	a := make([]*string, 0)
	aa := make([]*string, 0)
	b := "11"
	a = append(a, &b)
	c := "222"
	a = append(a, &c)
	fmt.Println(a)
	fmt.Println(len(a))
	clear(a)
	fmt.Println(a)
	fmt.Println(len(a))
	a = make([]*string, 0)
	fmt.Println(aa)
	fmt.Println(len(aa))
}
