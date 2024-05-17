package service

import (
	"context"
	"encoding/json"
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
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

type SerUser struct{}

var serUser SerUser

func (SerUser) Login(stage *kk_stage.Stage, param *kk_etcd_models.LoginParam) (tokenString string, err error) {
	//get md5 password
	rawPassword := param.Password
	//todo get by model
	err, value := serKV.KVGet(kk_etcd_const.User + param.UserName)
	if err != nil {
		return "", err
	}
	var userTemp kk_etcd_models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		return "", err
	}

	//validate password
	equal := kk_crypto.CheckPasswordHash(stage, userTemp.Password, rawPassword)
	if !equal {
		return "", errors.New("wrong password")
	}
	//generate token
	tokenString, err = kk_jwt.GenerateToken[string](kk_jwt.JwtPayload[string]{
		UserId: userTemp.UserName,
	})
	if err != nil {
		return "", err
	}
	//put into etcd
	err = serKV.KVPut(kk_etcd_const.Jwt+userTemp.UserName, tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (SerUser) Logout(stage *kk_stage.Stage, _ *kk_etcd_models.LogoutParam) error {
	err := serKV.KVDel(kk_etcd_const.Jwt + global_model.GetLoginUser(stage).GetUserName())
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
	pbUser := kk_etcd_models.PBUser{
		UserName: user.GetUserName(),
		Password: newPassword,
		Roles:    user.GetRoles(),
	}
	err = s.Store(&pbUser)
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
	pbUser := kk_etcd_models.PBUser{
		UserName: user.GetUserName(),
		Password: newPassword,
		Roles:    user.GetRoles(),
	}
	err = s.Update(&pbUser)
	if err != nil {
		return err
	}
	return nil
}
func (s SerUser) UserDelete(stage *kk_stage.Stage, userName string) error {
	user := kk_etcd_models.NewPBUser(userName)
	err := s.Delete(stage, user)
	return err
}
func (s SerUser) GetUser(userName string) (user *kk_etcd_models.PBUser, err error) {
	user = kk_etcd_models.NewPBUser(userName)
	err = s.Load(user)
	return user, err
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

func (SerUser) Store(m *kk_etcd_models.PBUser) error {
	if err := toolUser.checkFields(m); err != nil {
		return err
	}
	err := serKV.PutJson(kk_etcd_const.User+m.GetUserName(), m)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), m.GetUserName(), m.GetPassword())
	if err != nil && !errors.Is(err, rpctypes.ErrGRPCUserAlreadyExist) {
		return err
	}
	return nil
}

func (SerUser) Update(m *kk_etcd_models.PBUser) error {
	if err := toolUser.checkFields(m); err != nil {
		return err
	}
	err := serKV.UpdateJson(kk_etcd_const.User+m.GetUserName(), m)
	return err
}

func (SerUser) Load(m *kk_etcd_models.PBUser) error {
	if err := toolUser.checkFields(m); err != nil {
		return err
	}
	err := serKV.GetJson(kk_etcd_const.User+m.GetUserName(), m)
	return err
}
func (SerUser) Delete(stage *kk_stage.Stage, m *kk_etcd_models.PBUser) error {
	if err := toolUser.checkFields(m); err != nil {
		return err
	}
	err := serKV.KVDel(kk_etcd_const.User + m.GetUserName())
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), kk_etcd_const.User+m.GetUserName())
	if err != nil {
		return err
	}
	return nil
}
func (SerUser) GetJwt(m *kk_etcd_models.PBUser) (string, error) {
	if err := toolUser.checkFields(m); err != nil {
		return "", err
	}
	err, value := serKV.KVGet(kk_etcd_const.Jwt + m.GetUserName())
	return string(value), err
}
func (SerUser) SetJwt(stage *kk_stage.Stage, m *kk_etcd_models.PBUser, token string) error {
	if err := toolUser.checkFields(m); err != nil {
		return err
	}
	err := serKV.KVPut(kk_etcd_const.Jwt+m.GetUserName(), token)
	return err
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
