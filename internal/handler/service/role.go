package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/client/v3"
)

type SerRole struct{}

var serRole SerRole

func (SerRole) RoleAdd(stage *kk_stage.Stage, param *kk_etcd_models.RoleAddParam) error {
	_, err := global_model.GetClient(stage).RoleAdd(context.Background(), param.GetName())
	if err != nil && !kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCRoleAlreadyExist) {
		return err
	}
	return nil
}
func (SerRole) RoleDelete(stage *kk_stage.Stage, roleName string) error {
	_, err := global_model.GetClient(stage).RoleDelete(context.Background(), roleName)
	if err != nil {
		return err
	}
	return nil
}
func (SerRole) RoleGet(stage *kk_stage.Stage, roleName string) (role *kk_etcd_models.PBRole, err error) {
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
func (SerRole) RoleList(stage *kk_stage.Stage) (err error, roles *kk_etcd_models.PBListRole) {
	list, err := global_model.GetClient(stage).RoleList(context.Background())
	if err != nil {
		return err, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, err := serRole.RoleGet(stage, roleName)
		if err != nil {
			return err, nil
		}
		roles.List = append(roles.List, role)
	}
	return nil, roles

}

func (SerRole) RoleGrantPermission(stage *kk_stage.Stage, param *kk_etcd_models.RoleGrantPermissionParam) error {
	_, err := global_model.GetClient(stage).RoleGrantPermission(
		context.Background(),
		param.GetName(),
		param.GetPerm().GetKey(),
		param.GetPerm().GetRangeEnd(),
		clientv3.PermissionType(param.GetPerm().GetPermissionType()))
	if err != nil {
		return err
	}
	return nil
}
func (SerRole) RoleRevokePermission(stage *kk_stage.Stage, param *kk_etcd_models.RoleRevokePermissionParam) error {
	_, err := global_model.GetClient(stage).RoleRevokePermission(
		context.Background(),
		param.GetName(),
		param.GetKey(),
		param.GetRangeEnd())
	if err != nil {
		return err
	}
	return nil
}
