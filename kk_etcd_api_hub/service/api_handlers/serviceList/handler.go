package serviceList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.ServiceList_Output, error) {

	service := &service_hub.SerService{}
	serviceList, err := service.ServiceList(x.In.GetServiceType(), x.In.GetServiceName())
	return &api_def.ServiceList_Output{
		ServiceList: serviceList,
	}, err
}
