package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var hKV handler.HKV

func KVPut(param *kk_etcd_models.KVPutParam) (error, *kk_etcd_models.KVPutResponse) {
	return hKV.KVPut(kk_etcd_client.GlobalStage, param)
}

func KVGet(param *kk_etcd_models.KVGetParam) (error, *kk_etcd_models.KVGetResponse) {
	return hKV.KVGet(kk_etcd_client.GlobalStage, param)
}

func KVDel(param *kk_etcd_models.KVDelParam) (error, *kk_etcd_models.KVDelResponse) {
	return hKV.KVDel(kk_etcd_client.GlobalStage, param)
}
func KVList(param *kk_etcd_models.KVListParam) (error, *kk_etcd_models.KVListResponse) {
	return hKV.KVList(kk_etcd_client.GlobalStage, param)
}

var serKV service.SerKV

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func GetYaml(key string, structPtr any) error {
	return serKV.GetYaml(kk_etcd_client.GlobalStage, key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func PutYaml(key string, structPtr any) error {
	return serKV.PutYaml(kk_etcd_client.GlobalStage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func UpdateYaml(key string, structPtr any) error {
	return serKV.UpdateYaml(kk_etcd_client.GlobalStage, key, structPtr)
}

func PutExistUpdateYaml(key string, structPtr any) error {
	return serKV.PutExistUpdateYaml(kk_etcd_client.GlobalStage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func GetJson(key string, structPtr any) error {
	return serKV.GetJson(kk_etcd_client.GlobalStage, key, structPtr)
}

// PutJson put struct to etcd in json format, key should not exist
func PutJson(key string, structPtr any) error {
	return serKV.PutJson(kk_etcd_client.GlobalStage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func UpdateJson(key string, structPtr any) error {
	return serKV.UpdateJson(kk_etcd_client.GlobalStage, key, structPtr)
}

func PutExistUpdateJson(key string, structPtr any) error {
	return serKV.PutExistUpdateJson(kk_etcd_client.GlobalStage, key, structPtr)
}
func CheckKeyExist(key string) error {
	return serKV.CheckKeyExist(kk_etcd_client.GlobalStage, key)
}
