package userList

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"

func (x *api) Handler() (*api_def.UserList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	users, err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.UserList_Output{
		ListUser: users,
	}, err
}
