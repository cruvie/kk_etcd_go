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
		Name: roleName,
	}
	//if role.Name == kk_etcd_const.UserRoot {
	//	role.PermissionType = int32(authpb.READWRITE)
	//}
	//[permType:READWRITE key:"dfdd" range_end:"ewrew" ]
	if len(r.Perm) != 0 {
		role.Key = string(r.Perm[0].Key)
		role.RangeEnd = string(r.Perm[0].RangeEnd)
		role.PermissionType = int32(r.Perm[0].PermType)
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

func (SerRole) RoleGrantPermission(role *kk_etcd_models.PBRole) error {
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(
		context.Background(),
		role.Name,
		role.Key,
		role.RangeEnd,
		clientv3.PermissionType(role.PermissionType))
	if err != nil {
		return err
	}
	return nil
}
