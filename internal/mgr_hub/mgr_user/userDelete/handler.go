package userDelete

import (
	"errors"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *api) handler() (error, *UserDelete_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	if x.in.GetUserName() == consts.UserRoot ||
		x.in.GetUserName() == global_model.GetLoginUser(x.stage).UserName {
		return errors.New("illegal delete root or current logged in user"), nil
	}
	err := x.service(x.in.GetUserName())
	return err, &UserDelete_Output{}
}
