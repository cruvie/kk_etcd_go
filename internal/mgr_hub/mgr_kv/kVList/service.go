package kVList

import (
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) service() (err error, list *kk_etcd_models.PBListKV) {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_kv.ListKV(x.stage, x.in.GetPrefix())
}
