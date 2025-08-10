package deregisterService

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.DeregisterService_Output, error) {
	server := &service_hub.SerService{}
	err := server.DeregisterService(stage,
		x.In.GetService().GetServiceRegistration())
	if err != nil {
		return nil, err
	}
	return &api_def.DeregisterService_Output{}, nil
}
