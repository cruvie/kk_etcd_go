package deregisterServer

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/server/api_def"
)

func (x *api) Handler() (*api_def.DeregisterServer_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	server := &server_hub.SerServer{}
	err := server.DeregisterServer(x.stage,
		x.In.GetServer().GetServerRegistration())
	if err != nil {
		return nil, err
	}
	return &api_def.DeregisterServer_Output{}, nil
}
