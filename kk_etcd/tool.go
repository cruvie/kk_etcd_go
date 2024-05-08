package kk_etcd

import (
	"context"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
)

func CheckKeyExist(key string) error {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil {
		return errors.Join(errors.New("CheckKeyExist error"), err)
	}
	if len(getResponse.Kvs) == 0 {
		return ErrKeyNotFound
	} else {
		return ErrKeyAlreadyExists
	}
}
