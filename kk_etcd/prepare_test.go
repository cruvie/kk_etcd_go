package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func initTestEnv() {
	var (
		stage     = kk_stage.NewStage(context.Background(), "Test", true)
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		userName  = "kk_etcd"
		password  = "kk_etcd"
	)
	kk_stage.InitSlog(stage.DebugMode, nil, nil)
	err := InitEtcd(endpoints, userName, password, stage.DebugMode)
	if err != nil {
		panic(err)
	}
}
