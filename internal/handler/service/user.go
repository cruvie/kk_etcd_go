package service

import (
	"context"
	"encoding/json"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_encrypt"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
)

func Login(user *kk_etcd_models.PBUser) (tokenString string, res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		slog.Info("illegal login root user!")
		return "", -1
	}
	/*
		2 User not exist
		4 Wrong username or password
	*/
	//get md5 password
	rawPassword := user.Password
	res, value := KVGet(kk_etcd_const.User + user.UserName)
	if res != 1 {
		slog.Info("failed to get user kv", "name", user.UserName)
		res = 2
		return
	}
	var userTemp kk_etcd_models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		slog.Info("failed to unmarshal user kv", "name", user.UserName, "err", err)
		return
	}

	//validate password
	equal := kku_encrypt.CheckPasswordHash(userTemp.Password, rawPassword)
	if equal == false {
		slog.Info("wrong password", "name", user.UserName)
		res = 4
		return
	}
	//generate token
	tokenString = kku_jwt.GenerateToken[string](userTemp.UserName, 0, time.Duration(config.Config.JWT.ExpireTime)*time.Hour)
	//put into etcd
	res = KVPut(kk_etcd_const.Jwt+userTemp.UserName, tokenString)
	if res != 1 {
		slog.Info("failed to put jwt kv", "name", userTemp.UserName)
		res = -1
		return
	} else {
		res = 1
		return
	}
}
func Logout(user *kk_etcd_models.PBUser) (res int) {
	res = KVDel(kk_etcd_const.Jwt + user.UserName)
	if res != 1 {
		slog.Info("failed to del jwt kv", "name", user.UserName)
		res = -1
		return
	} else {
		res = 1
		return
	}
}

func UserAdd(user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		slog.Info("illegal add root user!")
		return -1
	}
	user.Password, _ = kku_encrypt.GeneratePassword(user.Password)
	jsonData, err := json.Marshal(&user)
	if err != nil {
		slog.Info("error in marshal user", "err", err)
		return -1
	}
	jsonStr := string(jsonData)
	//add user kv to etcd used for user login
	res = KVPut(kk_etcd_const.User+user.UserName, jsonStr)
	if res != 1 {
		slog.Info("failed to add user kv", "name", user.UserName, "err", err)
		return -1
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), user.UserName, user.Password)
	if err != nil && err.Error() != "etcdserver: user name already exists" {
		slog.Info("failed to add user to etcd user", "name", user.UserName, "err", err)
		return -1
	}
	return 1
}

func UserDelete(c *gin.Context, userName string, admin bool) (res int) {
	if (!admin) && (userName == kk_etcd_const.UserRoot || userName == config.Config.Admin.UserName || userName == global_model.GetLoginUser(c).UserName) {
		slog.Info("illegal delete root admin or current logged in user!")
		return 2
	}

	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), kk_etcd_const.User+userName)
	if err != nil {
		slog.Info("failed to delete user kv", "name", userName, "err", err)
		return -1
	}

	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), userName)
	if err != nil {
		slog.Info("failed to delete user", "name", userName, "err", err)
		return -1
	}
	return 1
}

func GetUser(userName string) (user *kk_etcd_models.PBUser, res int) {
	rolesResp, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), userName)
	if err != nil {
		slog.Info("failed to get user", "name", userName, "err", err)
		return nil, -1
	}
	user = &kk_etcd_models.PBUser{}
	user.UserName = userName
	user.Roles = rolesResp.Roles
	return user, 1
}

func UserList() (res int, users *kk_etcd_models.PBListUser) {
	list, err := kk_etcd_client.EtcdClient.UserList(context.Background())
	if err != nil {
		slog.Info("failed to get user list", "err", err)
		return -1, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, res := GetUser(userName)
		if res != 1 {
			slog.Info("failed to get user", "name", userName, "err", err)
			return -1, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return 1, users
}
func UserGrantRole(user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		slog.Info("Illegal modification of root user's role!")
		return -1
	}
	toolUser.deleteAllRoles(user.UserName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.UserName, role)
		if err != nil {
			slog.Info("failed to grant role", "name", user.UserName, "err", err)
			return -3
		}
	}
	return 1
}
