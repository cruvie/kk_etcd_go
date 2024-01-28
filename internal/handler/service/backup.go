package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"time"
)

// Snapshot get snapshot, similar to `etcdctl snapshot save snapshot.db`
func Snapshot(stage *kk_stage.Stage, backupFileName string) (err error) {
	//etcdctl --endpoints=127.0.0.1:2379 --user kk_etcd:kk_etcd snapshot save /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshot.db
	backupData, err := kk_etcd_client.EtcdClient.Snapshot(context.Background())
	defer func(backupData io.ReadCloser) {
		if backupData != nil {
			err := backupData.Close()
			if err != nil {
				slog.Error("Failed to close backup data", kk_stage.NewLog(stage).Error(err).Args()...)
			}
		}
	}(backupData)
	if err != nil {
		slog.Error("Failed to create etcd snapshot", kk_stage.NewLog(stage).Error(err).Args()...)
		return err
	}

	backupBytes, err := io.ReadAll(backupData)
	if err != nil {
		slog.Error("Failed to read backup data", kk_stage.NewLog(stage).Error(err).Args()...)
		return err
	}

	if backupFileName == "" {
		workDir, _ := os.Getwd()
		timeStr := time.Now().Format(kk_time.DateTime)
		backupFileName = workDir + "/backup/etcd_backup_" + timeStr + ".db"
	}

	err = os.WriteFile(backupFileName, backupBytes, 0644)
	if err != nil {
		slog.Error("Failed to write backup data to file", kk_stage.NewLog(stage).Error(err).Args()...)
		return err
	}
	return nil
}
func SnapshotDownload(stage *kk_stage.Stage, backupFileName string) (cmdStr string, err error) {

	return cmdStr, nil
}

// SnapshotRestore todo migrate to etcd cluster manager
func SnapshotRestore(stage *kk_stage.Stage, backupFileName string) (cmdStr string, err error) {
	//https://etcd.io/docs/v3.5/op-guide/recovery/
	cmdStr = `
			etcdutl snapshot restore /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshot.db \
		  --data-dir  /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/out \
		  --name node1 \
		  --initial-cluster node1=http://127.0.0.1:2380 \
		  --initial-advertise-peer-urls http://127.0.0.1:2380
`
	return cmdStr, nil
}

func SnapshotInfo(stage *kk_stage.Stage, backupFileName string) (info string, err error) {
	//etcdutl --write-out=table snapshot status etcd_backup.db
	cmd := exec.Command("etcdutl", "--write-out", "table", "snapshot", "status", backupFileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		s := cmd.String()
		slog.Error("failed to get snapshot info", kk_stage.NewLog(stage).Error(err).String("cmd", s).Args()...)
		return "", err
	}

	return string(output), nil
}

// todo cmk
func AllKVsBackup(stage *kk_stage.Stage, backupFileName string) (cmdStr string, err error) {

	//getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), "")
	//if err != nil || getResponse.Kvs == nil {
	//	slog.Error("failed to get kv", kk_stage.NewLog(stage).Error(err).Args()...)
	//	return -1, err, nil
	//}
	return cmdStr, nil
}

// todo cmk
func AllKVsRestore(stage *kk_stage.Stage, backupFileName string) (cmdStr string, err error) {
	//
	//getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), "")
	//if err != nil || getResponse.Kvs == nil {
	//	slog.Error("failed to get kv", kk_stage.NewLog(stage).Error(err).Args()...)
	//	return -1, err, nil
	//}
	return cmdStr, nil
}
