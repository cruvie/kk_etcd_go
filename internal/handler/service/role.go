package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
)

func RoleAdd(stage *kku_stage.Stage, role *kk_etcd_models.PBRole) (res int) {
	if role.Name == kk_etcd_const.RoleRoot {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		slog.Error("illegal add root role!", logBody.GetLogArgs()...)
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), role.Name)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("roleName", role.Name)
		slog.Error("failed to add role", logBody.GetLogArgs()...)
		return -1
	}
	return 1
}

func RoleGrantPermission(stage *kku_stage.Stage, role *kk_etcd_models.PBRole) (res int) {
	//PermissionType at pkg authpb
	//authpb.READ 0
	//authpb.WRITE 1
	//authpb.READWRITE 2
	//todo 一经设定无法修改？？
	_, err := kk_etcd_client.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("roleName", role.Name)
		slog.Error("failed to grant permission", logBody.GetLogArgs()...)
		return -2
	}
	return 1
}

func RoleDelete(stage *kku_stage.Stage, roleName string) (res int) {
	if roleName == kk_etcd_const.RoleRoot {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		slog.Error("illegal delete root role!", logBody.GetLogArgs()...)
		return -1
	}
	_, err := kk_etcd_client.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("roleName", roleName)
		slog.Error("failed to delete role", logBody.GetLogArgs()...)
		return -2
	}
	return 1
}
func RoleList(stage *kku_stage.Stage) (res int, roles *kk_etcd_models.PBListRole) {
	list, err := kk_etcd_client.EtcdClient.RoleList(context.Background())
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		slog.Error("failed to get role list", logBody.GetLogArgs()...)
		return -1, nil
	}
	roles = &kk_etcd_models.PBListRole{}
	for _, roleName := range list.Roles {
		role, res := RoleGet(stage, roleName)
		if res != 1 {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("roleName", role.Name)
			slog.Error("failed to get role", logBody.GetLogArgs()...)
			return -1, nil
		}
		roles.List = append(roles.List, role)
	}
	return 1, roles
}
func RoleGet(stage *kku_stage.Stage, roleName string) (role *kk_etcd_models.PBRole, res int) {
	r, err := kk_etcd_client.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("roleName", roleName)
		slog.Error("failed to get role", logBody.GetLogArgs()...)
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
