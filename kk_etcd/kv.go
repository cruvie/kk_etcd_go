package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
)

type MgrKV struct{}

func NewMgrKV() *MgrKV {
	return &MgrKV{}
}

//func (*MgrKV) KVPut(param *api_def2.KVPut_Input) (*api_def2.KVPut_Output, error) {
//	api := kVPut.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}
//
//func (*MgrKV) KVGet(param *api_def2.KVGet_Input) (*api_def2.KVGet_Output, error) {
//	api := kVGet.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}
//
//func (*MgrKV) KVDel(param *api_def2.KVDel_Input) (*api_def2.KVDel_Output, error) {
//	api := kVDel.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}
//
//func (*MgrKV) KVList(param *api_def2.KVList_Input) (*api_def2.KVList_Output, error) {
//	api := kVList.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}

// GetYaml get yaml from etcd and unmarshal to structPtr
// eg: GetYaml("configKey", &Config)
func (*MgrKV) GetYaml(key string, structPtr any) error {
	return util_kv.GetYaml(internal_client.GlobalStage, key, structPtr)
}

// PutYaml put struct to etcd in yaml format, key should not exist
// recommend: use PutJson instead which is more efficient
func (*MgrKV) PutYaml(key string, structPtr any) error {
	return util_kv.PutYaml(internal_client.GlobalStage, key, structPtr)
}

// UpdateYaml update struct in etcd, key should exist
func (*MgrKV) UpdateYaml(key string, structPtr any) error {
	return util_kv.UpdateYaml(internal_client.GlobalStage, key, structPtr)
}

func (*MgrKV) PutExistUpdateYaml(key string, structPtr any) error {
	return util_kv.PutExistUpdateYaml(internal_client.GlobalStage, key, structPtr)
}

// GetJson get json from etcd and unmarshal to structPtr
// eg: GetJson("configKey", &Config)
func (*MgrKV) GetJson(key string, structPtr any) error {
	return util_kv.GetJson(global_model.GetClient(internal_client.GlobalStage), key, structPtr)
}

// PutJson put struct to etcd in json format, key should not exist
func (*MgrKV) PutJson(key string, structPtr any) error {
	return util_kv.PutJson(internal_client.GlobalStage, key, structPtr)
}

// UpdateJson update struct in etcd, key should exist
func (*MgrKV) UpdateJson(key string, structPtr any) error {
	return util_kv.UpdateJson(internal_client.GlobalStage, key, structPtr)
}

func (*MgrKV) PutExistUpdateJson(key string, structPtr any) error {
	return util_kv.PutExistUpdateJson(global_model.GetClient(internal_client.GlobalStage), key, structPtr)
}

func (*MgrKV) CheckKeyExist(key string) error {
	return util_kv.CheckKeyExist(internal_client.GlobalStage, key)
}
