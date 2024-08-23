package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var hBackup handler.HBackup

func Snapshot() (error, *kk_etcd_models.SnapshotResponse) {
	return hBackup.Snapshot(global_stage.GlobalStage, nil)
}

func SnapshotRestore() (error, *kk_etcd_models.SnapshotRestoreResponse) {
	return hBackup.SnapshotRestore(global_stage.GlobalStage, nil)
}

func SnapshotInfo(param *kk_etcd_models.SnapshotInfoParam) (error, *kk_etcd_models.SnapshotInfoResponse) {
	return hBackup.SnapshotInfo(global_stage.GlobalStage, param)
}

func AllKVsBackup() (error, *kk_etcd_models.AllKVsBackupResponse) {
	return hBackup.AllKVsBackup(global_stage.GlobalStage, nil)
}

// AllKVsRestore will overwrite exist kv
func AllKVsRestore(param *kk_etcd_models.AllKVsRestoreParam) (error, *kk_etcd_models.AllKVsRestoreResponse) {
	return hBackup.AllKVsRestore(global_stage.GlobalStage, param)
}
