package kk_etcd_test

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func initTestEnv() kk_etcd.CloseFunc {
	var (
		stage = kk_stage.NewStage(context.Background(), "Test", true)
	)
	configLog := kk_log.ConfigLog{
		DebugMode:  stage.DebugMode,
		Lumberjack: kk_log.DefaultLogConfig(time.Now(), consts.ServiceName),
		StartTime:  time.Now(),
	}
	configLog.Init()
	defer configLog.Close()
	closeFunc, err := kk_etcd.InitClient(&kk_etcd.InitClientConfig{
		Config: clientv3.Config{
			Endpoints: []string{"http://127.0.0.1:2379"},
			Username:  "root",
			Password:  "root",
		},
		DebugMode: stage.DebugMode})
	if err != nil {
		panic(err)
	}
	return closeFunc
}
