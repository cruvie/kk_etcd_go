package getUser

import "github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/util_user"

func (x *api) handler() (error, *GetUser_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	user, err := util_user.GetUser(x.stage, x.in.GetUserName())
	if err != nil {
		return err, nil
	}
	return err, &GetUser_Output{
		User: user,
	}
}
