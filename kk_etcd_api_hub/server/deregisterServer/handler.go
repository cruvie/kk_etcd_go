package deregisterServer

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) Handler() (*DeregisterServer_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := server_hub.SerServer{}.DeregisterServer(x.stage,
		kk_etcd_models.ServerType(x.In.GetServer().GetServerType()),
		x.In.GetServer().GetEndpointKey())
	if err != nil {
		return nil, err
	}
	return &DeregisterServer_Output{}, nil
}
