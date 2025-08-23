package roleList

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def"
)

type Api struct {
	*kk_grpc.DefaultApi[api_def.RoleList_Input]
}

func NewApi() *Api {
	return &Api{
		DefaultApi: kk_grpc.NewDefaultApi[api_def.RoleList_Input](),
	}
}
