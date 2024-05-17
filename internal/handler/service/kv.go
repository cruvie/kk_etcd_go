package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"strings"
)

type SerKV struct{}

var serKV SerKV

func (SerKV) KVPut(key string, value string) error {
	_, err := kk_etcd_client.EtcdClient.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}

func (SerKV) KVGet(key string) (err error, value []byte) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key)
	if err != nil {
		return err, nil
	}
	if len(getResponse.Kvs) == 0 {
		return kk_etcd_error.KeyNotFound, nil
	}
	return nil, getResponse.Kvs[0].Value
}

func (SerKV) KVDel(key string) error {
	_, err := kk_etcd_client.EtcdClient.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (SerKV) KVList(prefix string) (err error, list *kk_etcd_models.PBListKV) {
	list = &kk_etcd_models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//skip all prefix match if prefix is empty
		//  User        = "kk_user/"
		//	Jwt         = "kk_jwt/"
		//	ServiceHttp = "kk_service_http/"
		//	ServiceGrpc = "kk_service_grpc/"
		if prefix == "" &&
			(strings.HasPrefix(cfg.Key, kk_etcd_const.User) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.Jwt) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.ServiceHttp) ||
				strings.HasPrefix(cfg.Key, kk_etcd_const.ServiceGrpc)) {
			continue
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return nil, list
}

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func (SerKV) GetYaml(key string, structPtr any) error {
	err, value := serKV.KVGet(key)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func (s SerKV) PutYaml(key string, structPtr any) error {
	err := s.CheckKeyExist(key)
	if !errors.Is(err, kk_etcd_error.KeyNotFound) {
		return err
	}

	return toolKV.putYaml(key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func (s SerKV) UpdateYaml(key string, structPtr any) error {
	err := s.CheckKeyExist(key)
	if !errors.Is(err, kk_etcd_error.KeyAlreadyExists) {
		return err
	}
	return toolKV.putYaml(key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func (SerKV) GetJson(key string, structPtr any) error {
	err, value := serKV.KVGet(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(value, structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutJson put struct to etcd in json format, key should not exist
func (s SerKV) PutJson(key string, structPtr any) error {
	err := s.CheckKeyExist(key)
	if !errors.Is(err, kk_etcd_error.KeyNotFound) {
		return err
	}
	return toolKV.putJson(key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func (s SerKV) UpdateJson(key string, structPtr any) error {
	err := s.CheckKeyExist(key)
	if !errors.Is(err, kk_etcd_error.KeyAlreadyExists) {
		return err
	}
	return toolKV.putJson(key, structPtr)
}
func (s SerKV) CheckKeyExist(key string) error {
	err, _ := s.KVGet(key)
	if err != nil {
		return err
	} else {
		return kk_etcd_error.KeyAlreadyExists
	}
}
