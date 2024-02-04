package kk_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
)

func InitEtcd(endpoints []string, userName string, password string) error {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	return service.InitEtcd(stage, endpoints, userName, password)
}
