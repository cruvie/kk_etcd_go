package roleGet

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/util_role"
)

func (x *api) Handler() (error, *RoleGet_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	role, err := util_role.GetRole(x.stage, x.In.GetName())
	return err, &RoleGet_Output{
		Role: role,
	}
}
