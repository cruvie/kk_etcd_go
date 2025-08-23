package kVGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.KVGet_Output, error) {
	value, err := x.Service(stage)
	return &api_def.KVGet_Output{
		KV: &kk_etcd_models.PBKV{
			Key:   x.In.GetKey(),
			Value: string(value),
		},
	}, err
}
