package kk_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func InitEtcd(stage *kk_stage.Stage, endpoints []string, userName string, password string) error {
	var serEtcd service.SerEtcd
	err := serEtcd.InitEtcd(stage, endpoints, userName, password)
	if err != nil {
		return err
	}
	global_model.InitGlobalStage(stage)
	var serUser service.SerUser
	user, err := serUser.GetUser(userName)
	if err != nil {
		return err
	}
	global_model.SetLoginUser(stage, user)

	return nil
}
