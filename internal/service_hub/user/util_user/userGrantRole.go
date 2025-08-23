package util_user

import (
	"context"
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_slice"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func UserGrantRole(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	oldUser, err := GetUser(stage, user.GetUserName())
	if err != nil {
		return err
	}

	del, add := kk_slice.Diff(oldUser.GetRoles(), user.GetRoles())
	for _, role := range del {
		_, err := global_model.GetClient(stage).UserRevokeRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return errors.Join(err, errors.New("failed to revoke role"))
		}
	}

	for _, role := range add {
		_, err := global_model.GetClient(stage).UserGrantRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return err
		}
	}

	return nil
}
