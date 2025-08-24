package api_impl

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_def"
)

func RegisterFileDesc() {
	kk_grpc.GFileDescHub.RegisterFileDesc(api_def.File_internal_service_hub_service_api_def_service_proto)
}
