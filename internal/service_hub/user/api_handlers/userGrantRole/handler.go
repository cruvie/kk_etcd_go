package userGrantRole

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.UserGrantRole_Output, error) {

	if x.In.GetUserName() == consts.UserRoot {
		return nil, errors.New("illegal modify root user's role")
	}
	err := x.Service(stage, &kk_etcd_models.PBUser{
		UserName: x.In.GetUserName(),
		Roles:    x.In.GetRoles(),
	})
	return &api_def.UserGrantRole_Output{}, err
}
