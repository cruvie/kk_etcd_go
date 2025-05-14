package roleGet

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/util_role"
)

func (x *api) Handler() (*api_def.RoleGet_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	role, err := util_role.GetRole(x.stage, x.In.GetName())
	return &api_def.RoleGet_Output{
		Role: role,
	}, err
}
