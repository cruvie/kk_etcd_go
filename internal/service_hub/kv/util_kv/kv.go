package util_kv

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetKV(client *clientv3.Client, key string) (value []byte, err error) {
	getOutput, err := client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}
	if len(getOutput.Kvs) == 0 {
		return nil, utils.ErrKeyNotFound
	}
	return getOutput.Kvs[0].Value, nil
}

func PutKV(stage *kk_stage.Stage, key string, value string, opts ...clientv3.OpOption) error {
	_, err := global_model.GetClient(stage).
		Put(context.Background(), key, value, opts...)
	if err != nil {
		return err
	}
	return nil
}

func DelKV(stage *kk_stage.Stage, key string, opts ...clientv3.OpOption) error {
	_, err := global_model.GetClient(stage).Delete(context.Background(), key, opts...)
	if err != nil {
		return err
	}
	return nil
}

func ListKV(stage *kk_stage.Stage, prefix string, opts ...clientv3.OpOption) (list *kk_etcd_models.PBListKV, err error) {
	list = &kk_etcd_models.PBListKV{}
	getOutput, err := global_model.GetClient(stage).Get(context.Background(), prefix, opts...)
	if err != nil {
		return nil, err
	}
	for _, kv := range getOutput.Kvs {
		cfg := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return list, nil
}
