package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

func InitEtcd(endpoints []string, userName string, password string) {
	if config.Config.RootPassword != "" {
		cfg := clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    consts.UserRoot,
			Password:    config.Config.RootPassword,
		}
		err := error(nil)
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			slog.Error("etcd client init failed", "err", err)
		}
	} else {
		//add root user
		cfg := clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    consts.UserRoot,
			Password:    consts.UserRoot,
		}
		err := error(nil)
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			panic("etcd client init failed" + err.Error())
		}
		if _, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), consts.UserRoot); err != nil {
			if err.Error() == "etcdserver: user name not found" {
				if _, err := kk_etcd_client.EtcdClient.UserAdd(context.Background(), consts.UserRoot, consts.UserRoot); err != nil {
					if err.Error() != "etcdserver: user name already exists" {
						panic(err)
					}
				}
			}
		}

		//check root role exist
		if _, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), consts.RoleRoot); err != nil {
			if err.Error() != "etcdserver: role name already exists" {
				panic(err)
			}
		}
		//grant root role to root user
		if _, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), consts.UserRoot, consts.RoleRoot); err != nil {
			panic(err)
		}
		//enable etcd auth
		if _, err := kk_etcd_client.EtcdClient.AuthEnable(context.Background()); err != nil {
			panic(err)
		}
		//add root(user defined) user as an administrator of the system
		user := &kk_etcd_models.PBUser{
			UserName: userName,
			Password: password,
			Roles:    []string{consts.RoleRoot},
		}
		UserDelete(nil, user.UserName, true)
		res := UserAdd(user)
		if res != 1 {
			panic("add root user as an administrator of the system failed")
		}
		res = UserGrantRole(user)
		if res != 1 {
			errStr := "grant " + user.Roles[0] + " role to " + user.UserName + " failed"
			panic(errStr)
		}
	}
}
