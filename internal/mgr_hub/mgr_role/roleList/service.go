package roleList

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/util_role"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) service() (err error, roles *kk_etcd_models.PBListRole) {
	span := x.stage.StartTrace("service")
	defer span.End()

	list, err := global_model.GetClient(x.stage).RoleList(context.Background())
	if err != nil {
		return err, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, err := util_role.GetRole(x.stage, roleName)
		if err != nil {
			return err, nil
		}
		roles.List = append(roles.List, role)
	}
	return nil, roles
}
