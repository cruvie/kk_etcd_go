package server_hub

//todo migrate to what kubernetes does on node update?

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/kubernetes"
	"log/slog"
)

var kc *kubernetes.Client

func InitKubernetesClient(cfg clientv3.Config) {
	var err error
	kc, err = kubernetes.New(cfg)
	if err != nil {
		panic(err)
	}
}

func CloseKubernetesClient() {
	if kc != nil {
		err := kc.Close()
		if err != nil {
			slog.Error("close kc failed", "err", err)
		}
	}
}
