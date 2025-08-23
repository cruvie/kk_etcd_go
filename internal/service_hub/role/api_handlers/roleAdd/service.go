package roleAdd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/util_role"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("service")
	defer span.End()

	return util_role.AddRole(stage, x.In.GetName())
}
