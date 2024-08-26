package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

type serToolEtcd struct{}

var toolEtcd serToolEtcd

func (x *serToolEtcd) checkAuthEnabled(cfg *InitKKEtcdConfig) (enabled bool) {
	c := clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: 5 * time.Second,
	}
	var err error
	kk_etcd_client.EtcdClient, err = clientv3.New(c)
	if err != nil {
		panic(err)
	}
	resp, err := kk_etcd_client.EtcdClient.AuthStatus(context.Background())
	if err != nil {
		if kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCUserEmpty) {
			//means auth enabled
			return true
		}
		panic(err)
	}
	return resp.Enabled
}

func (x *serToolEtcd) initRootRolePermission(c *InitKKEtcdConfig) error {
	//https://etcd.io/docs/v3.5/op-guide/authentication/rbac/
	newLog := kk_log.NewLog(&kk_log.LogOption{})
	user := &kk_etcd_models.PBUser{
		UserName: kk_etcd_const.UserRoot,
		Password: c.RootPassword,
		Roles:    []string{kk_etcd_const.RoleRoot},
	}
	//check root user exist
	err := serUser.UserAdd(user)
	if err != nil {
		slog.Error("add root user failed", newLog.Error(err).Args()...)
		return err
	}

	//check root role exist
	err = serRole.RoleAdd(&kk_etcd_models.RoleAddParam{
		Name: kk_etcd_const.RoleRoot,
	})
	if err != nil {
		slog.Error("add etcd role failed", newLog.Error(err).Args()...)
		return err
	}

	//grant root role to root user
	err = serUser.UserGrantRole(user)
	if err != nil {
		slog.Error("grant root role to root user failed", newLog.Error(err).Args()...)
		return err
	}
	return nil
}
