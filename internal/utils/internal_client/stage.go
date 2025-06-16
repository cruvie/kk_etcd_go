package internal_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
)

// GlobalStage for kk_etcd client sdk and main service
var GlobalStage *kk_stage.Stage

var cancelFunc context.CancelFunc

func InitGlobalStage() {
	var ctx context.Context
	ctx, cancelFunc = context.WithCancel(context.Background())
	GlobalStage = kk_stage.NewStage(ctx, consts.ServiceName)
}

func CloseGlobalStage() {
	cancelFunc()
}
