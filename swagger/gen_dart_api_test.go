package swagger

import (
	_ "embed"
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"testing"
)

//go:embed swagger.json
var jsonSwagger string

func TestGenDartApi(t *testing.T) {
	tagPaths := kk_swagger.ClassifyByTag(jsonSwagger)
	kk_swagger.GenerateDartApi("../lib/kk_etcd_api_hub", tagPaths)
}
