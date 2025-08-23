package roleList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.RoleList_Output, error) {
	roles, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &api_def.RoleList_Output{
		ListRole: roles,
	}, nil
}
