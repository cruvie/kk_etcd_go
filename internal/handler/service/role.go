package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
)

type SerRole struct{}

var serRole SerRole

func (SerRole) RoleAdd(stage *kk_stage.Stage, param *kk_etcd_models.RoleAddParam) error {
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err
	}
	if param.GetName() == kk_etcd_const.RoleRoot {
		return errors.New("illegal add root role")
	}
	_, err = kk_etcd_client.EtcdClient.RoleAdd(context.Background(), param.GetName())
	if err != nil {
		return err
	}
	return nil
}
func (SerRole) RoleDelete(stage *kk_stage.Stage, roleName string) error {
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err
	}
	if roleName == kk_etcd_const.RoleRoot {
		return errors.New("illegal delete root role")
	}
	_, err = kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
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
	role = &kk_etcd_models.PBRole{}
	role.Name = roleName
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

func (SerRole) RoleGrantPermission(stage *kk_stage.Stage, role *kk_etcd_models.PBRole) error {
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err
	}
	//PermissionType at pkg authpb
	//authpb.READ 0
	//authpb.WRITE 1
	//authpb.READWRITE 2
	//todo 一经设定无法修改？？
	_, err = kk_etcd_client.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {
		return err
	}
	return nil
}
