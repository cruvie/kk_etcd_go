package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_slice"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

type SerUser struct{}

var serUser SerUser

func (SerUser) Login(stage *kk_stage.Stage, param *kk_etcd_models.LoginParam) (tokenString string, err error) {
	//todo gen jwt
	return "todo gen jwt from etcd", nil
}

func (SerUser) Logout(stage *kk_stage.Stage, _ *kk_etcd_models.LogoutParam) error {
	//todo remove jwt from etcd
	return nil
}
func (s SerUser) UserAdd(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	//only need username and password
	_, err := global_model.GetClient(stage).UserAdd(
		context.Background(),
		user.GetUserName(),
		user.GetPassword(),
	)
	if err != nil && !kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCUserAlreadyExist) {
		return err
	}
	return nil
}

func (s SerUser) UserDelete(stage *kk_stage.Stage, userName string) error {
	_, err := global_model.GetClient(stage).UserDelete(context.Background(), userName)
	if err != nil {
		return err
	}
	return err
}
func (s SerUser) GetUser(stage *kk_stage.Stage, userName string) (pbUser *kk_etcd_models.PBUser, err error) {
	resp, err := global_model.GetClient(stage).UserGet(context.Background(), userName)
	if err != nil {
		return nil, err
	}
	pbUser = &kk_etcd_models.PBUser{
		UserName: userName,
		Password: "",
		Roles:    resp.Roles,
	}
	return pbUser, err
}

func (SerUser) UserList(stage *kk_stage.Stage) (err error, users *kk_etcd_models.PBListUser) {
	list, err := global_model.GetClient(stage).UserList(context.Background())
	if err != nil {
		return err, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, err := serUser.GetUser(stage, userName)
		if err != nil {
			return err, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return nil, users
}

func (SerUser) UserGrantRole(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	oldUser, err := serUser.GetUser(stage, user.GetUserName())
	if err != nil {
		return err
	}

	del, add := kk_slice.Diff(oldUser.GetRoles(), user.GetRoles())
	for _, role := range del {
		_, err := global_model.GetClient(stage).UserRevokeRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return errors.Join(err, errors.New("failed to revoke role"))
		}
	}

	for _, role := range add {
		_, err := global_model.GetClient(stage).UserGrantRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return err
		}
	}

	return nil
}

func (SerUser) CheckRootRole(stage *kk_stage.Stage) error {
	for _, role := range global_model.GetLoginUser(stage).Roles {
		if role == consts.RoleRoot {
			return nil
		}
	}
	return kk_etcd_error.NoRootRole
}
