package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

// RegisterService register service to etcd
func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	err := service.RegisterService(stage, registration)
	return err
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (serverList *kk_etcd_models.PBListServer, err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	_, servers, err := service.ServerList(stage, serviceName)
	return servers, err
}

// WatchServerList watch server list change
func WatchServerList(ctx context.Context, serviceName string, serverListChan chan *kk_etcd_models.PBListServer) (err error) {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, serviceName)

	if err != nil {
		slog.Error("failed to new endpoints.Manager", kk_stage.NewLog(stage).Any("serviceName", serviceName).Error(err).Args()...)
		return err
	}
	channel, err := etcdManager.NewWatchChannel(ctx)
	if err != nil {
		slog.Error("failed to new watch channel", kk_stage.NewLog(stage).Any("serviceName", serviceName).Error(err).Args()...)
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
