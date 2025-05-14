package logout

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"

func (x *api) Handler() (*api_def.Logout_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return &api_def.Logout_Output{}, err
}
