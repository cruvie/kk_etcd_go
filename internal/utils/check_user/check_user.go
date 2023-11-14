package check_user

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"go.etcd.io/etcd/api/v3/authpb"
)

func CheckRootRole(stage *kku_stage.Stage) bool {
	var hasRootRole bool
	for _, role := range global_model.GetLoginUser(stage.GinCtx).Roles {
		if role == kk_etcd_const.RoleRoot {
			hasRootRole = true
			break
		}
	}
	return hasRootRole
}

func CheckWritePermission(stage *kku_stage.Stage) bool {
	var hasWritePermission bool
	for _, roleName := range global_model.GetLoginUser(stage.GinCtx).Roles {
		role, _ := service.RoleGet(stage, roleName)
		if role.PermissionType == int32(authpb.WRITE) || role.PermissionType == int32(authpb.READWRITE) {
			hasWritePermission = true
			break
		}
	}
	return hasWritePermission
}
