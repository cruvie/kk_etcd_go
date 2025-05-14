package snapshotRestore

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"

func (x *api) Handler() (*api_def.SnapshotRestore_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()
	cmdStr, err := x.service()
	if err != nil {
		return nil, err
	}
	return &api_def.SnapshotRestore_Output{CmdStr: cmdStr}, err
}
