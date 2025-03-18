package userAdd

import "github.com/cruvie/kk_etcd_go/kk_etcd_models"

func (x *api) handler() (error, *UserAdd_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service(&kk_etcd_models.PBUser{
		UserName: x.in.GetUserName(),
		Password: x.in.GetPassword(),
		Roles:    x.in.GetRoles(),
	})
	return err, &UserAdd_Output{}
}
