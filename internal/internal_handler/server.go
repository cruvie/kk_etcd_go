package internal_handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/internal_handler/internal_service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HServer struct{}

var serInternalServer internal_service.SerServer

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func (HServer) ServerList(stage *kk_stage.Stage, param *kk_etcd_models.ServerListParam) (error, *kk_etcd_models.ServerListResponse) {
	span := stage.StartTrace("ServerList")
	defer span.End()
	serverList, err := serInternalServer.ServerList(param.GetPrefix())
	return err, &kk_etcd_models.ServerListResponse{
		ServerList: serverList,
	}
}
