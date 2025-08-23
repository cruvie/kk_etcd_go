package snapshotInfo

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.SnapshotInfo_Output, error) {
	info, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &api_def.SnapshotInfo_Output{Info: info}, err
}
