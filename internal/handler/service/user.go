package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_crypto"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_slice"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/authpb"
)

type SerUser struct{}

var serUser SerUser

func (SerUser) Login(stage *kk_stage.Stage, param *kk_etcd_models.LoginParam) (tokenString string, err error) {
	//get md5 password
	rawPassword := param.GetPassword()

	userTemp := kk_etcd_models.NewPBUser(param.GetUserName())
	extPBUser := NewExtPBUser(userTemp)
	err = extPBUser.Load()
	if err != nil {
		return "", err
	}

	//validate password
	err = kk_crypto.CheckPasswordHash(userTemp.GetPassword(), rawPassword)
	if err != nil {
		return "", errors.New("wrong password")
	}
	//generate token
	tokenString, err = kk_jwt.GenerateToken[string](kk_jwt.JwtPayload[string]{
		Username: userTemp.GetUserName(),
	})
	if err != nil {
		return "", err
	}
	//put into etcd
	err = extPBUser.JwtSet(tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (SerUser) Logout(stage *kk_stage.Stage, _ *kk_etcd_models.LogoutParam) error {
	userTemp := kk_etcd_models.NewPBUser(global_model.GetLoginUser(stage).GetUserName())
	extPBUser := NewExtPBUser(userTemp)
	err := extPBUser.JwtDel()
	if err != nil {
		return err
	}
	return nil
}
func (s SerUser) UserAdd(user *kk_etcd_models.PBUser) error {
	//only need username and password
	err := NewExtPBUser(&kk_etcd_models.PBUser{
		UserName: user.GetUserName(),
		Password: user.GetPassword(),
	}).Store()
	if err != nil {
		return err
	}
	return nil
}
func (s SerUser) UserUpdate(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	err := NewExtPBUser(user).Update()
	if err != nil {
		return err
	}
	return nil
}
func (s SerUser) UserDelete(userName string) error {
	pbUser := kk_etcd_models.NewPBUser(userName)
	err := NewExtPBUser(pbUser).Delete()
	return err
}
func (s SerUser) GetUser(userName string) (pbUser *kk_etcd_models.PBUser, err error) {
	pbUser = kk_etcd_models.NewPBUser(userName)
	err = NewExtPBUser(pbUser).Load()
	return pbUser, err
}

func (SerUser) UserList() (err error, users *kk_etcd_models.PBListUser) {
	list, err := kk_etcd_client.EtcdClient.UserList(context.Background())
	if err != nil {
		return err, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, err := serUser.GetUser(userName)
		if err != nil {
			return err, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return nil, users
}

func (SerUser) UserGrantRole(user *kk_etcd_models.PBUser) error {
	oldUser, err := serUser.GetUser(user.GetUserName())
	if err != nil {
		return err
	}

	del, add := kk_slice.Diff(oldUser.GetRoles(), user.GetRoles())
	for _, role := range del {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return errors.Join(err, errors.New("failed to revoke role"))
		}
	}

	for _, role := range add {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return err
		}
	}

	oldUser.Roles = user.GetRoles()
	err = serUser.UserUpdate(nil, oldUser)
	if err != nil {
		return err
	}
	return nil
}

func (SerUser) CheckRootRole(stage *kk_stage.Stage) error {
	for _, role := range global_model.GetLoginUser(stage).Roles {
		if role == kk_etcd_const.RoleRoot {
			return nil
		}
	}
	return kk_etcd_error.NoRootRole
}

func (SerUser) CheckWritePermission(stage *kk_stage.Stage) error {
	for _, roleName := range global_model.GetLoginUser(stage).Roles {
		role, err := serRole.RoleGet(roleName)
		if err != nil {
			return err
		}
		if role.PermissionType == int32(authpb.WRITE) || role.PermissionType == int32(authpb.READWRITE) {
			return nil
		}
	}
	return kk_etcd_error.PermissionDenied
}
