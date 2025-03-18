package roleGet

import "github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/util_role"

func (x *api) handler() (error, *RoleGet_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	role, err := util_role.GetRole(x.stage, x.in.GetName())
	return err, &RoleGet_Output{
		Role: role,
	}
}
