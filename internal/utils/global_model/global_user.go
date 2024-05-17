package global_model

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

const globalUser = "globalUser"

// SetLoginUser set current login user to gin context
func SetLoginUser(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) {
	stage.Set(globalUser, user)
}

// GetLoginUser get current login user from gin context
func GetLoginUser(stage *kk_stage.Stage) *kk_etcd_models.PBUser {
	value, exists := stage.Get(globalUser)
	if !exists {
		return &kk_etcd_models.PBUser{}
	}
	return value.(*kk_etcd_models.PBUser)
}
