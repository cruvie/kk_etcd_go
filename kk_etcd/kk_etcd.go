package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type InitClientConfig struct {
	Endpoints []string
	UserName  string
	Password  string
	DebugMode bool
}

func (x *InitClientConfig) Check() {
	if len(x.Endpoints) == 0 {
		panic("endpoints is empty")
	}
	if x.UserName == "" {
		panic("userName is empty")
	}
	if x.Password == "" {
		panic("password is empty")
	}
}

func InitClient(cfg *InitClientConfig) (err error) {
	cfg.Check()
	global_stage.InitGlobalStage(cfg.DebugMode)

	kk_etcd_client.EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    cfg.UserName,
		Password:    cfg.Password,
	})
	if err != nil {
		return err
	}

	var serUser service.SerUser
	user, err := serUser.GetUser(cfg.UserName)
	if err != nil {
		return err
	}
	global_model.SetLoginUser(global_stage.GlobalStage, user)

	return nil
}
