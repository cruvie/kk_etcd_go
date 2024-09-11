package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"strings"
)

type HKV struct{}

var serKV service.SerKV

func (HKV) KVPut(stage *kk_stage.Stage, param *kk_etcd_models.KVPutParam) (error, *kk_etcd_models.KVPutResponse) {
	span := stage.StartTrace("KVPut")
	defer span.End()
	err := serKV.KVPut(stage, param.GetKey(), param.GetValue())
	return err, &kk_etcd_models.KVPutResponse{}
}

func (HKV) KVGet(stage *kk_stage.Stage, param *kk_etcd_models.KVGetParam) (error, *kk_etcd_models.KVGetResponse) {
	span := stage.StartTrace("KVGet")
	defer span.End()
	err, value := serKV.KVGet(stage, param.GetKey())
	return err, &kk_etcd_models.KVGetResponse{KV: &kk_etcd_models.PBKV{
		Key:   param.GetKey(),
		Value: string(value),
	}}
}

func (HKV) KVDel(stage *kk_stage.Stage, param *kk_etcd_models.KVDelParam) (error, *kk_etcd_models.KVDelResponse) {
	span := stage.StartTrace("KVDel")
	defer span.End()

	err := serKV.KVDel(stage, param.GetKey())
	return err, &kk_etcd_models.KVDelResponse{}
}

func (HKV) KVList(stage *kk_stage.Stage, param *kk_etcd_models.KVListParam) (error, *kk_etcd_models.KVListResponse) {
	span := stage.StartTrace("KVList")
	defer span.End()
	err, list := serKV.KVList(stage, param.GetPrefix())
	slices.DeleteFunc(list.GetListKV(), func(kv *kk_etcd_models.PBKV) bool {
		if strings.HasPrefix(kv.Key, kk_etcd_models.Server) {
			return true
		}
		return false
	})
	return err, &kk_etcd_models.KVListResponse{KVList: list}
}
