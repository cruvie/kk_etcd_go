package kVPut

import (
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/util_kv"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_kv.PutKV(x.stage, x.in.GetKey(), x.in.GetValue())
}
