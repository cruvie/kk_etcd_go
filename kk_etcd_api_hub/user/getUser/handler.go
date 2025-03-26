package getUser

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/util_user"
)

func (x *api) Handler() (*GetUser_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	user, err := util_user.GetUser(x.stage, x.In.GetUserName())
	if err != nil {
		return nil, err
	}
	return &GetUser_Output{
		User: user,
	}, err
}
