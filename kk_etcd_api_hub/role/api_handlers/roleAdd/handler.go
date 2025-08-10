package roleAdd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.RoleAdd_Output, error) {
	err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &api_def.RoleAdd_Output{}, nil
}
