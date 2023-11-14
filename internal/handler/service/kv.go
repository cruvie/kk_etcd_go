package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
	"strings"
)

func KVPut(stage *kku_stage.Stage, key string, value string) (res int) {
	_, err := kk_etcd_client.EtcdClient.Put(context.Background(), key, value)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("key", key)
		slog.Error("failed to put kv", logBody.GetLogArgs()...)
		return -1
	}
	return 1
}

func KVGet(stage *kku_stage.Stage, key string) (res int, value []byte) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil || getResponse.Kvs == nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("key", key)
		slog.Error("failed to get kv", logBody.GetLogArgs()...)
		return -1, nil
	}
	return 1, getResponse.Kvs[0].Value
}

func KVDel(stage *kku_stage.Stage, key string) (res int) {
	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), key)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("key", key)
		slog.Error("failed to delete kv", logBody.GetLogArgs()...)
		return -1
	}
	return 1
}

func KVList(stage *kku_stage.Stage, prefix string) (res int, list *kk_etcd_models.PBListKV) {
	list = &kk_etcd_models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).SetAny("prefix", prefix)
		slog.Error("failed to get config list", logBody.GetLogArgs()...)
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
