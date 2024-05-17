package service

import (
	"context"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type userFunc struct{}

var toolUser userFunc

func (t *userFunc) deleteAllRoles(userName string) error {
	user, _ := serUser.GetUser(userName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), userName, role)
		if err != nil {
			return errors.Join(err, errors.New("failed to revoke role"))
		}
	}
	return nil
}
func (t *userFunc) checkFields(m *kk_etcd_models.PBUser) error {
	if m == nil {
		return errors.New("user is nil")
	}
	if m.GetUserName() == "" {
		return errors.New("userName is empty")
	}
	return nil
}
