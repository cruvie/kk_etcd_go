package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
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
		kku_log.SlogPanic("etcd client init failed", err)
	}
	//check root user exist
	if _, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), "root", "root"); err != nil {
		if err.Error() != "etcdserver: user name already exists" {
			kku_log.SlogPanic("", err)
		}
	}
	//check root role exist
	if _, err = kk_etcd_client.EtcdClient.RoleAdd(context.Background(), "root"); err != nil {
		if err.Error() != "etcdserver: role name already exists" {
			kku_log.SlogPanic("", err)
		}
	}
	//grant root role to root user
	if _, err = kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), "root", "root"); err != nil {
		kku_log.SlogPanic("", err)
	}
	//enable etcd auth
	if _, err = kk_etcd_client.EtcdClient.AuthEnable(context.Background()); err != nil {
		kku_log.SlogPanic("", err)
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
		kku_log.SlogPanic("add root user as an administrator of the system failed")
	}
	res = service.UserGrantRole(user)
	if res != 1 {
		kku_log.SlogPanic("grant " + user.Roles[0] + " role to " + user.UserName + " failed")
	}
}
