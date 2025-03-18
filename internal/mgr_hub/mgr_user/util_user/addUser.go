package util_user

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

func AddUser(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	//only need username and password
	_, err := global_model.GetClient(stage).UserAdd(
		context.Background(),
		user.GetUserName(),
		user.GetPassword(),
	)
	if err != nil && !kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCUserAlreadyExist) {
		return err
	}
	return nil
}
