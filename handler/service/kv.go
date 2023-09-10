package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
	"log"
	"strings"
)

func KVPut(key string, value string) (res int) {
	_, err := kk_etcd_client.EtcdClient.Put(context.Background(), key, value)
	if err != nil {
		log.Println("failed to put kv:", key, err)
		return -1
	}
	return 1
}

func KVGet(key string) (res int, value []byte) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil || getResponse.Kvs == nil {
		log.Println("failed to get kv:", key, err)
		return -1, nil
	}
	return 1, getResponse.Kvs[0].Value
}

func KVDel(key string) (res int) {
	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), key)
	if err != nil {
		log.Println("failed to delete kv:", key, err)
		return -1
	}
	return 1
}

func KVList(prefix string) (res int, list *models.PBListKV) {
	list = &models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Println("failed to get config list:", err)
		return -1, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &models.PBKV{
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
			(strings.HasPrefix(cfg.Key, key_prefix.User) ||
				strings.HasPrefix(cfg.Key, key_prefix.Jwt) ||
				strings.HasPrefix(cfg.Key, key_prefix.Config) ||
				strings.HasPrefix(cfg.Key, key_prefix.ServiceHttp) ||
				strings.HasPrefix(cfg.Key, key_prefix.ServiceGrpc)) {
			continue
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return 1, list
}
