package userDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"
)

func (x *api) Handler() (*api_def.UserDelete_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	if x.In.GetUserName() == consts.UserRoot ||
		x.In.GetUserName() == global_model.GetLoginUser(x.stage).UserName {
		return nil, errors.New("illegal delete root or current logged in user")
	}
	err := x.service(x.In.GetUserName())
	return &api_def.UserDelete_Output{}, err
}
