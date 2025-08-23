package backup

import (
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_grpc/grpc_api_gen"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def"
)

func TestGenImpl(t *testing.T) {
	grpc_api_gen.GenerateImpl(grpc_api_gen.GenerateImplInput{
		ServerName:      "Backup",
		Methods:         api_def.Backup_ServiceDesc.Methods,
		ApiDefPkgPath:   "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def",
		HandlersPkgPath: "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_handlers",
	})
}
