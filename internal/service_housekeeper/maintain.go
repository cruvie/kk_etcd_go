package service_housekeeper

import (
	"context"
	"log/slog"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdMaintain https://etcd.io/docs/v3.5/op-guide/maintenance/#space-quota
func NewEtcdMaintain() *kk_server.KKRunServer {
	tick := time.NewTicker(1 * time.Hour)
	run := func() {
		maintainEtcd()
		for range tick.C {
			maintainEtcd()
		}
	}
	done := func(quitCh <-chan struct{}) {
		<-quitCh
		tick.Stop()
	}
	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}

func maintainEtcd() {
	for _, endpoint := range config.Config.Etcd.Endpoints {
		// get current revision
		currentRevision, err := getCurrentRevision(endpoint)
		if err != nil {
			slog.Error("Failed to get current revision", kk_stage.NewLog(nil).Error(err).Args()...)
			return
		}

		slog.Info("Current revision", slog.Any("revision", currentRevision))

		// compact
		err = executeCompact(currentRevision)
		if err != nil {
			slog.Error("Failed to execute compact", kk_stage.NewLog(nil).Error(err).Args()...)
			return
		}
		slog.Info("Compact executed successfully.")

		//  Defragment
		err = executeDefragment(endpoint)
		if err != nil {
			slog.Error("Failed to execute defragment", kk_stage.NewLog(nil).Error(err).Args()...)
			return
		}
		slog.Info("Defrag executed successfully.")

		// disarm alarm
		err = disarmAlarm()
		if err != nil {
			slog.Error("Failed to disarm alarm", kk_stage.NewLog(nil).Error(err).Args()...)
			return
		}
		slog.Info("Alarm disarmed.")
	}
}

func getCurrentRevision(endpoint string) (int64, error) {
	resp, err := global_model.GetClient(internal_client.GlobalStage).Status(context.Background(), endpoint)
	if err != nil {
		return 0, err
	}
	return resp.Header.Revision, nil
}

func executeCompact(revision int64) error {
	_, err := global_model.GetClient(internal_client.GlobalStage).Compact(context.Background(), revision)
	if utils.ErrorIs(err, rpctypes.ErrCompacted) {
		return nil
	}
	return err
}

func executeDefragment(endpoint string) error {
	_, err := global_model.GetClient(internal_client.GlobalStage).Defragment(context.Background(), endpoint)
	return err
}

func disarmAlarm() error {
	cli := global_model.GetClient(internal_client.GlobalStage)
	list, err := cli.AlarmList(context.Background())
	if err != nil {
		return err
	}

	for _, alarm := range list.Alarms {
		a, err := cli.AlarmDisarm(context.Background(), (*clientv3.AlarmMember)(alarm))
		if err != nil {
			return err
		}
		slog.Info("Disarmed alarm", kk_stage.NewLog(nil).Any("alarm", a).Args()...)
	}

	return nil
}
