package internal_handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/internal_handler/internal_service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HServer struct{}

var serInternalServer internal_service.SerServer

// ServerList
// serverName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func (HServer) ServerList(stage *kk_stage.Stage, param *kk_etcd_models.ServerListParam) (error, *kk_etcd_models.ServerListResponse) {
	span := stage.StartTrace("ServerList")
	defer span.End()
	serverList, err := serInternalServer.ServerList(global_model.GetClient(stage), kk_etcd_models.ServerType(param.GetServerType()))
	return err, &kk_etcd_models.ServerListResponse{
		ServerList: serverList,
	}
}

func (HServer) DeregisterServer(stage *kk_stage.Stage, param *kk_etcd_models.DeregisterServerParam) (error, *kk_etcd_models.DeregisterServerResponse) {
	span := stage.StartTrace("DeregisterServer")
	defer span.End()
	err := serInternalServer.DeregisterServer(stage, param)
	if err != nil {
		return err, nil
	}
	return nil, &kk_etcd_models.DeregisterServerResponse{}
}
