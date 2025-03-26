package userAdd

import "github.com/cruvie/kk_etcd_go/kk_etcd_models"

func (x *api) Handler() (*UserAdd_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service(&kk_etcd_models.PBUser{
		UserName: x.In.GetUserName(),
		Password: x.In.GetPassword(),
		Roles:    x.In.GetRoles(),
	})
	return &UserAdd_Output{}, err
}
