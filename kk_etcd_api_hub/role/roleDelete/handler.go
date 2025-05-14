package roleDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/role/api_def"
)

func (x *api) Handler() (*api_def.RoleDelete_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	if x.In.GetName() == consts.RoleRoot {
		return nil, errors.New("illegal delete root role")
	}
	err := x.service()
	return &api_def.RoleDelete_Output{}, err
}
