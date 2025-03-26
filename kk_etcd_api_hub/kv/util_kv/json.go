package util_kv

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
)

func PutExistUpdateJson(stage *kk_stage.Stage, key string, structPtr any) error {
	return putJson(stage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := GetKV(stage, key)
	if err != nil {
		return err
	}

	d := json.NewDecoder(bytes.NewReader(value))
	// prevent json number overflow
	d.UseNumber()
	err = d.Decode(structPtr)
	if err != nil {
		return err
	}
	return nil
}

// PutJson put struct to etcd in json format, key should not exist
func PutJson(stage *kk_stage.Stage, key string, structPtr any) error {
	err := CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.ErrKeyNotFound) {
		return err
	}
	return putJson(stage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func UpdateJson(stage *kk_stage.Stage, key string, structPtr any) error {
	err := CheckKeyExist(stage, key)
	if !errors.Is(err, kk_etcd_error.ErrKeyAlreadyExists) {
		return err
	}
	return putJson(stage, key, structPtr)
}

func putJson(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := json.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = global_model.GetClient(stage).Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}
