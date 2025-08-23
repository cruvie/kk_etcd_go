package kVPut

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/util_kv"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("service")
	defer span.End()

	return util_kv.PutKV(stage, x.In.GetKey(), x.In.GetValue())
}
