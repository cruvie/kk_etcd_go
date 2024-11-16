package kk_etcd_models

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestBindCheckGen(t *testing.T) {
	kk_api_gen.BindCheckGen(
		//RoleRevokePermissionParam{},
		ServerListParam{},
		//DeregisterServerParam{},
	)
}
