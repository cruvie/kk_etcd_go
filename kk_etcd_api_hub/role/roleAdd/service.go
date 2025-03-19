package roleAdd

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/util_role"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_role.AddRole(x.stage, x.In.GetName())
}
