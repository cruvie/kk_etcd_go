package internal_service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

type serverTool struct{}

var toolServer serverTool

func (t *serverTool) serverList(client *clientv3.Client, serverType kk_etcd_models.ServerType, serverName string) (endpoints.Key2EndpointMap, error) {
	etcdManager, err := endpoints.NewManager(client, serverType.EndpointManagerTarget(serverName))
	if err != nil {
		return nil, err
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		return nil, err
	}
	return endpointMap, nil
}

func (t *serverTool) registerServer(stage *kk_stage.Stage, registration *kk_etcd_models.ServerRegistration) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})

	endpointManager, err := endpoints.NewManager(global_model.GetClient(stage), registration.EndpointManagerTarget())
	if err != nil {
		msg := "failed to create etcd manager"
		slog.Error(msg, newLog.Error(err).String("target", registration.EndpointManagerTarget()).Args()...)
		return err
	}
	info := newServerStatus(registration)
	endpoint := endpoints.Endpoint{
		Addr:     registration.ServerAddr,
		Metadata: info,
	}
	//add endpoint to etcd
	err = endpointManager.AddEndpoint(
		context.Background(),
		info.EndpointKey(),
		endpoint)
	if err != nil {
		msg := "failed to add endpoint to etcd"
		slog.Error(msg, newLog.Error(err).Args()...)
		return err
	}
	return nil
}
