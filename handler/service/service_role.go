package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
	"log"
)

func RoleAdd(role *models.PBRole) (res int) {
	if role.Name == "root" {
		log.Println("illegal add root role!")
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), role.Name)
	if err != nil {
		log.Println("failed to add role:", role.Name, err)
		return -1
	}
	return 1
}

func RoleGrantPermission(role *models.PBRole) (res int) {
	//PermissionType at pkg authpb
	//authpb.READ 0
	//authpb.WRITE 1
	//authpb.READWRITE 2
	//todo 一经设定无法修改？？
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {
		log.Println("failed to grant permission:", role.Name, err)
		return -2
	}
	return 1
}

func RoleDelete(roleName string) (res int) {
	if roleName == "root" {
		log.Println("illegal delete root role!")
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {
		log.Println("failed to delete role:", roleName, err)
		return -2
	}
	return 1
}
func RoleList() (res int, roles *models.PBListRole) {
	list, err := kk_etcd_client.EtcdClient.RoleList(context.Background())
	if err != nil {
		log.Println("failed to get role list:", err)
		return -1, nil
	}
	roles = &models.PBListRole{}
	for _, roleName := range list.Roles {
		role, res := RoleGet(roleName)
		if res != 1 {
			log.Println(err)
			return -1, nil
		}
		roles.List = append(roles.List, role)
	}
	return 1, roles
}
func RoleGet(roleName string) (role *models.PBRole, res int) {
	r, err := kk_etcd_client.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {
		log.Println("failed to get role:", roleName, err)
		return nil, -1
	}
	role = &models.PBRole{}
	role.Name = roleName
	//[permType:READWRITE key:"dfdd" range_end:"ewrew" ]
	if len(r.Perm) != 0 {
		role.Key = string(r.Perm[0].Key)
		role.RangeEnd = string(r.Perm[0].RangeEnd)
		role.PermissionType = int32(r.Perm[0].PermType)
	}
	return role, 1
}
