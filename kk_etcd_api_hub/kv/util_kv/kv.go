package util_kv

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func GetKV(stage *kk_stage.Stage, key string) (value []byte, err error) {
	getOutput, err := global_model.GetClient(stage).Get(context.Background(), key)
	if err != nil {
		return nil, err
	}
	if len(getOutput.Kvs) == 0 {
		return nil, kk_etcd_error.ErrKeyNotFound
	}
	return getOutput.Kvs[0].Value, nil
}

func PutKV(stage *kk_stage.Stage, key string, value string) error {

	_, err := global_model.GetClient(stage).
		Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}

func DelKV(stage *kk_stage.Stage, key string) error {
	_, err := global_model.GetClient(stage).Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func ListKV(stage *kk_stage.Stage, prefix string) (list *kk_etcd_models.PBListKV, err error) {

	list = &kk_etcd_models.PBListKV{}
	getOutput, err := global_model.GetClient(stage).Get(context.Background(), prefix, clientv3.WithPrefix())
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
