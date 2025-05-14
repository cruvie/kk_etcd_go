package roleAdd

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"

func (x *api) Handler() (*api_def.RoleAdd_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.RoleAdd_Output{}, nil
}
