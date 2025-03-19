package userGrantRole

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) Handler() (error, *UserGrantRole_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	if x.In.GetUserName() == consts.UserRoot {
		return errors.New("illegal modify root user's role"), nil
	}
	err := x.service(&kk_etcd_models.PBUser{
		UserName: x.In.GetUserName(),
		Roles:    x.In.GetRoles(),
	})
	return err, &UserGrantRole_Output{}
}
