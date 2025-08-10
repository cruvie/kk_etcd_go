package userList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.UserList_Output, error) {
	users, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &api_def.UserList_Output{
		ListUser: users,
	}, err
}
