package userAdd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Service(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	span := stage.StartTrace("service")
	defer span.End()

	return util_user.AddUser(stage, user)
}
