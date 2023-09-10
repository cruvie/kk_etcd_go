package kk_etcd

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

func InitEtcd(endpoints []string, userName string, password string) {
	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "root",
	}
	err := error(nil)
	kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
	if err != nil {
		slog.Error("etcd client init failed", "err:", err)
	}
	//check root user exist
	if _, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), "root", "root"); err != nil {
		if err.Error() != "etcdserver: user name already exists" {
			slog.Error(err.Error())
		}
	}
	//check root role exist
	if _, err = kk_etcd_client.EtcdClient.RoleAdd(context.Background(), "root"); err != nil {
		if err.Error() != "etcdserver: role name already exists" {
			slog.Error(err.Error())
		}
	}
	//grant root role to root user
	if _, err = kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), "root", "root"); err != nil {
		slog.Error(err.Error())
	}
	//enable etcd auth
	if _, err = kk_etcd_client.EtcdClient.AuthEnable(context.Background()); err != nil {
		slog.Error(err.Error())
	}
	//add root(user defined) user as an administrator of the system
	user := &models.PBUser{
		UserName: userName,
		Password: password,
		Roles:    []string{consts.RoleRoot},
	}
	service.UserDelete(nil, user.UserName, true)
	res := service.UserAdd(user)
	if res != 1 {
		slog.Error("add root user as an administrator of the system failed")
	}
	res = service.UserGrantRole(user)
	if res != 1 {
		slog.Error("grant " + user.Roles[0] + " role to " + user.UserName + " failed")
	}
}
