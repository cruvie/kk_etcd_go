package kVGet

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) Handler() (*api_def.KVGet_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	value, err := x.service()
	return &api_def.KVGet_Output{KV: &kk_etcd_models.PBKV{
		Key:   x.In.GetKey(),
		Value: string(value),
	},
	}, err
}
