package kVGet

import (
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
)

func (x *api) service() (value []byte, err error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_kv.GetKV(global_model.GetClient(x.stage), x.In.GetKey())
}
