package myInfo

import "github.com/cruvie/kk_etcd_go/internal/utils/global_model"

func (x *api) handler() (error, *MyInfo_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	loginUser := global_model.GetLoginUser(x.stage)
	return nil, &MyInfo_Output{
		UserName: loginUser.GetUserName(),
		Roles:    loginUser.GetRoles(),
	}
}
