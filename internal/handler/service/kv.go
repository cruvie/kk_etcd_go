package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
	"strings"
)

func KVPut(stage *kk_stage.Stage, key string, value string) (err error) {
	_, err = kk_etcd_client.EtcdClient.Put(context.Background(), key, value)
	if err != nil {
		slog.Error("failed to put kv", kk_stage.NewLog(stage).Error(err).Any("key", key).Args()...)
		return err
	}
	return nil
}

func KVGet(stage *kk_stage.Stage, key string) (res int, value []byte) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil || getResponse.Kvs == nil {

		slog.Error("failed to get kv", kk_stage.NewLog(stage).Error(err).Any("key", key).Args()...)
		return -1, nil
	}
	return 1, getResponse.Kvs[0].Value
}

func KVDel(stage *kk_stage.Stage, key string) (res int) {
	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), key)
	if err != nil {

		slog.Error("failed to delete kv", kk_stage.NewLog(stage).Error(err).Any("key", key).Args()...)
		return -1
	}
	return 1
}

func KVList(stage *kk_stage.Stage, prefix string) (res int, list *kk_etcd_models.PBListKV) {
	list = &kk_etcd_models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {

		slog.Error("failed to get config list", kk_stage.NewLog(stage).Error(err).Any("prefix", prefix).Args()...)
		return -1, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//skip all prefix match if prefix is empty
		//  User        = "kk_user/"
		//	Jwt         = "kk_jwt/"
		//	Config      = "kk_config/"
		//	ServiceHttp = "kk_service_http/"
		//	ServiceGrpc = "kk_service_grpc/"
		if prefix == "" &&
			(strings.HasPrefix(cfg.Key, kk_etcd_const.User) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.Jwt) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.Config) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.ServiceHttp) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.ServiceGrpc)) {
			continue
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return 1, list
}
