package userList

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Service(stage *kk_stage.Stage) (users []*kk_etcd_models.PBUser, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	list, err := global_model.GetClient(stage).UserList(context.Background())
	if err != nil {
		return nil, err
	}
	users = make([]*kk_etcd_models.PBUser, 0)
	for _, userName := range list.Users {
		user, err := util_user.GetUser(stage, userName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
