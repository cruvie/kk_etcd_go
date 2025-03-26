package roleDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
)

func (x *api) Handler() (*RoleDelete_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	if x.In.GetName() == consts.RoleRoot {
		return nil, errors.New("illegal delete root role")
	}
	err := x.service()
	return &RoleDelete_Output{}, err
}
