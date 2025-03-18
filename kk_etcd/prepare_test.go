package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"time"
)

func initTestEnv() CloseFunc {
	var (
		stage = kk_stage.NewStage(context.Background(), "Test", true)
	)
	configLog := kk_log.ConfigLog{
		DebugMode:  stage.DebugMode,
		Lumberjack: kk_log.DefaultLogConfig(time.Now(), consts.ServerName),
	}
	configLog.Init()
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
