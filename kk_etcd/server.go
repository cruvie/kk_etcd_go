package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/internal_handler/internal_service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

var serInternalServer internal_service.SerServer

// RegisterService register service to etcd
func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	err := serInternalServer.RegisterService(internal_client.GlobalStage, registration)
	return err
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(param *kk_etcd_models.ServerListParam) (*kk_etcd_models.PBListServer, error) {
	return serInternalServer.ServerList(global_model.GetClient(internal_client.GlobalStage), kk_etcd_models.ServiceType(param.GetServerType()), param.GetServerName())
}

// WatchServerList watch server list change
func WatchServerList(ctx context.Context, serverType kk_etcd_models.ServiceType, serviceName string, serverListChan chan<- *kk_etcd_models.PBListServer) (err error) {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: internal_client.GlobalStage.TraceId})
	etcdManager, err := endpoints.NewManager(global_model.GetClient(internal_client.GlobalStage), serverType.String()+serviceName)

	if err != nil {
		slog.Error("failed to new endpoints.Manager", newLog.Any("serviceName", serviceName).Error(err).Args()...)
		return err
	}
	channel, err := etcdManager.NewWatchChannel(ctx)
	if err != nil {
		slog.Error("failed to new watch channel", newLog.Any("serviceName", serviceName).Error(err).Args()...)
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case updates := <-channel:
				var pBListServer kk_etcd_models.PBListServer
				for _, update := range updates {
					pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
						ServiceName: update.Key,
						ServiceAddr: update.Endpoint.Addr,
					})
				}
				if len(pBListServer.ListServer) > 0 {
					serverListChan <- &pBListServer
				}
			}
		}
	}()
	return nil
}
