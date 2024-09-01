package service

import (
	"context"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"
)

type SerBackup struct{}

var serBackup SerBackup

// Snapshot get snapshot, similar to `etcdctl snapshot save snapshot.db`
func (SerBackup) Snapshot(stage *kk_stage.Stage) (error, *kk_etcd_models.SnapshotResponse) {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	//etcdctl --endpoints=127.0.0.1:2379 --user kk_etcd:kk_etcd snapshot save /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshot.db
	backupData, err := global_model.GetClient(stage).Snapshot(context.Background())
	defer func(backupData io.ReadCloser) {
		if backupData != nil {
			err := backupData.Close()
			if err != nil {
				slog.Error("Failed to close backup data", newLog.Error(err).Args()...)
			}
		}
	}(backupData)
	if err != nil {
		return err, nil
	}

	snapshotBytes, err := io.ReadAll(backupData)
	if err != nil {
		return err, nil
	}
	timeStr := time.Now().Format(time.DateTime)
	fileName := "etcd_" + timeStr + ".snapshot"

	return nil, &kk_etcd_models.SnapshotResponse{
		Name: fileName,
		File: snapshotBytes,
	}
}

// SnapshotRestore todo migrate to etcd cluster manager
func (SerBackup) SnapshotRestore() (cmdStr string, err error) {

	//https://etcd.io/docs/v3.5/op-guide/recovery/
	// 执行etcdctl命令来恢复etcd数据
	//cmd := exec.Command("etcdctl", "snapshot", "restore", "/path/to/snapshot.db", "--name", "my-etcd")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Println("Failed to restore etcd data:", err)
	//	return
	//}
	cmdStr = `
			etcdutl snapshot restore /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshotFile.snapshot \
		  --data-dir  /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/out \
		  --name node1 \
		  --initial-cluster node1=http://127.0.0.1:2380 \
		  --initial-advertise-peer-urls http://127.0.0.1:2380
`
	return cmdStr, nil
}
func (SerBackup) SnapshotInfo(stage *kk_stage.Stage, param *kk_etcd_models.SnapshotInfoParam) (info string, err error) {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	tempFile, err := os.CreateTemp("", "temp")
	if err != nil {
		return "", err
	}
	defer func() {
		err = os.Remove(tempFile.Name())
		if err != nil {
			slog.Error("delete temp file failed", newLog.Error(err).Args()...)
		}
	}()
	if _, err := tempFile.Write(param.GetFile()); err != nil {
		return "", err
	}
	if err := tempFile.Close(); err != nil {
		return "", err
	}

	//etcdutl --write-out=table snapshot status etcd_backup.snapshot
	cmd := exec.Command("etcdutl", "--write-out", "table", "snapshot", "status", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
func (SerBackup) AllKVsBackup(stage *kk_stage.Stage) (error, *kk_etcd_models.AllKVsBackupResponse) {
	list := &kk_etcd_models.PBListKV{}
	getResponse, err := global_model.GetClient(stage).Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		return err, nil
	}
	for _, kv := range getResponse.Kvs {
		pbKV := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//todo only support backup normal kv and config
		if strings.HasPrefix(pbKV.Key, kk_etcd_models.ServiceHttp.String()) ||
			strings.HasPrefix(pbKV.Key, kk_etcd_models.ServiceGrpc.String()) {
			continue
		}

		list.ListKV = append(list.ListKV, pbKV)
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return err, nil
	}
	timeStr := time.Now().Format(time.DateTime)
	fileName := "etcd_all_kv_" + timeStr + ".json"
	return nil, &kk_etcd_models.AllKVsBackupResponse{
		Name: fileName,
		File: marshal,
	}
}
func (SerBackup) AllKVsRestore(stage *kk_stage.Stage, param *kk_etcd_models.AllKVsRestoreParam) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	list := &kk_etcd_models.PBListKV{}
	err := json.Unmarshal(param.GetFile(), list)
	if err != nil {
		slog.Error("failed to Unmarshal kv", newLog.Error(err).Args()...)
		return err
	}
	for _, pbkv := range list.ListKV {
		err := serKV.KVPut(stage, pbkv.GetKey(), pbkv.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}
