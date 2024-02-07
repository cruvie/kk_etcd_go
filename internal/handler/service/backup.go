package service

import (
	"context"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Snapshot get snapshot, similar to `etcdctl snapshot save snapshot.db`
func Snapshot(stage *kk_stage.Stage) (pbFile *kk_pb_type.PBFile, err error) {
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
		return nil, err
	}

	snapshotBytes, err := io.ReadAll(backupData)
	if err != nil {
		slog.Error("Failed to read backup data", kk_stage.NewLog(stage).Error(err).Args()...)
		return nil, err
	}
	timeStr := time.Now().Format(kk_time.DateTime)
	fileName := "etcd_" + timeStr + ".snapshot"

	return &kk_pb_type.PBFile{
		Name:  fileName,
		Bytes: snapshotBytes,
	}, nil
}

// SnapshotRestore todo migrate to etcd cluster manager
func SnapshotRestore(stage *kk_stage.Stage) (cmdStr string, err error) {
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

func SnapshotInfo(stage *kk_stage.Stage, pbFile *kk_pb_type.PBFile) (info string, err error) {
	tempFile, err := os.CreateTemp("", "temp")
	if err != nil {
		slog.Error("create temp file failed", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}
	defer func() {
		err = os.Remove(tempFile.Name())
		if err != nil {
			slog.Error("delete temp file failed", kk_stage.NewLog(stage).Error(err).Args()...)
		}
	}()
	if _, err := tempFile.Write(pbFile.Bytes); err != nil {
		slog.Error("write temp file failed", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}
	if err := tempFile.Close(); err != nil {
		slog.Error("close temp file failed", kk_stage.NewLog(stage).Error(err).Args()...)
		return "", err
	}

	//etcdutl --write-out=table snapshot status etcd_backup.snapshot
	cmd := exec.Command("etcdutl", "--write-out", "table", "snapshot", "status", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		s := cmd.String()
		slog.Error("failed to get snapshot info", kk_stage.NewLog(stage).Error(err).String("cmd", s).Args()...)
		return "", err
	}

	return string(output), nil
}

func AllKVsBackup(stage *kk_stage.Stage) (pbFile *kk_pb_type.PBFile, err error) {
	list := &kk_etcd_models.PBListKV{}
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil || getResponse.Kvs == nil {
		slog.Error("failed to get kv", kk_stage.NewLog(stage).Error(err).Args()...)
		return nil, err
	}
	for _, kv := range getResponse.Kvs {
		pbKV := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//todo only support backup normal kv and config
		if strings.HasPrefix(pbKV.Key, kk_etcd_const.User) ||
			strings.HasPrefix(pbKV.Key, kk_etcd_const.Jwt) ||
			strings.HasPrefix(pbKV.Key, kk_etcd_const.ServiceHttp) ||
			strings.HasPrefix(pbKV.Key, kk_etcd_const.ServiceGrpc) {
			continue
		}

		list.ListKV = append(list.ListKV, pbKV)
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().Format(kk_time.DateTime)
	fileName := "etcd_all_kv_" + timeStr + ".json"
	return &kk_pb_type.PBFile{
		Name:  fileName,
		Bytes: marshal,
	}, nil
}

func AllKVsRestore(stage *kk_stage.Stage, pbFile *kk_pb_type.PBFile) (err error) {
	list := &kk_etcd_models.PBListKV{}
	err = json.Unmarshal(pbFile.GetBytes(), list)
	if err != nil {
		slog.Error("failed to Unmarshal kv", kk_stage.NewLog(stage).Error(err).Args()...)
		return err
	}
	for _, pbkv := range list.ListKV {
		err := KVPut(stage, pbkv.GetKey(), pbkv.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}
