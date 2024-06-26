package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HServer struct{}

var serServerV service.SerServer

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func (HServer) ServerList(stage *kk_stage.Stage, param *kk_etcd_models.ServerListParam) (error, *kk_etcd_models.ServerListResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	serverList, err := serServerV.ServerList(param.GetPrefix())
	return err, &kk_etcd_models.ServerListResponse{
		ServerList: serverList,
	}
}
