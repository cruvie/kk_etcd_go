package kVGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Service(stage *kk_stage.Stage) (value []byte, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	return util_kv.GetKV(global_model.GetClient(stage), x.In.GetKey())
}
