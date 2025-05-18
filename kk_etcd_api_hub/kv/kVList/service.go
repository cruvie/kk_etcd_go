package kVList

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (x *api) service() (list *kk_etcd_models.PBListKV, err error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	return util_kv.ListKV(x.stage, x.In.GetPrefix(), clientv3.WithPrefix())
}
