package global_model

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

const globalUser = "globalUser"

// SetLoginUser set current login user to gin context
func SetLoginUser(c *gin.Context, user *kk_etcd_models.PBUser) {
	c.Set(globalUser, user)
}

// GetLoginUser get current login user from gin context
func GetLoginUser(c *gin.Context) *kk_etcd_models.PBUser {
	value, exists := c.Get(globalUser)
	if !exists {
		return &kk_etcd_models.PBUser{}
	}
	return value.(*kk_etcd_models.PBUser)
}
