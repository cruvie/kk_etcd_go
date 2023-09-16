package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/handler/service"
)

func InitEtcd(endpoints []string, userName string, password string) {
	service.InitEtcd(endpoints, userName, password)
}
