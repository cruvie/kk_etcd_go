package service

import (
	"context"
	"encoding/json"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"strings"
)

type SerKV struct{}

var serKV SerKV

func (SerKV) KVPut(stage *kk_stage.Stage, key string, value string) error {
	_, err := global_model.GetClient(stage).Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	return nil
}

func (SerKV) KVGet(stage *kk_stage.Stage, key string) (err error, value []byte) {
	getResponse, err := global_model.GetClient(stage).Get(context.Background(), key)
	if err != nil {
		return err, nil
	}
	if len(getResponse.Kvs) == 0 {
		return kk_etcd_error.KeyNotFound, nil
	}
	return nil, getResponse.Kvs[0].Value
}

func (SerKV) KVDel(stage *kk_stage.Stage, key string) error {
	_, err := global_model.GetClient(stage).Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (SerKV) KVList(stage *kk_stage.Stage, prefix string) (err error, list *kk_etcd_models.PBListKV) {
	list = &kk_etcd_models.PBListKV{}
	getResponse, err := global_model.GetClient(stage).Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err, nil
	}
	for _, kv := range getResponse.Kvs {
		cfg := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//skip all prefix match if prefix is empty
		//	ServiceHttp = "kk_service_http/"
		//	ServiceGrpc = "kk_service_grpc/"
		if prefix == "" &&
			(strings.HasPrefix(cfg.Key, kk_etcd_models.ServiceHttp.String()) ||
				strings.HasPrefix(cfg.Key, kk_etcd_models.ServiceGrpc.String())) {
			continue
		}
		list.ListKV = append(list.ListKV, cfg)
	}
	return nil, list
}

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func (SerKV) GetYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	err, value := serKV.KVGet(stage, key)
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
func (s SerKV) PutYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	err := s.CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.KeyNotFound) {
		return err
	}

	return toolKV.putYaml(stage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func (s SerKV) UpdateYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	err := s.CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.KeyAlreadyExists) {
		return err
	}
	return toolKV.putYaml(stage, key, structPtr)
}

func (s SerKV) PutExistUpdateYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	return toolKV.putYaml(stage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func (SerKV) GetJson(stage *kk_stage.Stage, key string, structPtr any) error {
	err, value := serKV.KVGet(stage, key)
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
func (s SerKV) PutJson(stage *kk_stage.Stage, key string, structPtr any) error {
	err := s.CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.KeyNotFound) {
		return err
	}
	return toolKV.putJson(stage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func (s SerKV) UpdateJson(stage *kk_stage.Stage, key string, structPtr any) error {
	err := s.CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.KeyAlreadyExists) {
		return err
	}
	return toolKV.putJson(stage, key, structPtr)
}
func (s SerKV) PutExistUpdateJson(stage *kk_stage.Stage, key string, structPtr any) error {
	return toolKV.putJson(stage, key, structPtr)
}

func (s SerKV) CheckKeyExist(stage *kk_stage.Stage, key string) error {
	err, _ := s.KVGet(stage, key)
	if err != nil {
		return err
	} else {
		return kk_etcd_error.KeyAlreadyExists
	}
}
