package service

import (
	"context"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_encrypt"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
)

func Login(stage *kku_stage.Stage, user *kk_etcd_models.PBUser) (tokenString string, res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "illegal login root user!"
		slog.Error(msg, logBody.GetLogArgs()...)
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
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("UserName", user.UserName)
		slog.Error("failed to get user kv", logBody.GetLogArgs()...)
		res = 2
		return
	}
	var userTemp kk_etcd_models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", user.UserName)
		slog.Error("failed to unmarshal user kv", logBody.GetLogArgs()...)
		return
	}

	//validate password
	equal := kku_encrypt.CheckPasswordHash(stage, userTemp.Password, rawPassword)
	if !equal {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("UserName", user.UserName)
		slog.Error("wrong password", logBody.GetLogArgs()...)
		res = 4
		return
	}
	//generate token
	tokenString = kku_jwt.GenerateToken[string](stage, userTemp.UserName, 0, time.Duration(config.Config.JWT.ExpireTime)*time.Hour)
	//put into etcd
	res = KVPut(stage, kk_etcd_const.Jwt+userTemp.UserName, tokenString)
	if res != 1 {
		res = -1
		return
	} else {
		res = 1
		return
	}
}
func Logout(stage *kku_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	res = KVDel(stage, kk_etcd_const.Jwt+user.UserName)
	if res != 1 {
		res = -1
		return
	} else {
		res = 1
		return
	}
}

func UserAdd(stage *kku_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		slog.Error("illegal add root user!", logBody.GetLogArgs()...)
		return -1
	}
	user.Password, _ = kku_encrypt.GeneratePassword(stage, user.Password)
	jsonData, err := json.Marshal(&user)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		slog.Error("error in marshal user", logBody.GetLogArgs()...)
		return -1
	}
	jsonStr := string(jsonData)
	//add user kv to etcd used for user login
	res = KVPut(stage, kk_etcd_const.User+user.UserName, jsonStr)
	if res != 1 {
		return -1
	}
	_, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), user.UserName, user.Password)
	if err != nil && err.Error() != "etcdserver: user name already exists" {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", user.UserName)
		slog.Error("failed to add user to etcd user", logBody.GetLogArgs()...)
		return -1
	}
	return 1
}

func UserDelete(stage *kku_stage.Stage, userName string, admin bool) (res int) {
	if (!admin) && (userName == kk_etcd_const.UserRoot || userName == config.Config.Admin.UserName || userName == global_model.GetLoginUser(stage.GinCtx).UserName) {
		slog.Error("illegal delete root admin or current logged in user!")
		return 2
	}

	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), kk_etcd_const.User+userName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", userName)
		slog.Error("failed to delete user kv", logBody.GetLogArgs()...)
		return -1
	}

	_, err = kk_etcd_client.EtcdClient.UserDelete(context.Background(), userName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", userName)
		slog.Error("failed to delete user", logBody.GetLogArgs()...)
		return -1
	}
	return 1
}

func GetUser(stage *kku_stage.Stage, userName string) (user *kk_etcd_models.PBUser, res int) {
	rolesResp, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), userName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", userName)
		slog.Error("failed to get user", logBody.GetLogArgs()...)
		return nil, -1
	}
	user = &kk_etcd_models.PBUser{}
	user.UserName = userName
	user.Roles = rolesResp.Roles
	return user, 1
}

func UserList(stage *kku_stage.Stage) (res int, users *kk_etcd_models.PBListUser) {
	list, err := kk_etcd_client.EtcdClient.UserList(context.Background())
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		slog.Error("failed to get user list", logBody.GetLogArgs()...)
		return -1, nil
	}
	users = &kk_etcd_models.PBListUser{}
	for _, userName := range list.Users {
		user, res := GetUser(stage, userName)
		if res != 1 {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", userName)
			slog.Error("failed to get user", logBody.GetLogArgs()...)
			return -1, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return 1, users
}
func UserGrantRole(stage *kku_stage.Stage, user *kk_etcd_models.PBUser) (res int) {
	if user.UserName == kk_etcd_const.UserRoot {
		slog.Error("Illegal modification of root user's role!")
		return -1
	}
	toolUser.deleteAllRoles(stage, user.UserName)
	for _, role := range user.Roles {
		_, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), user.UserName, role)
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("UserName", user.UserName)
			slog.Error("failed to grant role", logBody.GetLogArgs()...)
			return -3
		}
	}
	return 1
}
