package kk_etcd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"gopkg.in/yaml.v3"
)

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(key string, structPtr any) error {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil {
		return err
	}
	if len(getResponse.Kvs) == 0 {
		return ErrKeyNotFound
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func PutYaml(key string, structPtr any) error {
	err := CheckKeyExist(key)
	if !errors.Is(err, ErrKeyNotFound) {
		return err
	}

	return putYaml(key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func UpdateYaml(key string, structPtr any) error {
	err := CheckKeyExist(key)
	if !errors.Is(err, ErrKeyAlreadyExists) {
		return err
	}
	return putYaml(key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(key string, structPtr any) error {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil {
		return err
	}
	if len(getResponse.Kvs) == 0 {
		return ErrKeyNotFound
	}
	err = json.Unmarshal(getResponse.Kvs[0].Value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutJson put struct to etcd in json format, key should not exist
func PutJson(key string, structPtr any) error {
	err := CheckKeyExist(key)
	if !errors.Is(err, ErrKeyNotFound) {
		return err
	}
	return putJson(key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func UpdateJson(key string, structPtr any) error {
	err := CheckKeyExist(key)
	if !errors.Is(err, ErrKeyAlreadyExists) {
		return err
	}
	return putJson(key, structPtr)
}
