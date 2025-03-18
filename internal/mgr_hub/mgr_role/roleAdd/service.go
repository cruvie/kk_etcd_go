package roleAdd

import (
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/util_role"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_role.AddRole(x.stage, x.in.GetName())
}
