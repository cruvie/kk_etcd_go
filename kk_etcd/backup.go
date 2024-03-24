package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
)

func Snapshot() (pBFile *kk_pb_type.PBFile, err error) {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	pBFile, err = service.Snapshot(stage)
	if err != nil {
		return nil, err
	} else {
		return pBFile, nil
	}
}

func SnapshotRestore() (cmdStr string, err error) {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)

	return service.SnapshotRestore(stage)
}

func SnapshotInfo(fileByte []byte) (info string, err error) {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)

	snapshotInfo, err := service.SnapshotInfo(stage, &kk_pb_type.PBFile{
		Name:  "",
		Bytes: fileByte,
	})
	if err != nil {
		return "", err
	}
	return snapshotInfo, nil
}

func AllKVsBackup() (pbFile *kk_pb_type.PBFile, err error) {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	return service.AllKVsBackup(stage)
}

// AllKVsRestore will overwrite exist kv
func AllKVsRestore(jsonBytes []byte) (err error) {
	stage := kk_stage.NewStage(context.Background(), kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	pbFile := kk_pb_type.PBFile{
		Bytes: jsonBytes,
	}
	err = service.AllKVsRestore(stage, &pbFile)
	if err != nil {
		return err
	}
	return nil
}
