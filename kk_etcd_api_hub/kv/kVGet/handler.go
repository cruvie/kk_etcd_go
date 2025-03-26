package kVGet

import "github.com/cruvie/kk_etcd_go/kk_etcd_models"

func (x *api) Handler() (*KVGet_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	value, err := x.service()
	return &KVGet_Output{KV: &kk_etcd_models.PBKV{
		Key:   x.In.GetKey(),
		Value: string(value),
	},
	}, err
}
