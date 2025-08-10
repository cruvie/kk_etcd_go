package kVGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
)

func (x *Api) Service(stage *kk_stage.Stage) (value []byte, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	return util_kv.GetKV(global_model.GetClient(stage), x.In.GetKey())
}
