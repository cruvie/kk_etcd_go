package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var hBackup handler.HBackup

func Snapshot() (error, *kk_etcd_models.SnapshotResponse) {
	return hBackup.Snapshot(kk_etcd_client.GlobalStage, nil)
}

func SnapshotRestore() (error, *kk_etcd_models.SnapshotRestoreResponse) {
	return hBackup.SnapshotRestore(kk_etcd_client.GlobalStage, nil)
}

func SnapshotInfo(param *kk_etcd_models.SnapshotInfoParam) (error, *kk_etcd_models.SnapshotInfoResponse) {
	return hBackup.SnapshotInfo(kk_etcd_client.GlobalStage, param)
}

func AllKVsBackup() (error, *kk_etcd_models.AllKVsBackupResponse) {
	return hBackup.AllKVsBackup(kk_etcd_client.GlobalStage, nil)
}

// AllKVsRestore will overwrite exist kv
func AllKVsRestore(param *kk_etcd_models.AllKVsRestoreParam) (error, *kk_etcd_models.AllKVsRestoreResponse) {
	return hBackup.AllKVsRestore(kk_etcd_client.GlobalStage, param)
}
