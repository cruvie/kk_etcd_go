package server_hub

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/kubernetes"
	"log/slog"
)

func (*SerServer) registerServer(registrations []*kk_etcd_models.PBServerRegistration) error {
	for _, registration := range registrations {
		val, err := registration.Marshal()
		if err != nil {
			slog.Error("failed to get Marshal", kk_log.NewLog(nil).Error(err).Any("server", registration).Args()...)
			return err
		}
		_, err = kc.Put(context.Background(), registration.Key(), val)
		if err != nil {
			slog.Error("failed to put to etcd", kk_log.NewLog(nil).Error(err).Any("server", registration).Error(err).Args()...)
			return err
		}
	}
	return nil
}

func (*SerServer) deRegisterServer(stage *kk_stage.Stage, registration *kk_etcd_models.PBServerRegistration) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	_, err := kc.Delete(context.Background(), registration.Key())
	if err != nil {
		slog.Error("failed to delete from etcd", newLog.Error(err).Args()...)
		return err
	}
	return nil
}

// serverList nil serverType means all serverType
func (*SerServer) serverList(keyPrefix string) (serverList []*kk_etcd_models.PBServerRegistration, err error) {
	resp, err := kc.List(context.Background(), keyPrefix, kubernetes.ListOptions{
		Revision: 0,
		Limit:    0,
		Continue: "",
	})
	if err != nil {
		return nil, err
	}
	serverList = make([]*kk_etcd_models.PBServerRegistration, 0)
	for _, kv := range resp.Kvs {
		if kv.Value != nil {
			registration := &kk_etcd_models.PBServerRegistration{}
			err = registration.UnMarshal(kv.Value)
			if err != nil {
				return nil, err
			}
			serverList = append(serverList, registration)
		}
	}
	return serverList, nil
}
