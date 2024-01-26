package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
)

func RoleAdd(stage *kk_stage.Stage, role *kk_etcd_models.PBRole) (res int) {
	if role.Name == kk_etcd_const.RoleRoot {

		slog.Error("illegal add root role!", kk_stage.NewLog(stage).Args()...)
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), role.Name)
	if err != nil {

		slog.Error("failed to add role", kk_stage.NewLog(stage).Error(err).Any("roleName", role.Name).Args()...)
		return -1
	}
	return 1
}

func RoleGrantPermission(stage *kk_stage.Stage, role *kk_etcd_models.PBRole) (res int) {
	//PermissionType at pkg authpb
	//authpb.READ 0
	//authpb.WRITE 1
	//authpb.READWRITE 2
	//todo 一经设定无法修改？？
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {

		slog.Error("failed to grant permission", kk_stage.NewLog(stage).Error(err).Any("roleName", role.Name).Args()...)
		return -2
	}
	return 1
}

func RoleDelete(stage *kk_stage.Stage, roleName string) (res int) {
	if roleName == kk_etcd_const.RoleRoot {

		slog.Error("illegal delete root role!", kk_stage.NewLog(stage).Args()...)
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {

		slog.Error("failed to delete role", kk_stage.NewLog(stage).Error(err).Any("roleName", roleName).Args()...)
		return -2
	}
	return 1
}
func RoleList(stage *kk_stage.Stage) (res int, roles *kk_etcd_models.PBListRole) {
	list, err := kk_etcd_client.EtcdClient.RoleList(context.Background())
	if err != nil {

		slog.Error("failed to get role list", kk_stage.NewLog(stage).Error(err).Args()...)
		return -1, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, res := RoleGet(stage, roleName)
		if res != 1 {

			slog.Error("failed to get role", kk_stage.NewLog(stage).Any("roleName", role.Name).Args()...)
			return -1, nil
		}
		roles.List = append(roles.List, role)
	}
	return 1, roles
}
func RoleGet(stage *kk_stage.Stage, roleName string) (role *kk_etcd_models.PBRole, res int) {
	r, err := kk_etcd_client.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {

		slog.Error("failed to get role", kk_stage.NewLog(stage).Error(err).Any("roleName", roleName).Args()...)
		return nil, -1
	}
	role = &kk_etcd_models.PBRole{}
	role.Name = roleName
	//[permType:READWRITE key:"dfdd" range_end:"ewrew" ]
	if len(r.Perm) != 0 {
		role.Key = string(r.Perm[0].Key)
		role.RangeEnd = string(r.Perm[0].RangeEnd)
		role.PermissionType = int32(r.Perm[0].PermType)
	}
	return role, 1
}
