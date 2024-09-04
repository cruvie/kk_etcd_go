package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_config_interface"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
)

func initTestEnv() CloseFunc {
	var (
		stage = kk_stage.NewStage(context.Background(), "Test", true)
	)
	configLog := kk_log.ConfigLog{
		Lumberjack: kk_log.DefaultLogConfig(consts.ServerName),
	}
	configLog.Init(&kk_config_interface.InitArgs{
		DebugMode:   stage.DebugMode,
		ServiceName: stage.ServiceName,
	})
	defer configLog.Close()
	closeFunc, err := InitClient(&InitClientConfig{
		Endpoints: []string{"http://127.0.0.1:2379"},
		UserName:  "root",
		Password:  "root",
		DebugMode: stage.DebugMode})
	if err != nil {
		panic(err)
	}
	return closeFunc
}
