package userDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *api) Handler() (error, *UserDelete_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	if x.In.GetUserName() == consts.UserRoot ||
		x.In.GetUserName() == global_model.GetLoginUser(x.stage).UserName {
		return errors.New("illegal delete root or current logged in user"), nil
	}
	err := x.service(x.In.GetUserName())
	return err, &UserDelete_Output{}
}
