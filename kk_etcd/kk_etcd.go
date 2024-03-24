package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
)

func InitEtcd(endpoints []string, userName string, password string, debugMode bool) error {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), debugMode)
	return service.InitEtcd(stage, endpoints, userName, password)
}
