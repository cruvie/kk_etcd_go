package serviceList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_housekeeper"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.ServiceList_Output, error) {

	service := &service_housekeeper.SerService{}
	serviceList, err := service.ServiceList(x.In.GetServiceType(), x.In.GetServiceName())
	return &api_def.ServiceList_Output{
		ServiceList: serviceList,
	}, err
}
