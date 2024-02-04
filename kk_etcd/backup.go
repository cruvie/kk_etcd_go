package kk_etcd

import (
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"log/slog"
	"os/exec"
)

func Snapshot(backupFileName string) (err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	err = service.Snapshot(stage, backupFileName)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SnapshotDownload(backupFileName string) (err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	_, err = service.SnapshotDownload(stage, backupFileName)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SnapshotRestore(backupFileName string) (err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	// 执行etcdctl命令来恢复etcd数据
	cmd := exec.Command("etcdctl", "snapshot", "restore", "/path/to/snapshot.db", "--name", "my-etcd")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to restore etcd data:", err)
		return
	}

	fmt.Println("Etcd data restore complete")

	_, err = service.SnapshotRestore(stage, backupFileName)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SnapshotInfo(backupFileName string) (info string, err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	snapshotInfo, err := service.SnapshotInfo(stage, backupFileName)
	if err != nil {
		slog.Error("Failed to get snapshot info", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}
	return snapshotInfo, nil
}

func AllKVsBackup(backupFileName string) (info string, err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	snapshotInfo, err := service.AllKVsBackup(stage, backupFileName)
	if err != nil {
		slog.Error("Failed to get snapshot info", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}
	return snapshotInfo, nil
}
func AllKVsRestore(backupFileName string) (info string, err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	snapshotInfo, err := service.AllKVsRestore(stage, backupFileName)
	if err != nil {
		slog.Error("Failed to get snapshot info", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}
	return snapshotInfo, nil
}
