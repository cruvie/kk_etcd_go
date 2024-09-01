package internal_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
)

// GlobalStage for kk_etcd client sdk and main server
var GlobalStage *kk_stage.Stage

var cancelFunc context.CancelFunc

func InitGlobalStage(debugMode bool) {
	var ctx context.Context
	ctx, cancelFunc = context.WithCancel(context.Background())
	GlobalStage = kk_stage.NewStage(ctx, kk_etcd_const.ServerName, debugMode)
}

func CloseGlobalStage() {
	cancelFunc()
}
