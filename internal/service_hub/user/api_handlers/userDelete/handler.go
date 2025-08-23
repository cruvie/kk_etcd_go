package userDelete

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.UserDelete_Output, error) {

	if x.In.GetUserName() == consts.UserRoot ||
		x.In.GetUserName() == global_model.GetLoginUser(stage).UserName {
		return nil, errors.New("illegal delete root or current logged in user")
	}
	err := x.Service(stage, x.In.GetUserName())
	return &api_def.UserDelete_Output{}, err
}
