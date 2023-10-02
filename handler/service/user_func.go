package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"log/slog"
)

type userFunc struct{}

var toolUser userFunc

func (t *userFunc) deleteAllRoles(userName string) {
	user, _ := GetUser(userName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), userName, role)
		if err != nil {
			slog.Info("failed to revoke role", "name", userName, "err", err)
			return
		}
	}
}
