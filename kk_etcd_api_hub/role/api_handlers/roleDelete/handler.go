package roleDelete

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.RoleDelete_Output, error) {
	if x.In.GetName() == consts.RoleRoot {
		return nil, errors.New("illegal delete root role")
	}
	err := x.Service(stage)
	return &api_def.RoleDelete_Output{}, err
}
