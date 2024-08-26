package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

type SerEtcd struct{}

var serEtcd SerEtcd

func (x *SerEtcd) InitEtcd(c *InitKKEtcdConfig) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{})
	cfg := clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    kk_etcd_const.UserRoot,
		Password:    c.RootPassword,
	}
	if toolEtcd.checkAuthEnabled(c) {
		var err error
		//refresh client
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			panic(err)
		}
		err = toolEtcd.initRootRolePermission(c)
		if err != nil {
			return err
		}
	} else {
		err := toolEtcd.initRootRolePermission(c)
		if err != nil {
			return err
		}
		//enable etcd auth
		_, err = kk_etcd_client.EtcdClient.AuthEnable(context.Background())
		if err != nil {
			slog.Error("Enable Auth failed", newLog.Error(err).Args()...)
			return err
		}
		//refresh client
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			panic(err)
		}
	}
	{
		//add admin(user defined) user with root role as an administrator of the system
		user := &kk_etcd_models.PBUser{
			UserName: c.UserName,
			Password: c.Password,
			Roles:    []string{kk_etcd_const.RoleRoot},
		}

		err := serUser.UserAdd(user)
		if err != nil {
			slog.Error("add admin user as an administrator of the system failed", newLog.Args()...)
			return err
		}

		err = serUser.UserGrantRole(user)
		if err != nil {
			errStr := "grant root role to " + user.UserName + " failed"
			slog.Error(errStr, newLog.Error(err).Args()...)
			return err
		}
	}
	return nil
}
