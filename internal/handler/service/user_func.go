package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"log/slog"
)

type userFunc struct{}

var toolUser userFunc

func (t *userFunc) deleteAllRoles(stage *kku_stage.Stage, userName string) {
	user, _ := GetUser(stage, userName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), userName, role)
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
				SetAny("userName", userName).SetAny("role", role)
			msg := "failed to revoke role"
			slog.Error(msg, logBody.GetLogArgs()...)
			return
		}
	}
}
