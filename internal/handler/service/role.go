package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/client/v3"
)

type SerRole struct{}

var serRole SerRole

func (SerRole) RoleAdd(param *kk_etcd_models.RoleAddParam) error {
	_, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), param.GetName())
	if err != nil && !kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCRoleAlreadyExist) {
		return err
	}
	return nil
}
func (SerRole) RoleDelete(roleName string) error {
	_, err := kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {
		return err
	}
	return nil
}
func (SerRole) RoleGet(roleName string) (role *kk_etcd_models.PBRole, err error) {
	r, err := kk_etcd_client.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {
		return nil, err
	}
	role = &kk_etcd_models.PBRole{
		Name:  roleName,
		Perms: make([]*kk_etcd_models.Permission, len(r.Perm)),
	}
	for _, permission := range r.Perm {
		role.Perms = append(role.Perms, &kk_etcd_models.Permission{
			Key:            string(permission.Key),
			RangeEnd:       string(permission.RangeEnd),
			PermissionType: int32(permission.PermType),
		})
	}
	return role, nil
}
func (SerRole) RoleList() (err error, roles *kk_etcd_models.PBListRole) {
	list, err := kk_etcd_client.EtcdClient.RoleList(context.Background())
	if err != nil {
		return err, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, err := serRole.RoleGet(roleName)
		if err != nil {
			return err, nil
		}
		roles.List = append(roles.List, role)
	}
	return nil, roles

}

func (SerRole) RoleGrantPermission(name string, permission *kk_etcd_models.Permission) error {
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(
		context.Background(),
		name,
		permission.Key,
		permission.RangeEnd,
		clientv3.PermissionType(permission.PermissionType))
	if err != nil {
		return err
	}
	return nil
}
