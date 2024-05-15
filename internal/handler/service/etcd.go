package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

type SerEtcd struct{}

var serEtcd SerEtcd

func (SerEtcd) InitEtcd(stage *kk_stage.Stage, endpoints []string, userName string, password string) error {
	//https://etcd.io/docs/v3.5/op-guide/authentication/rbac/
	if config.Config.RootPassword != "" {
		cfg := clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    kk_etcd_const.UserRoot,
			Password:    config.Config.RootPassword,
		}
		err := error(nil)
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			slog.Error("etcd client init failed", kk_stage.NewLog(stage).Error(err).Args()...)
			return err
		}
	} else {
		//add root user
		cfg := clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    kk_etcd_const.UserRoot,
			Password:    kk_etcd_const.UserRoot,
		}
		err := error(nil)
		kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
		if err != nil {
			slog.Error("etcd client create failed", kk_stage.NewLog(stage).Error(err).Args()...)
			return err
		}
		if _, err := kk_etcd_client.EtcdClient.UserGet(context.Background(), kk_etcd_const.UserRoot); err != nil {
			if err.Error() == "etcdserver: user name not found" {
				if _, err := kk_etcd_client.EtcdClient.UserAdd(context.Background(), kk_etcd_const.UserRoot, kk_etcd_const.UserRoot); err != nil {
					if err.Error() != "etcdserver: user name already exists" {
						slog.Error("add etcd user failed", kk_stage.NewLog(stage).Error(err).Args()...)
						return err
					}
				}
			} else {
				slog.Error("get etcd user failed", kk_stage.NewLog(stage).Error(err).Args()...)
				return err
			}
		}

		//check root role exist
		if _, err := kk_etcd_client.EtcdClient.RoleAdd(context.Background(), kk_etcd_const.RoleRoot); err != nil {
			if err.Error() != "etcdserver: role name already exists" {
				slog.Error("add etcd role failed", kk_stage.NewLog(stage).Error(err).Args()...)
				return err
			}
		}
		//grant root role to root user
		if _, err := kk_etcd_client.EtcdClient.UserGrantRole(context.Background(), kk_etcd_const.UserRoot, kk_etcd_const.RoleRoot); err != nil {
			slog.Error("grant role to user failed", kk_stage.NewLog(stage).Error(err).Args()...)
			return err
		}
		//enable etcd auth
		if _, err := kk_etcd_client.EtcdClient.AuthEnable(context.Background()); err != nil {
			slog.Error("Enable Auth failed", kk_stage.NewLog(stage).Error(err).Args()...)
			return err
		}
		//add admin(user defined) user with root role as an administrator of the system
		user := &kk_etcd_models.PBUser{
			UserName: userName,
			Password: password,
			Roles:    []string{kk_etcd_const.RoleRoot},
		}

		//err = serUser.UserDelete(stage, user.UserName, true)
		//if err != nil {
		//	return err
		//}
		err = serUser.UserAdd(stage, user)
		if err != nil {
			if errors.Is(err, kk_etcd_error.KeyAlreadyExists) {
				err := serUser.UserUpdate(stage, user)
				if err != nil {
					slog.Error("update admin user failed", kk_stage.NewLog(stage).Error(err).Args()...)
					return err
				}
			} else {
				slog.Error("add admin user as an administrator of the system failed", kk_stage.NewLog(stage).Args()...)
				return err
			}
		}
		err = serUser.UserGrantRole(user)
		if err != nil {
			errStr := "grant " + user.Roles[0] + " role to " + user.UserName + " failed"
			slog.Error(errStr, kk_stage.NewLog(stage).Args()...)
			return err
		}
	}
	return nil
}
