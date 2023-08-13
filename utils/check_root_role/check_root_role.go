package check_root_role

import (
	"github.com/cruvie/kk_etcd_go/utils/global_model"
	"github.com/gin-gonic/gin"
)

func CheckRootRole(c *gin.Context) bool {
	var hasRootRole bool
	for _, role := range global_model.GetLoginUser(c).Roles {
		if role == "root" {
			hasRootRole = true
			break
		}
	}
	return hasRootRole
}
