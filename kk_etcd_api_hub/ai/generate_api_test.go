package ai

import (
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_grpc/grpc_api_gen"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/ai/api_def"
)

func TestGenImpl(t *testing.T) {
	grpc_api_gen.GenerateImpl(grpc_api_gen.GenerateImplInput{
		ServerName:      "AI",
		Methods:         api_def.AI_ServiceDesc.Methods,
		ApiDefPkgPath:   "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/ai/api_def",
		HandlersPkgPath: "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/ai/api_handlers",
	})
}
