package snapshot

import (
	"context"
	"io"
	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *Api) Service(stage *kk_stage.Stage) (*api_def.Snapshot_Output, error) {
	span := stage.StartTrace("service")
	defer span.End()

	newLog := kk_stage.NewLog(stage)
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
