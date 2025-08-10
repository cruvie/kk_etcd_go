package deregisterService

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
)

type Api struct {
	*kk_grpc.DefaultApi[api_def.DeregisterService_Input]
}

func NewApi() *Api {
	return &Api{
		DefaultApi: kk_grpc.NewDefaultApi[api_def.DeregisterService_Input](),
	}
}
