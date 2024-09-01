package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var hKV handler.HKV

func KVPut(param *kk_etcd_models.KVPutParam) (error, *kk_etcd_models.KVPutResponse) {
	return hKV.KVPut(internal_client.GlobalStage, param)
}

func KVGet(param *kk_etcd_models.KVGetParam) (error, *kk_etcd_models.KVGetResponse) {
	return hKV.KVGet(internal_client.GlobalStage, param)
}

func KVDel(param *kk_etcd_models.KVDelParam) (error, *kk_etcd_models.KVDelResponse) {
	return hKV.KVDel(internal_client.GlobalStage, param)
}
func KVList(param *kk_etcd_models.KVListParam) (error, *kk_etcd_models.KVListResponse) {
	return hKV.KVList(internal_client.GlobalStage, param)
}

var serKV service.SerKV

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(key string, structPtr any) error {
	return serKV.GetYaml(internal_client.GlobalStage, key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func PutYaml(key string, structPtr any) error {
	return serKV.PutYaml(internal_client.GlobalStage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func UpdateYaml(key string, structPtr any) error {
	return serKV.UpdateYaml(internal_client.GlobalStage, key, structPtr)
}

func PutExistUpdateYaml(key string, structPtr any) error {
	return serKV.PutExistUpdateYaml(internal_client.GlobalStage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(key string, structPtr any) error {
	return serKV.GetJson(internal_client.GlobalStage, key, structPtr)
}

// PutJson put struct to etcd in json format, key should not exist
func PutJson(key string, structPtr any) error {
	return serKV.PutJson(internal_client.GlobalStage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func UpdateJson(key string, structPtr any) error {
	return serKV.UpdateJson(internal_client.GlobalStage, key, structPtr)
}

func PutExistUpdateJson(key string, structPtr any) error {
	return serKV.PutExistUpdateJson(internal_client.GlobalStage, key, structPtr)
}
func CheckKeyExist(key string) error {
	return serKV.CheckKeyExist(internal_client.GlobalStage, key)
}
