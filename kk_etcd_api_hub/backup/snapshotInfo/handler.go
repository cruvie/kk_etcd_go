package snapshotInfo

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"

func (x *api) Handler() (*api_def.SnapshotInfo_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	info, err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.SnapshotInfo_Output{Info: info}, err
}
