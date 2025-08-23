package util_role

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func GetRole(stage *kk_stage.Stage, roleName string) (role *kk_etcd_models.PBRole, err error) {
	r, err := global_model.GetClient(stage).RoleGet(context.Background(), roleName)
	if err != nil {
		return nil, err
	}
	role = &kk_etcd_models.PBRole{
		Name:  roleName,
		Perms: make([]*kk_etcd_models.PBPermission, len(r.Perm)),
	}
	for _, permission := range r.Perm {
		role.Perms = append(role.Perms, &kk_etcd_models.PBPermission{
			Key:            string(permission.Key),
			RangeEnd:       string(permission.RangeEnd),
			PermissionType: int32(permission.PermType),
		})
	}
	return role, nil
}
