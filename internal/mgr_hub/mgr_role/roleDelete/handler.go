package roleDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
)

func (x *api) handler() (error, *RoleDelete_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	if x.in.GetName() == consts.RoleRoot {
		return errors.New("illegal delete root role"), nil
	}
	err := x.service()
	return err, &RoleDelete_Output{}
}
