package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

var serBackup service.SerBackup

func Snapshot() (error, *kk_etcd_models.SnapshotResponse) {

	pBFile, err := serBackup.Snapshot(global_model.GlobalStage)
	if err != nil {
		return nil, err
	} else {
		return pBFile, nil
	}
}

func SnapshotRestore() (cmdStr string, err error) {
	return serBackup.SnapshotRestore(global_model.GlobalStage)
}

func SnapshotInfo(fileByte []byte) (info string, err error) {

	snapshotInfo, err := serBackup.SnapshotInfo(global_model.GlobalStage, &kk_etcd_models.SnapshotInfoParam{
		File: fileByte,
	})
	if err != nil {
		return "", err
	}
	return snapshotInfo, nil
}

func AllKVsBackup() (error, *kk_etcd_models.AllKVsBackupResponse) {
	return serBackup.AllKVsBackup()
}

// AllKVsRestore will overwrite exist kv
func AllKVsRestore(jsonBytes []byte) (err error) {

	pbFile := kk_etcd_models.AllKVsRestoreParam{
		File: jsonBytes,
	}
	err = serBackup.AllKVsRestore(global_model.GlobalStage, &pbFile)
	if err != nil {
		return err
	}
	return nil
}
