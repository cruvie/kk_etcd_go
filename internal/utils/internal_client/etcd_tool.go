package internal_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
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

func (x *serToolEtcd) checkAuthEnabled() (enabled bool) {
	c := clientv3.Config{
		Endpoints:   []string{config.Config.Etcd.Endpoint},
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

func (x *serToolEtcd) initRootRolePermission() {
	//https://etcd.io/docs/v3.5/op-guide/authentication/rbac/
	newLog := kk_log.NewLog(&kk_log.LogOption{})
	user := &kk_etcd_models.PBUser{
		UserName: kk_etcd_const.UserRoot,
		Password: config.Config.RootPassword,
		Roles:    []string{kk_etcd_const.RoleRoot},
	}
	//add root user
	var serUser service.SerUser
	err := serUser.UserAdd(user)
	if err != nil {
		slog.Error("add root user failed", newLog.Error(err).Args()...)
		panic(err)
	}

	//add root role
	var serRole service.SerRole
	err = serRole.RoleAdd(&kk_etcd_models.RoleAddParam{
		Name: kk_etcd_const.RoleRoot,
	})
	if err != nil {
		slog.Error("add etcd role failed", newLog.Error(err).Args()...)
		panic(err)
	}

	//grant all permissions to root role
	//https://etcd.io/docs/v3.5/learning/api/
	err = serRole.RoleGrantPermission(kk_etcd_const.RoleRoot,
		&kk_etcd_models.Permission{
			Key:            `/`,
			RangeEnd:       `\0`,
			PermissionType: int32(clientv3.PermReadWrite),
		})
	if err != nil {
		slog.Error("grant all permissions to root role failed", newLog.Error(err).Args()...)
		panic(err)
	}

	//grant root role to root user
	err = serUser.UserGrantRole(user)
	if err != nil {
		slog.Error("grant root role to root user failed", newLog.Error(err).Args()...)
		panic(err)
	}
}
