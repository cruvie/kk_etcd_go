package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/server/serverList"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"log/slog"
)

var serServerHub server_hub.SerServer

// RegisterService register service to etcd
func RegisterService(registration *kk_etcd_models.ServerRegistration) error {
	err := serServerHub.RegisterService(internal_client.GlobalStage, registration)
	return err
}

// ServerList
// serverName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(param *serverList.ServerList_Input) (*kk_etcd_models.PBListServer, error) {
	return serServerHub.ServerList(GetClient(), kk_etcd_models.ServerType(param.GetServerType()))
}

// WatchServerList watch server list change
func WatchServerList(ctx context.Context, serverType kk_etcd_models.ServerType, serverName string, serverListChan chan<- *kk_etcd_models.PBListServer) (err error) {
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: internal_client.GlobalStage.TraceId})
	etcdManager, err := serverType.NewEndpointManager(GetClient())

	if err != nil {
		slog.Error("failed to new endpoints.Manager", newLog.Any("serverName", serverName).Error(err).Args()...)
		return err
	}
	channel, err := etcdManager.NewWatchChannel(ctx)
	if err != nil {
		slog.Error("failed to new watch channel", newLog.Any("serverName", serverName).Error(err).Args()...)
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
						EndpointKey:  update.Key,
						EndpointAddr: update.Endpoint.Addr,
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
