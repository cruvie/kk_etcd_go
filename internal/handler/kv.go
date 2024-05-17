package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HKV struct{}

var serKV service.SerKV

func (HKV) KVPut(stage *kk_stage.Stage, param *kk_etcd_models.KVPutParam) (error, *kk_etcd_models.KVPutResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckWritePermission(stage)
	if err != nil {
		return err, nil
	}
	err = serKV.KVPut(param.GetKey(), param.GetValue())
	return err, &kk_etcd_models.KVPutResponse{}
}

func (HKV) KVGet(stage *kk_stage.Stage, param *kk_etcd_models.KVGetParam) (error, *kk_etcd_models.KVGetResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err, value := serKV.KVGet(param.GetKey())
	return err, &kk_etcd_models.KVGetResponse{KV: &kk_etcd_models.PBKV{
		Key:   param.GetKey(),
		Value: string(value),
	}}
}

func (HKV) KVDel(stage *kk_stage.Stage, param *kk_etcd_models.KVDelParam) (error, *kk_etcd_models.KVDelResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckWritePermission(stage)
	if err != nil {
		return err, nil
	}
	err = serKV.KVDel(param.GetKey())
	return err, &kk_etcd_models.KVDelResponse{}
}

func (HKV) KVList(stage *kk_stage.Stage, param *kk_etcd_models.KVListParam) (error, *kk_etcd_models.KVListResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err, list := serKV.KVList(param.GetPrefix())
	return err, &kk_etcd_models.KVListResponse{KVList: list}
}
