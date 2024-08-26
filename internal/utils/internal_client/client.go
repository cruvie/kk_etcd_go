package internal_client

import (
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// InternalClient is the internal client for kk_etcd self use
var InternalClient *clientv3.Client

func InitInternalClient() {
	cfg := clientv3.Config{
		Endpoints:   []string{config.Config.Etcd.Endpoint},
		DialTimeout: 5 * time.Second,
		Username:    kk_etcd_const.UserRoot,
		Password:    config.Config.RootPassword,
	}
	var err error
	InternalClient, err = clientv3.New(cfg)
	if err != nil {
		panic(err)
	}
}
