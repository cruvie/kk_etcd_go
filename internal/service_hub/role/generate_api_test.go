package api_def

import (
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_grpc/grpc_api_gen"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def"
	"github.com/cruvie/kk_etcd_go/internal/utils"
)

func TestGenImpl(t *testing.T) {
	grpc_api_gen.GenerateImpl(grpc_api_gen.GenerateImplInput{
		ServerName:      "Role",
		Methods:         api_def.Role_ServiceDesc.Methods,
		ApiDefPkgPath:   "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def",
		HandlersPkgPath: "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_handlers",
	})
	utils.GenerateImpl(utils.GenerateImplInput{
		ServerName:      "Role",
		Methods:         api_def.Role_ServiceDesc.Methods,
		ApiDefPkgPath:   "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def",
		HandlersPkgPath: "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_handlers",
	})
}
