package util_user

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func GetUser(stage *kk_stage.Stage, userName string) (pbUser *kk_etcd_models.PBUser, err error) {

	resp, err := global_model.GetClient(stage).
		UserGet(context.Background(), userName)
	if err != nil {
		return nil, err
	}
	pbUser = &kk_etcd_models.PBUser{
		UserName: userName,
		Password: "",
		Roles:    resp.Roles,
	}
	return pbUser, err
}
