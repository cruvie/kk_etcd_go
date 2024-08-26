package internal_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

func InitEtcd() {
	newLog := kk_log.NewLog(&kk_log.LogOption{})
	defer func() {
		err := kk_etcd_client.EtcdClient.Close()
		if err != nil {
			slog.Error("close etcd client failed", newLog.Error(err).Args()...)
		}
		kk_etcd_client.EtcdClient = nil
	}()

	cfg := clientv3.Config{
		Endpoints:   []string{config.Config.Etcd.Endpoint},
		DialTimeout: 5 * time.Second,
		Username:    kk_etcd_const.UserRoot,
		Password:    config.Config.RootPassword,
	}
	if toolEtcd.checkAuthEnabled() {
		var err error
		//refresh client
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			slog.Error("New etcd client failed", newLog.Error(err).Args()...)
			panic(err)
		}
		toolEtcd.initRootRolePermission()
	} else {
		toolEtcd.initRootRolePermission()
		//enable etcd auth
		_, err := kk_etcd_client.EtcdClient.AuthEnable(context.Background())
		if err != nil {
			slog.Error("Enable Auth failed", newLog.Error(err).Args()...)
			panic(err)
		}
		//refresh client
		//kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		//if err != nil {
		//	panic(err)
		//}
	}
	{
		////add admin(user defined) user with root role as an administrator of the system
		//user := &kk_etcd_models.PBUser{
		//	UserName: c.UserName,
		//	Password: c.Password,
		//	Roles:    []string{kk_etcd_const.RoleRoot},
		//}
		//
		//err := serUser.UserAdd(user)
		//if err != nil {
		//	slog.Error("add admin user as an administrator of the system failed", newLog.Args()...)
		//	panic(err)
		//}
		//
		//err = serUser.UserGrantRole(user)
		//if err != nil {
		//	errStr := "grant root role to " + user.UserName + " failed"
		//	slog.Error(errStr, newLog.Error(err).Args()...)
		//	panic(err)
		//}
	}
}
