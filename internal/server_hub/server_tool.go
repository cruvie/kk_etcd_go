package server_hub

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/util_kv"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

type serverTool struct{}

var toolServer serverTool

func (t *serverTool) serverList(client *clientv3.Client, serverType kk_etcd_models.ServerType) (endpoints.Key2EndpointMap, error) {
	etcdManager, err := serverType.NewEndpointManager(client)
	if err != nil {
		return nil, err
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		return nil, err
	}
	return endpointMap, nil
}
func (t *serverTool) services() (map[string]*serverStatus, error) {
	services := make(map[string]*serverStatus)
	err, v := util_kv.ListKV(internal_client.GlobalStage, kk_etcd_models.InternalServerStatus)
	if err != nil {
		return nil, err
	}
	for _, kv := range v.GetListKV() {
		var info serverStatus
		err := info.FromJson(kv.GetValue())
		if err != nil {
			return nil, err
		}
		services[info.kVKey()] = &info
	}

	return services, nil
}

func (t *serverTool) registerServer(stage *kk_stage.Stage, registration *kk_etcd_models.ServerRegistration) error {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})

	endpointManager, err := registration.ServerType.NewEndpointManager(global_model.GetClient(stage))
	if err != nil {
		msg := "failed to create etcd manager"
		slog.Error(msg, newLog.Error(err).String("key", registration.EndpointKey()).Args()...)
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
