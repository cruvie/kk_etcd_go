package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_crypto"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
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
	equal := kk_crypto.CheckPasswordHash(stage, userTemp.GetPassword(), rawPassword)
	if !equal {
		return "", errors.New("wrong password")
	}
	//generate token
	tokenString, err = kk_jwt.GenerateToken[string](kk_jwt.JwtPayload[string]{
		UserId: userTemp.GetUserName(),
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
func (s SerUser) UserAdd(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	newPassword, err := kk_crypto.GeneratePassword(stage, user.GetPassword())
	if err != nil {
		return err
	}
	pbUser := &kk_etcd_models.PBUser{
		UserName: user.GetUserName(),
		Password: newPassword,
		Roles:    user.GetRoles(),
	}
	err = NewExtPBUser(pbUser).Store()
	if err != nil {
		return err
	}
	return nil
}
func (s SerUser) UserUpdate(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) error {
	newPassword, err := kk_crypto.GeneratePassword(stage, user.GetPassword())
	if err != nil {
		return err
	}
	pbUser := &kk_etcd_models.PBUser{
		UserName: user.GetUserName(),
		Password: newPassword,
		Roles:    user.GetRoles(),
	}
	err = NewExtPBUser(pbUser).Update()
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
	//fixme not a good way to do this
	err := toolUser.deleteAllRoles(user.GetUserName())
	if err != nil {
		return err
	}
	for _, role := range user.GetRoles() {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.GetUserName(), role)
		if err != nil {
			return err
		}
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
		role, _ := serRole.RoleGet(roleName)
		if role.PermissionType == int32(authpb.WRITE) || role.PermissionType == int32(authpb.READWRITE) {
			return nil
		}
	}
	return kk_etcd_error.PermissionDenied
}
