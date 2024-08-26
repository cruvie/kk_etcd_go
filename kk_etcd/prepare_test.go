package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_config_interface"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
)

func initTestEnv() {
	var (
		stage = kk_stage.NewStage(context.Background(), "Test", true)
	)
	configLog := kk_log.ConfigLog{
		Lumberjack: kk_log.DefaultLogConfig(kk_etcd_const.ServerName),
	}
	configLog.Init(&kk_config_interface.InitArgs{
		DebugMode:   stage.DebugMode,
		ServiceName: stage.ServiceName,
	})
	defer configLog.Close()
	err := InitClient(&InitClientConfig{
		Endpoints: []string{"http://127.0.0.1:2379"},
		UserName:  "kk_etcd",
		Password:  "kk_etcd",
		DebugMode: stage.DebugMode})
	if err != nil {
		panic(err)
	}
}
