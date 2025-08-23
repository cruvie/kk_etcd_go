package myInfo

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.MyInfo_Output, error) {
	loginUser := global_model.GetLoginUser(stage)
	return &api_def.MyInfo_Output{
		UserName: loginUser.GetUserName(),
		Roles:    loginUser.GetRoles(),
	}, nil
}
