package deregisterServer

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) handler() (error, *DeregisterServer_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := server_hub.SerServer{}.DeregisterServer(x.stage,
		kk_etcd_models.ServerType(x.in.GetServer().GetServerType()),
		x.in.GetServer().GetEndpointKey())
	if err != nil {
		return err, nil
	}
	return nil, &DeregisterServer_Output{}
}
