package allKVsRestore

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.AllKVsRestore_Output, error) {
	err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &api_def.AllKVsRestore_Output{}, nil
}
