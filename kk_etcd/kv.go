package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var hKV handler.HKV
var serKV service.SerKV

func KVGet(param *kk_etcd_models.KVGetParam) (error, *kk_etcd_models.KVGetResponse) {
	return hKV.KVGet(global_model.GlobalStage, param)
}

func KVPut(param *kk_etcd_models.KVPutParam) (error, *kk_etcd_models.KVPutResponse) {
	return hKV.KVPut(global_model.GlobalStage, param)
}

func KVDel(param *kk_etcd_models.KVDelParam) (error, *kk_etcd_models.KVDelResponse) {
	return hKV.KVDel(global_model.GlobalStage, param)
}

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(key string, structPtr any) error {
	return serKV.GetYaml(key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func PutYaml(key string, structPtr any) error {
	return serKV.PutYaml(key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func UpdateYaml(key string, structPtr any) error {
	return serKV.UpdateYaml(key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(key string, structPtr any) error {
	return serKV.GetJson(key, structPtr)
}

// PutJson put struct to etcd in json format, key should not exist
func PutJson(key string, structPtr any) error {
	return serKV.PutJson(key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func UpdateJson(key string, structPtr any) error {
	return serKV.UpdateJson(key, structPtr)
}
func CheckKeyExist(key string) error {
	return serKV.CheckKeyExist(key)
}
