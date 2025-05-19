package deregisterService

import (
	"github.com/cruvie/kk_etcd_go/internal/service_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
)

func (x *api) Handler() (*api_def.DeregisterService_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	server := &service_hub.SerService{}
	err := server.DeregisterService(x.stage,
		x.In.GetService().GetServiceRegistration())
	if err != nil {
		return nil, err
	}
	return &api_def.DeregisterService_Output{}, nil
}
