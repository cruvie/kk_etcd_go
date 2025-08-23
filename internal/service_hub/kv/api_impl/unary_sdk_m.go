package api_impl

import (
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
)

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func (x *CServer) GetYaml(key string, structPtr any) error {
	return util_kv.GetYaml(internal_client.GlobalStage, key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func (x *CServer) PutYaml(key string, structPtr any) error {
	return util_kv.PutYaml(internal_client.GlobalStage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func (x *CServer) UpdateYaml(key string, structPtr any) error {
	return util_kv.UpdateYaml(internal_client.GlobalStage, key, structPtr)
}

func (x *CServer) PutExistUpdateYaml(key string, structPtr any) error {
	return util_kv.PutExistUpdateYaml(internal_client.GlobalStage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func (x *CServer) GetJson(key string, structPtr any) error {
	return util_kv.GetJson(global_model.GetClient(internal_client.GlobalStage), key, structPtr)
}

// PutJson put struct to etcd in json format, key should not exist
func (x *CServer) PutJson(key string, structPtr any) error {
	return util_kv.PutJson(internal_client.GlobalStage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func (x *CServer) UpdateJson(key string, structPtr any) error {
	return util_kv.UpdateJson(internal_client.GlobalStage, key, structPtr)
}

func (x *CServer) PutExistUpdateJson(key string, structPtr any) error {
	return util_kv.PutExistUpdateJson(global_model.GetClient(internal_client.GlobalStage), key, structPtr)
}

func (x *CServer) CheckKeyExist(key string) error {
	return util_kv.CheckKeyExist(internal_client.GlobalStage, key)
}
