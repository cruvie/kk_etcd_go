package getUser

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.GetUser_Output, error) {

	user, err := util_user.GetUser(stage, x.In.GetUserName())
	if err != nil {
		return nil, err
	}
	return &api_def.GetUser_Output{
		User: user,
	}, err
}
