package global_model

import (
	"github.com/cruvie/kk_etcd_go/models"
	"github.com/gin-gonic/gin"
)

const globalUser = "globalUser"

// SetLoginUser set current login user to gin context
func SetLoginUser(c *gin.Context, user *models.PBUser) {
	c.Set(globalUser, user)
}

// GetLoginUser get current login user from gin context
func GetLoginUser(c *gin.Context) *models.PBUser {
	value, exists := c.Get(globalUser)
	if !exists {
		return &models.PBUser{}
	}
	return value.(*models.PBUser)
}
