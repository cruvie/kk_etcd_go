package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HBackup struct{}

var serBackup service.SerBackup

func (HBackup) Snapshot(stage *kk_stage.Stage, _ *kk_etcd_models.SnapshotParam) (error, *kk_etcd_models.SnapshotResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err, response := serBackup.Snapshot(stage)
	return err, response
}

func (HBackup) SnapshotRestore(stage *kk_stage.Stage, _ *kk_etcd_models.SnapshotRestoreParam) (error, *kk_etcd_models.SnapshotRestoreResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	cmdStr, err := serBackup.SnapshotRestore(stage)
	return err, &kk_etcd_models.SnapshotRestoreResponse{CmdStr: cmdStr}
}
func (HBackup) SnapshotInfo(stage *kk_stage.Stage, param *kk_etcd_models.SnapshotInfoParam) (error, *kk_etcd_models.SnapshotInfoResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	info, err := serBackup.SnapshotInfo(stage, param)
	return err, &kk_etcd_models.SnapshotInfoResponse{Info: info}
}
func (HBackup) AllKVsBackup(stage *kk_stage.Stage, _ *kk_etcd_models.AllKVsBackupParam) (error, *kk_etcd_models.AllKVsBackupResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err, response := serBackup.AllKVsBackup()
	return err, response
}

func (HBackup) AllKVsRestore(stage *kk_stage.Stage, param *kk_etcd_models.AllKVsRestoreParam) (error, *kk_etcd_models.AllKVsRestoreResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serBackup.AllKVsRestore(stage, param)
	return err, &kk_etcd_models.AllKVsRestoreResponse{}
}
