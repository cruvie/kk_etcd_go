package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
)

func RoleAdd(role *kk_etcd_models.PBRole) (res int) {
	if role.Name == "root" {
		slog.Info("illegal add root role!")
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), role.Name)
	if err != nil {
		slog.Info("failed to add role", "name", role.Name, "err", err)
		return -1
	}
	return 1
}

func RoleGrantPermission(role *kk_etcd_models.PBRole) (res int) {
	//PermissionType at pkg authpb
	//authpb.READ 0
	//authpb.WRITE 1
	//authpb.READWRITE 2
	//todo 一经设定无法修改？？
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {
		slog.Info("failed to grant permission", "name", role.Name, "err", err)
		return -2
	}
	return 1
}

func RoleDelete(roleName string) (res int) {
	if roleName == "root" {
		slog.Info("illegal delete root role!")
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {
		slog.Info("failed to delete role", "name", roleName, "err", err)
		return -2
	}
	return 1
}
func RoleList() (res int, roles *kk_etcd_models.PBListRole) {
	list, err := kk_etcd_client.EtcdClient.RoleList(context.Background())
	if err != nil {
		slog.Info("failed to get role list", "err", err)
		return -1, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, res := RoleGet(roleName)
		if res != 1 {
			slog.Info("failed to get role", "name", roleName, "err", err)
			return -1, nil
		}
		roles.List = append(roles.List, role)
	}
	return 1, roles
}
func RoleGet(roleName string) (role *kk_etcd_models.PBRole, res int) {
	r, err := kk_etcd_client.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {
		slog.Info("failed to get role", "name", roleName, "err", err)
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
