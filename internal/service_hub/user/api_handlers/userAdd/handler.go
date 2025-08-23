package userAdd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.UserAdd_Output, error) {
	err := x.Service(stage, &kk_etcd_models.PBUser{
		UserName: x.In.GetUserName(),
		Password: x.In.GetPassword(),
		Roles:    x.In.GetRoles(),
	})
	return &api_def.UserAdd_Output{}, err
}
