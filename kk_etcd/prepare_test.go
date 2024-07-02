package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
)

func initTestEnv() {
	var (
		stage     = kk_stage.NewStage(context.Background(), "Test", true)
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)
	configLog := kk_log.ConfigLog{
		Lumberjack: kk_log.DefaultLogConfig(kk_etcd_const.ServerName),
	}
	configLog.Init(stage.DebugMode)
	defer configLog.Close()
	err := InitEtcd(endpoints, userName, password, stage.DebugMode)
	if err != nil {
		panic(err)
	}
}
