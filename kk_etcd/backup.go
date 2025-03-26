package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/allKVsBackup"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/allKVsRestore"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/snapshot"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/snapshotInfo"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/snapshotRestore"
)

type MgrBackup struct{}

func NewMgrBackup() *MgrBackup {
	return &MgrBackup{}
}

func (*MgrBackup) Snapshot() (*snapshot.Snapshot_Output, error) {
	api := snapshot.NewApi(internal_client.GlobalStage)
	api.In = &snapshot.Snapshot_Input{}
	return api.Handler()
}

func (*MgrBackup) SnapshotRestore() (*snapshotRestore.SnapshotRestore_Output, error) {
	api := snapshotRestore.NewApi(internal_client.GlobalStage)
	api.In = &snapshotRestore.SnapshotRestore_Input{}
	return api.Handler()
}

func (*MgrBackup) SnapshotInfo(param *snapshotInfo.SnapshotInfo_Input) (*snapshotInfo.SnapshotInfo_Output, error) {
	api := snapshotInfo.NewApi(internal_client.GlobalStage)
	api.In = param
	return api.Handler()
}

func (*MgrBackup) AllKVsBackup() (*allKVsBackup.AllKVsBackup_Output, error) {
	api := allKVsBackup.NewApi(internal_client.GlobalStage)
	api.In = &allKVsBackup.AllKVsBackup_Input{}
	return api.Handler()
}

// AllKVsRestore will overwrite exist kv
func (*MgrBackup) AllKVsRestore(param *allKVsRestore.AllKVsRestore_Input) (*allKVsRestore.AllKVsRestore_Output, error) {
	api := allKVsRestore.NewApi(internal_client.GlobalStage)
	api.In = param
	return api.Handler()
}
