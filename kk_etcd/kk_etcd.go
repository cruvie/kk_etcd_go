package kk_etcd

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
	"log"
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
		log.Fatalln(err)
	}
	//check root user exist
	if _, err = kk_etcd_client.EtcdClient.UserAdd(context.Background(), "root", "root"); err != nil {
		if err.Error() != "etcdserver: user name already exists" {
			log.Fatalln(err)
		}
	}
	//check root role exist
	if _, err = kk_etcd_client.EtcdClient.RoleAdd(context.Background(), "root"); err != nil {
		if err.Error() != "etcdserver: role name already exists" {
			log.Fatalln(err)
		}
	}
	//grant root role to root user
	if _, err = kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), "root", "root"); err != nil {
		log.Fatalln(err)
	}
	//enable etcd auth
	if _, err = kk_etcd_client.EtcdClient.AuthEnable(context.Background()); err != nil {
		log.Fatalln(err)
	}
	//add root(user defined) user as an administrator of the system
	user := &models.PBUser{
		UserName: userName,
		Password: password,
		Roles:    []string{consts.RoleRoot},
	}
	service.DeleteUser(nil, user.UserName, true)
	res := service.AddUser(user)
	if res != 1 {
		log.Fatalln("add root user as an administrator of the system failed")
	}
	res = service.UserGrantRole(user.UserName, user.Roles[0])
	if res != 1 {
		log.Fatalln("grant " + user.Roles[0] + " role to " + user.UserName + " failed")
	}
}
