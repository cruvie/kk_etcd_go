package service

import (
	"context"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_crypto"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"log/slog"
)

func Login(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) (tokenString string, res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		msg := "illegal login root user!"
		slog.Error(msg, kk_stage.NewLog(stage).Args()...)
		return "", -1
	}
	/*
		2 User not exist
		4 Wrong username or password
	*/
	//get md5 password
	rawPassword := user.Password
	res, value := KVGet(stage, kk_etcd_const.User+user.UserName)
	if res != 1 {
		slog.Error("failed to get user kv", kk_stage.NewLog(stage).Any("UserName", user.UserName).Args()...)
		res = 2
		return
	}
	var userTemp kk_etcd_models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		slog.Error("failed to unmarshal user kv", kk_stage.NewLog(stage).Error(err).Any("UserName", user.UserName).Args()...)
		return
	}

	//validate password
	equal := kk_crypto.CheckPasswordHash(stage, userTemp.Password, rawPassword)
	if !equal {
		slog.Error("wrong password", kk_stage.NewLog(stage).Any("UserName", user.UserName).Args()...)
		res = 4
		return
	}
	//generate token
	tokenString, err := kk_jwt.GenerateToken[string](userTemp.UserName, 0)
	if err != nil {
		res = -1
		return
	}
	//put into etcd
	err = KVPut(stage, kk_etcd_const.Jwt+userTemp.UserName, tokenString)
	if err != nil {
		res = -1
		return
	} else {
		res = 1
		return
	}
}
func Logout(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	res = KVDel(stage, kk_etcd_const.Jwt+user.UserName)
	if res != 1 {
		res = -1
		return
	} else {
		res = 1
		return
	}
}

func UserAdd(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {

		slog.Error("illegal add root user!", kk_stage.NewLog(stage).Args()...)
		return -1
	}
	user.Password, _ = kk_crypto.GeneratePassword(stage, user.Password)
	jsonData, err := json.Marshal(&user)
	if err != nil {

		slog.Error("error in marshal user", kk_stage.NewLog(stage).Error(err).Args()...)
		return -1
	}
	jsonStr := string(jsonData)
	//add user kv to etcd used for user login
	err = KVPut(stage, kk_etcd_const.User+user.UserName, jsonStr)
	if err != nil {
		return -1
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), user.UserName, user.Password)
	if err != nil && err.Error() != "etcdserver: user name already exists" {

		slog.Error("failed to add user to etcd user", kk_stage.NewLog(stage).Error(err).Any("UserName", user.UserName).Args()...)
		return -1
	}
	return 1
}

func UserDelete(stage *kk_stage.Stage, userName string, admin bool) (res int) {
	if (!admin) && (userName == kk_etcd_const.UserRoot || userName == config.Config.Admin.UserName || userName == global_model.GetLoginUser(stage.GinCtx).UserName) {
		slog.Error("illegal delete root admin or current logged in user!")
		return 2
	}

	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), kk_etcd_const.User+userName)
	if err != nil {

		slog.Error("failed to delete user kv", kk_stage.NewLog(stage).Error(err).Any("UserName", userName).Args()...)
		return -1
	}

	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), userName)
	if err != nil {

		slog.Error("failed to delete user", kk_stage.NewLog(stage).Error(err).Any("UserName", userName).Args()...)
		return -1
	}
	return 1
}

func GetUser(stage *kk_stage.Stage, userName string) (user *kk_etcd_models.PBUser, res int) {
	rolesResp, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), userName)
	if err != nil {
		slog.Error("failed to get user", kk_stage.NewLog(stage).Error(err).Any("UserName", userName).Args()...)
		return nil, -1
	}
	user = &kk_etcd_models.PBUser{}
	user.UserName = userName
	user.Roles = rolesResp.Roles
	return user, 1
}

func UserList(stage *kk_stage.Stage) (res int, users *kk_etcd_models.PBListUser) {
	list, err := kk_etcd_client.EtcdClient.UserList(context.Background())
	if err != nil {
		slog.Error("failed to get user list", kk_stage.NewLog(stage).Error(err).Args()...)
		return -1, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, res := GetUser(stage, userName)
		if res != 1 {
			slog.Error("failed to get user", kk_stage.NewLog(stage).Error(err).Any("UserName", userName).Args()...)
			return -1, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return 1, users
}
func UserGrantRole(stage *kk_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		slog.Error("Illegal modification of root user's role!")
		return -1
	}
	toolUser.deleteAllRoles(stage, user.UserName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.UserName, role)
		if err != nil {
			slog.Error("failed to grant role", kk_stage.NewLog(stage).Error(err).Any("UserName", user.UserName).Args()...)
			return -3
		}
	}
	return 1
}
