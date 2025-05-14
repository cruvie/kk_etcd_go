package userGrantRole

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) Handler() (*api_def.UserGrantRole_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	if x.In.GetUserName() == consts.UserRoot {
		return nil, errors.New("illegal modify root user's role")
	}
	err := x.service(&kk_etcd_models.PBUser{
		UserName: x.In.GetUserName(),
		Roles:    x.In.GetRoles(),
	})
	return &api_def.UserGrantRole_Output{}, err
}
