package kk_etcd

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"gopkg.in/yaml.v3"
	"log/slog"
)

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key_prefix.Config+configKey)
	if err != nil {
		slog.Error("failed to get kv", "configKey", configKey, "err", err)
	}
	if getResponse.Kvs == nil {
		slog.Error("failed to get kv, getResponse.Kvs is nil", "configKey", configKey)
		return
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		slog.Error("failed Unmarshal config", "err", err, "config", string(getResponse.Kvs[0].Value))
	}
}

func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	err := service.RegisterService(registration)
	return err
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (serverList *kk_etcd_models.PBListServer, err error) {
	_, server, err := service.ServerList(serviceName)
	return server, err
}
