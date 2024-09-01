package internal_service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

type serverTool struct{}

var toolServer serverTool

func (t *serverTool) serverList(client *clientv3.Client, serverType kk_etcd_models.ServiceType, serverName string) (endpoints.Key2EndpointMap, error) {
	etcdManager, err := endpoints.NewManager(client, serverType.String())
	if err != nil {
		return nil, err
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		return nil, err
	}
	return endpointMap, nil
}

func (t *serverTool) registerServer(stage *kk_stage.Stage, registration *kk_etcd_models.ServiceRegistration) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})

	endpointManager, err := endpoints.NewManager(global_model.GetClient(stage), registration.ServerType.String())
	if err != nil {
		msg := "failed to create etcd manager"
		slog.Error(msg, newLog.Error(err).String("key", registration.EndpointKey()).Args()...)
		return err
	}
	info := checkInfo{
		Key:         registration.EndpointKey(),
		CheckConfig: registration.CheckConfig,
		ServiceType: registration.ServerType,
		Status:      kk_etcd_models.PBServer_Init,
		LastCheck:   kk_time.DefaultTime,
		Msg:         "init checker",
		Metadata:    registration.Metadata,
	}
	endpoint := endpoints.Endpoint{
		Addr:     registration.Addr,
		Metadata: info,
	}
	//add endpoint to etcd
	err = endpointManager.AddEndpoint(
		context.Background(),
		registration.EndpointKey(),
		endpoint)
	if err != nil {
		msg := "failed to add endpoint to etcd"
		slog.Error(msg, newLog.Error(err).Args()...)
		return err
	}
	return nil
}
