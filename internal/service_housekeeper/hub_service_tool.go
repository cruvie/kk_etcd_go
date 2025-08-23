package service_housekeeper

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/kubernetes"
	"log/slog"
)

func (*SerService) registerService(registrations []*kk_etcd_models.PBServiceRegistration) error {
	for _, registration := range registrations {
		val, err := registration.Marshal()
		if err != nil {
			slog.Error("failed to get Marshal", kk_stage.NewLog(nil).Error(err).Any("service", registration).Args()...)
			return err
		}
		_, err = kc.Put(context.Background(), registration.UniqueKey(), val)
		if err != nil {
			slog.Error("failed to put to etcd", kk_stage.NewLog(nil).Error(err).Any("service", registration).Error(err).Args()...)
			return err
		}
	}
	return nil
}

func (*SerService) deRegisterService(stage *kk_stage.Stage, registration *kk_etcd_models.PBServiceRegistration) error {
	newLog := kk_stage.NewLog(stage)
	_, err := kc.Delete(context.Background(), registration.UniqueKey())
	if err != nil {
		slog.Error("failed to delete from etcd", newLog.Error(err).Args()...)
		return err
	}
	return nil
}

// serviceList nil serviceType means all serviceType
func (*SerService) serviceList(keyPrefix string) (serviceList []*kk_etcd_models.PBServiceRegistration, err error) {
	resp, err := kc.List(context.Background(), keyPrefix, kubernetes.ListOptions{
		Revision: 0,
		Limit:    0,
		Continue: "",
	})
	if err != nil {
		return nil, err
	}
	serviceList = make([]*kk_etcd_models.PBServiceRegistration, 0)
	for _, kv := range resp.Kvs {
		if kv.Value != nil {
			registration := &kk_etcd_models.PBServiceRegistration{}
			err = registration.UnMarshal(kv.Value)
			if err != nil {
				return nil, err
			}
			serviceList = append(serviceList, registration)
		}
	}
	return serviceList, nil
}
