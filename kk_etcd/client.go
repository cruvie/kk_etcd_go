package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/util_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
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

type CloseFunc func() error

func InitClient(cfg *InitClientConfig) (CloseFunc, error) {
	cfg.Check()
	internal_client.InitGlobalStage(cfg.DebugMode)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    cfg.UserName,
		Password:    cfg.Password,
	})
	if err != nil {
		return nil, err
	}
	global_model.SetClient(internal_client.GlobalStage, client)

	user, err := util_user.GetUser(internal_client.GlobalStage, cfg.UserName)
	if err != nil {
		return nil, err
	}

	global_model.SetLoginUser(internal_client.GlobalStage, user)
	closeFunc := func() error {
		err := GetClient().Close()
		internal_client.CloseGlobalStage()
		return err
	}
	return closeFunc, nil
}

func GetClient() *clientv3.Client {
	return global_model.GetClient(internal_client.GlobalStage)
}
