package login

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"

func (x *api) Handler() (*api_def.Login_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	tokenString, err := x.service()
	return &api_def.Login_Output{
		Token: tokenString,
	}, err
}
