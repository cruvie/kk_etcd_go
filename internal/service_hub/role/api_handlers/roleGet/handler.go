package roleGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/util_role"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.RoleGet_Output, error) {
	role, err := util_role.GetRole(stage, x.In.GetName())
	return &api_def.RoleGet_Output{
		Role: role,
	}, err
}
