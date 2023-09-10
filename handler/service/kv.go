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
	if err != nil {
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

func KVGetConfigList() (res int, list *models.PBListKV) {
	list = &models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key_prefix.Config, clientv3.WithPrefix())
	if err != nil {
		log.Println("failed to get config list:", err)
		return -1, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &models.PBKV{}
		key := string(kv.Key)
		split := strings.Split(key, key_prefix.Config)
		if len(split) >= 2 {
			cfg.Key = strings.Split(key, key_prefix.Config)[1]
		}
		cfg.Value = string(kv.Value)
		if err != nil {
			log.Println("failed to unmarshal config:", err)
			return -1, nil
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return 1, list
}
