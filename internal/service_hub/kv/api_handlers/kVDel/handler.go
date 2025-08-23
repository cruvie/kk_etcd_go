package kVDel

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.KVDel_Output, error) {

	err := x.Service(stage)
	return &api_def.KVDel_Output{}, err

}
