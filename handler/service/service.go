package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/client/v3"
	"kk_etcd_go/config"
	"kk_etcd_go/consts"
	"kk_etcd_go/kk_etcd"
	"kk_etcd_go/models"
	"kk_etcd_go/utils/global_model"
	"log"
	"strings"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_encrypt"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
)

func Login(user *models.PBUser) (tokenString string, res int) {
	if user.UserName == "root" {
		log.Println("illegal login root user!")
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
		log.Println("failed to get user kv:", user.UserName)
		res = 2
		return
	}
	var userTemp models.PBUser
	if err := json.Unmarshal(value, &userTemp); err != nil {
		log.Println("failed to unmarshal user kv:", user.UserName, err)
		return
	}

	//validate password
	equal := kku_encrypt.CheckPasswordHash(userTemp.Password, rawPassword)
	if equal == false {
		log.Println("wrong password:", user.UserName)
		res = 4
		return
	}
	//generate token
	tokenString = kku_jwt.GenerateToken[string](userTemp.UserName, 0, time.Duration(config.GlobalConfig.JWT.ExpireTime)*time.Hour)
	//put into etcd
	res = KVPut(consts.EtcdJwtPrefix+userTemp.UserName, tokenString)
	if res != 1 {
		log.Println("failed to put jwt kv:", userTemp.UserName)
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
		log.Println("failed to del jwt kv:", user.UserName)
		res = -1
		return
	} else {
		res = 1
		return
	}
}

func AddUser(user *models.PBUser) (res int) {
	if user.UserName == "root" {
		log.Println("illegal add root user!")
		return -1
	}
	user.Password, _ = kku_encrypt.GeneratePassword(user.Password)
	jsonData, err := json.Marshal(&user)
	if err != nil {
		log.Println("error in marshal user:", err)
		return -1
	}
	jsonStr := string(jsonData)
	//add user kv to etcd used for user login
	res = KVPut(consts.EtcdUserPrefix+user.UserName, jsonStr)
	if res != 1 {
		log.Println("failed to add user kv:", user.UserName, err)
		return -1
	}
	_, err = kk_etcd.EtcdClient.UserAdd(context.Background(), user.UserName, user.Password)
	if err != nil && err.Error() != "etcdserver: user name already exists" {
		log.Println("failed to add user to etcd user:", user.UserName, err)
		return -1
	}
	return 1
}

func DeleteUser(c *gin.Context, userName string, admin bool) (res int) {
	if (!admin) && (userName == "root" || userName == config.GlobalConfig.Admin.UserName || userName == global_model.GetLoginUser(c).UserName) {
		log.Println("illegal delete root admin or your own user!")
		return 2
	}

	_, err := kk_etcd.EtcdClient.Delete(context.Background(), consts.EtcdUserPrefix+userName)
	if err != nil {
		log.Println("failed to delete user kv:", userName, err)
		return -1
	}

	_, err = kk_etcd.EtcdClient.UserDelete(context.Background(), userName)
	if err != nil {
		log.Println("failed to delete user:", userName, err)
		return -1
	}
	return 1
}

func GetUser(userName string) (user *models.PBUser, res int) {
	rolesResp, err := kk_etcd.EtcdClient.UserGet(context.Background(), userName)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	user = &models.PBUser{}
	user.UserName = userName
	user.Roles = rolesResp.Roles
	return user, 1
}

func UserList() (res int, users *models.PBListUser) {
	list, err := kk_etcd.EtcdClient.UserList(context.Background())
	if err != nil {
		log.Println("failed to get user list:", err)
		return -1, nil
	}
	users = &models.PBListUser{}
	for _, userName := range list.Users {
		user, res := GetUser(userName)
		if res != 1 {
			log.Println(err)
			return -1, nil
		}
		users.ListUser = append(users.ListUser, user)
	}
	return 1, users
}

