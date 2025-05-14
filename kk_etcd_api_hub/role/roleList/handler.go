package roleList

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"

func (x *api) Handler() (*api_def.RoleList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	roles, err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.RoleList_Output{
		ListRole: roles,
	}, nil
}
