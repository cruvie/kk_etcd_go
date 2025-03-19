package userAdd

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) service(user *kk_etcd_models.PBUser) error {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_user.AddUser(x.stage, user)
}