func AddRole(role *models.PBRole) (res int) {
	if role.Name == "root" {
		log.Println("illegal add root role!")
		return -1
	}
	_, err := kk_etcd.EtcdClient.RoleAdd(context.Background(), role.Name)
	if err != nil {
		log.Println("failed to add role:", role.Name, err)
		return -1
	}
	return 1
}

func DeleteRole(roleName string) (res int) {
	if roleName == "root" {
		log.Println("illegal delete root role!")
		return -1
	}
	_, err := kk_etcd.EtcdClient.RoleDelete(context.Background(), roleName)
	if err != nil {
		log.Println("failed to delete role:", roleName, err)
		return -1
	}
	return 1
}
func RoleList() (res int, roles *models.PBListRole) {
	list, err := kk_etcd.EtcdClient.RoleList(context.Background())
	if err != nil {
		log.Println("failed to get role list:", err)
		return -1, nil
	}
	roles = &models.PBListRole{}
	for _, roleName := range list.Roles {
		role, res := RoleGet(roleName)
		if res != 1 {
			log.Println(err)
			return -1, nil
		}
		roles.List = append(roles.List, role)
	}
	return 1, roles
}
func RoleGet(roleName string) (role *models.PBRole, res int) {
	r, err := kk_etcd.EtcdClient.RoleGet(context.Background(), roleName)
	if err != nil {
		log.Println("failed to get role:", roleName, err)
		return nil, -1
	}
	role = &models.PBRole{}
	role.Name = roleName
	//[permType:READWRITE key:"dfdd" range_end:"ewrew" ]
	role.Key = string(r.Perm[0].Key)
	role.RangeEnd = string(r.Perm[0].RangeEnd)
	role.PermissionType = int32(r.Perm[0].PermType)
	return role, 1
}
func RoleGrantPermission(role *models.PBRole) (res int) {
	if role.Name == "root" {
		log.Println("illegal change root role permission!")
		return -1
	}
	_, err := kk_etcd.EtcdClient.RoleGrantPermission(context.Background(), role.Name, role.Key, role.RangeEnd, clientv3.PermissionType(role.PermissionType))
	if err != nil {
		log.Println("failed to grant permission:", role.Name, err)
		return -1
	}
	return 1
}
func UserGrantRole(userName string, roleName string) (res int) {
	if userName == "root" {
		log.Println("illegal change root role!")
		return -1
	}
	_, err := kk_etcd.EtcdClient.UserGrantRole(context.Background(), userName, roleName)
	if err != nil {
		log.Println("failed to grant role:", userName, err)
		return -1
	}
	return 1
}

func KVPut(key string, value string) (res int) {
	_, err := kk_etcd.EtcdClient.Put(context.Background(), key, value)
	if err != nil {
		log.Println("failed to put kv:", key, err)
		return -1
	}
	return 1
}

func KVGet(key string) (res int, value []byte) {
	getResponse, err := kk_etcd.EtcdClient.Get(context.Background(), key)
	if err != nil {
		log.Println("failed to get kv:", key, err)
		return -1, nil
	}
	return 1, getResponse.Kvs[0].Value
}

func KVDel(key string) (res int) {
	_, err := kk_etcd.EtcdClient.Delete(context.Background(), key)
	if err != nil {
		log.Println("failed to delete kv:", key, err)
		return -1
	}
	return 1
}

func KVGetConfigList() (res int, list *models.PBListKV) {
	list = &models.PBListKV{}
	getResponse, err := kk_etcd.EtcdClient.Get(context.Background(), consts.EtcdConfig, clientv3.WithPrefix())
	if err != nil {
		log.Println("failed to get config list:", err)
		return -1, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &models.PBKV{}
		cfg.Key = strings.Split(string(kv.Key), consts.EtcdConfig)[1]
		cfg.Value = string(kv.Value)
		if err != nil {
			log.Println("failed to unmarshal config:", err)
			return -1, nil
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return 1, list
}
