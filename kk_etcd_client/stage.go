package kk_etcd_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
)

// GlobalStage for kk_etcd client sdk
var GlobalStage *kk_stage.Stage

func initGlobalStage(debugMode bool) {
	GlobalStage = kk_stage.NewStage(context.Background(), kk_etcd_const.ServerName, debugMode)
}
