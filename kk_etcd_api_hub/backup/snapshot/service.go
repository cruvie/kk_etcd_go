package snapshot

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"
	"io"
	"log/slog"
	"time"
)

func (x *api) service() (*api_def.Snapshot_Output, error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: x.stage.TraceId})
	//etcdctl --endpoints=127.0.0.1:2379 --user kk_etcd:kk_etcd snapshot save /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshot.db
	backupData, err := global_model.GetClient(x.stage).Snapshot(context.Background())
	defer func(backupData io.ReadCloser) {
		if backupData != nil {
			err := backupData.Close()
			if err != nil {
				slog.Error("Failed to close backup data", newLog.Error(err).Args()...)
			}
		}
	}(backupData)
	if err != nil {
		return nil, err
	}

	snapshotBytes, err := io.ReadAll(backupData)
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().Format(time.DateTime)
	fileName := "etcd_" + timeStr + ".snapshot"

	return &api_def.Snapshot_Output{
		Name: fileName,
		File: snapshotBytes,
	}, nil
}
