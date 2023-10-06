package check_user

import (
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/api/v3/authpb"
)

func CheckRootRole(c *gin.Context) bool {
	var hasRootRole bool
	for _, role := range global_model.GetLoginUser(c).Roles {
		if role == kk_etcd_const.RoleRoot {
			hasRootRole = true
			break
		}
	}
	return hasRootRole
}

func CheckWritePermission(c *gin.Context) bool {
	var hasWritePermission bool
	for _, roleName := range global_model.GetLoginUser(c).Roles {
		role, _ := service.RoleGet(roleName)
		if role.PermissionType == int32(authpb.WRITE) || role.PermissionType == int32(authpb.READWRITE) {
			hasWritePermission = true
			break
		}
	}
	return hasWritePermission
}
