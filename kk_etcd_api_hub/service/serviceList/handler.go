package serviceList

import (
	"github.com/cruvie/kk_etcd_go/internal/service_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
)

func (x *api) Handler() (*api_def.ServiceList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	service := &service_hub.SerService{}
	serviceList, err := service.ServiceList(x.In.GetServiceType(), x.In.GetServiceName())
	return &api_def.ServiceList_Output{
		ServiceList: serviceList,
	}, err
}
