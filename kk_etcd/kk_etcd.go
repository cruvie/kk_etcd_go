package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"gopkg.in/yaml.v3"
	"log/slog"
)

func InitEtcd(endpoints []string, userName string, password string) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	return service.InitEtcd(stage, endpoints, userName, password)
}

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), kk_etcd_const.Config+configKey)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("configKey", configKey)
		slog.Error("failed to get config", logBody.GetLogArgs()...)
		return err
	}
	if getResponse.Kvs == nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		slog.Warn("failed to get kv, getResponse.Kvs is nil", logBody.GetLogArgs()...)
		return nil
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("config", string(getResponse.Kvs[0].Value))
		slog.Error("failed Unmarshal config", logBody.GetLogArgs()...)
		return err
	}
	return nil
}

func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	err := service.RegisterService(stage, registration)
	return err
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (serverList *kk_etcd_models.PBListServer, err error) {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	_, servers, err := service.ServerList(stage, serviceName)
	return servers, err
}

// todo cmk watch server list change update to cache
// 使用errgroup启动这个协程方便监听错误
func WatchServerList(ctx context.Context, serviceName string, serverListChan chan *kk_etcd_models.PBListServer) (err error) {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, serviceName)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("serviceName", serviceName)
		slog.Error("failed to new endpoints.Manager", logBody.GetLogArgs()...)
		return err
	}
	channel, err := etcdManager.NewWatchChannel(ctx)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("serviceName", serviceName)
		slog.Error("failed to new watch channel", logBody.GetLogArgs()...)
		return err
	}
	var pBListServer kk_etcd_models.PBListServer
	for {
		select {
		case <-ctx.Done():
			return nil
		case updates := <-channel:
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
}

//监听指定服务的服务列表变化，更新grpc客户端 ss_grpc存储一个client列表，ss_go使用算法每次请求随机选择一个client
