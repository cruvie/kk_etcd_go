package myInfo

import (
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"
)

func (x *api) Handler() (*api_def.MyInfo_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	loginUser := global_model.GetLoginUser(x.stage)
	return &api_def.MyInfo_Output{
		UserName: loginUser.GetUserName(),
		Roles:    loginUser.GetRoles(),
	}, nil
}
