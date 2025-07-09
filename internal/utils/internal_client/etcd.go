package internal_client

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"time"
)

func InitEtcd(stage *kk_stage.Stage) {
	newLog := kk_stage.NewLog(stage)
	InitGlobalStage()
	defer CloseGlobalStage()
	defer global_model.CloseClient(GlobalStage)

	cfg := clientv3.Config{
		Endpoints:   config.Config.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    consts.UserRoot,
		Password:    config.Config.RootPassword,
	}
	//refresh client
	client, err := clientv3.New(cfg)
	if err != nil {
		slog.Error("New etcd client failed", newLog.Error(err).Args()...)
		panic(err)
	}
	global_model.SetClient(GlobalStage, client)

	if toolEtcd.checkAuthEnabled() {
		toolEtcd.initRootRolePermission(GlobalStage)
	} else {
		toolEtcd.initRootRolePermission(GlobalStage)
		//enable etcd auth
		_, err := global_model.GetClient(GlobalStage).AuthEnable(context.Background())
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
