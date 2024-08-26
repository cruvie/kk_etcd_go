package service

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type userTool struct{}

var toolUser userTool

func (t *userTool) checkFields(m *kk_etcd_models.PBUser) error {
	if m == nil {
		return errors.New("user is nil")
	}
	if m.GetUserName() == "" {
		return errors.New("userName is empty")
	}
	return nil
}
