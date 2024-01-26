package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"log/slog"
)

type userFunc struct{}

var toolUser userFunc

func (t *userFunc) deleteAllRoles(stage *kk_stage.Stage, userName string) {
	user, _ := GetUser(stage, userName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), userName, role)
		if err != nil {

			msg := "failed to revoke role"
			slog.Error(msg, kk_stage.NewLog(stage).Error(err).
				Any("userName", userName).Any("role", role).Args()...)
			return
		}
	}
}
