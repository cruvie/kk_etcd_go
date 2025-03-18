package kVGet

import "github.com/cruvie/kk_etcd_go/kk_etcd_models"

func (x *api) handler() (error, *KVGet_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err, value := x.service()
	return err, &KVGet_Output{KV: &kk_etcd_models.PBKV{
		Key:   x.in.GetKey(),
		Value: string(value),
	}}
}
