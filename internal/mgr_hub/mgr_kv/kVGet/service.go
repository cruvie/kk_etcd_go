package kVGet

import (
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/util_kv"
)

func (x *api) service() (err error, value []byte) {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_kv.GetKV(x.stage, x.in.GetKey())
}
