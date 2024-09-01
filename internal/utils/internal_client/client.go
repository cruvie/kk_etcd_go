package internal_client

import (
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type InitClientConfig struct {
	Endpoints []string
	UserName  string
	Password  string
	DebugMode bool
}

func (x *InitClientConfig) check() {
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
	cfg.check()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    cfg.UserName,
		Password:    cfg.Password,
	})
	if err != nil {
		return err
	}
	global_model.SetClient(GlobalStage, client)

	var serUser service.SerUser
	user, err := serUser.GetUser(GlobalStage, cfg.UserName)
	if err != nil {
		return err
	}

	global_model.SetLoginUser(GlobalStage, user)

	return nil
}
