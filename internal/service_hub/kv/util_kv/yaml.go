package util_kv

import (
	"context"
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"gopkg.in/yaml.v3"
)

func PutExistUpdateYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	return putYaml(stage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func UpdateYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	err := CheckKeyExist(stage, key)
	if !errors.Is(err, utils.ErrKeyAlreadyExists) {
		return err
	}
	return putYaml(stage, key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func PutYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	err := CheckKeyExist(stage, key)
	if !errors.Is(err, utils.ErrKeyNotFound) {
		return err
	}

	return putYaml(stage, key, structPtr)
}

func putYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := yaml.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = global_model.GetClient(stage).Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := GetKV(global_model.GetClient(stage), key)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(value, structPtr)
	if err != nil {
		return err
	}
	return nil
}
