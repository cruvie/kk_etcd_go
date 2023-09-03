package service

import (
	"context"
	"encoding/json"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"github.com/cruvie/kk_etcd_go/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_encrypt"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
)

func Login(user *models.PBUser) (tokenString string, res int) {
	if user.UserName == "root" {
		slog.Info("illegal login root user!")
		return "", -1
	}
	/*
		2 User not exist
		4 Wrong username or password
	*/
	//get md5 password
	rawPassword := user.Password
	res, value := KVGet(consts.EtcdUserPrefix + user.UserName)
	if res != 1 {
		slog.Info("failed to get user kv:", user.UserName)
		res = 2
		return
	}
	var userTemp models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		slog.Info("failed to unmarshal user kv:", user.UserName, err)
		return
	}

	//validate password
	equal := kku_encrypt.CheckPasswordHash(userTemp.Password, rawPassword)
	if equal == false {
		slog.Info("wrong password:", user.UserName)
		res = 4
		return
	}
	//generate token
	tokenString = kku_jwt.GenerateToken[string](userTemp.UserName, 0, time.Duration(config.Config.JWT.ExpireTime)*time.Hour)
	//put into etcd
	res = KVPut(consts.EtcdJwtPrefix+userTemp.UserName, tokenString)
	if res != 1 {
		slog.Info("failed to put jwt kv:", userTemp.UserName)
		res = -1
		return
	} else {
		res = 1
		return
	}
}
func Logout(user *models.PBUser) (res int) {
	res = KVDel(consts.EtcdJwtPrefix + user.UserName)
	if res != 1 {
		slog.Info("failed to del jwt kv:", user.UserName)
		res = -1
		return
	} else {
		res = 1
		return
	}
}

func UserAdd(user *models.PBUser) (res int) {
	if user.UserName == "root" {
		slog.Info("illegal add root user!")
		return -1
	}
	user.Password, _ = kku_encrypt.GeneratePassword(user.Password)
	jsonData, err := json.Marshal(&user)
	if err != nil {
		slog.Info("error in marshal user:", err)
		return -1
	}
	jsonStr := string(jsonData)
	//add user kv to etcd used for user login
	res = KVPut(consts.EtcdUserPrefix+user.UserName, jsonStr)
	if res != 1 {
		slog.Info("failed to add user kv:", user.UserName, err)
		return -1
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), user.UserName, user.Password)
	if err != nil && err.Error() != "etcdserver: user name already exists" {
		slog.Info("failed to add user to etcd user:", user.UserName, err)
		return -1
	}
	return 1
}

func UserDelete(c *gin.Context, userName string, admin bool) (res int) {
	if (!admin) && (userName == "root" || userName == config.Config.Admin.UserName || userName == global_model.GetLoginUser(c).UserName) {
		slog.Info("illegal delete root admin or current logged in user!")
		return 2
	}

	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), consts.EtcdUserPrefix+userName)
	if err != nil {
		slog.Info("failed to delete user kv:", userName, err)
		return -1
	}

	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), userName)
	if err != nil {
		slog.Info("failed to delete user:", userName, err)
		return -1
	}
	return 1
}

func GetUser(userName string) (user *models.PBUser, res int) {
	rolesResp, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), userName)
	if err != nil {
		slog.Info("failed to get user:", userName, err)
		return nil, -1
	}
	user = &models.PBUser{}
	user.UserName = userName
	user.Roles = rolesResp.Roles
	return user, 1
}

func UserList() (res int, users *models.PBListUser) {
	list, err := kk_etcd_client.EtcdClient.UserList(context.Background())
	if err != nil {
		slog.Info("failed to get user list:", err)
		return -1, nil
	}
	users = &models.PBListUser{}
	for _, userName := range list.Users {
		user, res := GetUser(userName)
		if res != 1 {
			slog.Info("failed to get user:", userName, err)
			return -1, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return 1, users
}
func UserGrantRole(user *models.PBUser) (res int) {
	if user.UserName == "root" {
		slog.Info("Illegal modification of root user's role!")
		return -1
	}
	deleteAllRoles(user.UserName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.UserName, role)
		if err != nil {
			slog.Info("failed to grant role:", user.UserName, err)
			return -3
		}
	}
	return 1
}

func deleteAllRoles(userName string) {
	user, _ := GetUser(userName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserRevokeRole(context.Background(), userName, role)
		if err != nil {
			slog.Info("failed to revoke role:", userName, err)
			return
		}
	}
}
