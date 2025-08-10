package api_def

import (
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_grpc/grpc_api_gen"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"
)

func TestGenImpl(t *testing.T) {
	grpc_api_gen.GenerateImpl(grpc_api_gen.GenerateImplInput{
		ServerName:      "Role",
		Methods:         api_def.Role_ServiceDesc.Methods,
		ApiDefPkgPath:   "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def",
		HandlersPkgPath: "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_handlers",
	})
}
