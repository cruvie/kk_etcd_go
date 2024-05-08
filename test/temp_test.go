package test

import (
	"context"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"testing"
)

func InitEtcd() {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	config.InitConfig()
	kk_stage.InitSlog(stage.DebugMode, nil, nil)
	err := kk_etcd.InitEtcd([]string{config.Config.Etcd.Endpoint}, config.Config.Admin.UserName, config.Config.Admin.Password, stage.DebugMode)
	if err != nil {
		return
	}
	_, list := service.KVList(stage, "")
	for _, pbKV := range list.ListKV {
		_ = service.KVPut(stage, pbKV.Key, pbKV.Value)
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
	fmt.Println(aa)
	fmt.Println(len(aa))
}
