package kVList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (x *Api) Service(stage *kk_stage.Stage) (list *kk_etcd_models.PBListKV, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	return util_kv.ListKV(stage, x.In.GetPrefix(), clientv3.WithPrefix())
}
