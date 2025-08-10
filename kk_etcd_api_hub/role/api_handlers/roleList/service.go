package roleList

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/util_role"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Service(stage *kk_stage.Stage) (roles *kk_etcd_models.PBListRole, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	list, err := global_model.GetClient(stage).RoleList(context.Background())
	if err != nil {
		return nil, err
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, err := util_role.GetRole(stage, roleName)
		if err != nil {
			return nil, err
		}
		roles.List = append(roles.List, role)
	}
	return roles, nil
}
