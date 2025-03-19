package userList

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) service() (err error, users *kk_etcd_models.PBListUser) {
	span := x.stage.StartTrace("service")
	defer span.End()

	list, err := global_model.GetClient(x.stage).UserList(context.Background())
	if err != nil {
		return err, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, err := util_user.GetUser(x.stage, userName)
		if err != nil {
			return err, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return nil, users
}
