package kk_etcd

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"gopkg.in/yaml.v3"
)

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(key string, structPtr any) error {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), kk_etcd_const.Config+key)
	if err != nil {
		return err
	}
	if len(getResponse.Kvs) == 0 {
		return errors.New("get value is empty")
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutYaml put struct to etcd in yaml format
// recommend: use PutJson instead which is more efficient
func PutYaml(key string, structPtr any) error {
	value, err := yaml.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.Put(context.Background(), kk_etcd_const.Config+key, string(value))
	if err != nil {
		return err
	}
	return nil
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(key string, structPtr any) error {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), kk_etcd_const.Config+key)
	if err != nil {
		return err
	}
	if len(getResponse.Kvs) == 0 {
		return errors.New("get value is empty")
	}
	err = json.Unmarshal(getResponse.Kvs[0].Value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutJson put struct to etcd in json format
func PutJson(key string, structPtr any) error {
	value, err := json.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.Put(context.Background(), kk_etcd_const.Config+key, string(value))
	if err != nil {
		return err
	}
	return nil
}
