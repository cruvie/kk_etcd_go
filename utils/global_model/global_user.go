package global_model

import (
	"github.com/gin-gonic/gin"
	"kk_etcd_go/models"
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

// ClearLoginUser clear current login user from gin context
func ClearLoginUser(c *gin.Context) {
	delete(c.Keys, globalUser)
}
