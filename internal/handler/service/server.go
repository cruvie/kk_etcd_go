package service

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
	"sort"
)

func RegisterService(stage *kk_stage.Stage, registration *kk_etcd_models.ServiceRegistration) error {
	if registration.ServerType != kk_etcd_const.ServiceHttp && registration.ServerType != kk_etcd_const.ServiceGrpc {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "server type is invalid"
		slog.Error(msg, logBody.GetLogArgs()...)
		return errors.New(msg)
	}
	if registration.ServerName == "" {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "server name cannot be empty"
		slog.Error(msg, logBody.GetLogArgs()...)
		return errors.New(msg)
	}
	if registration.Address == "" {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "server address cannot be empty"
		slog.Error(msg, logBody.GetLogArgs()...)
		return errors.New(msg)
	}
	if registration.Check == nil {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "server Check cannot be empty"
		slog.Error(msg, logBody.GetLogArgs()...)
		return errors.New(msg)
	}
	switch registration.Check.Type {
	case kk_etcd_models.CheckTypeHttp:
		if registration.Check.HTTP == "" {
			registration.Check.HTTP = "http://" + registration.Address + kk_etcd_models.HealthCheckPath
		}
	case kk_etcd_models.CheckTypeGrpc:
		if registration.Check.GRPC == "" {
			registration.Check.GRPC = registration.Address
		}
	default:
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "server Check Type is invalid"
		slog.Error(msg, logBody.GetLogArgs()...)
		return errors.New(msg)
	}
	if registration.Check.TTL == 0 {
		registration.Check.TTL = 30
	}
	if registration.Check.Timeout == 0 {
		registration.Check.Timeout = registration.Check.TTL / 3
	}
	if registration.Check.Interval == 0 {
		registration.Check.Interval = registration.Check.TTL / 3
	}
	return toolServer.registerServer(stage, registration)
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(stage *kk_stage.Stage, serviceName string) (res int, serverList *kk_etcd_models.PBListServer, err error) {
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, serviceName)
	if err != nil {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "failed to create etcd manager"
		slog.Error(msg, logBody.GetLogArgs()...)
		return -1, nil, err
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId)
		msg := "failed to list endpoints"
		slog.Error(msg, logBody.GetLogArgs()...)
		return -1, nil, err
	}
	//ListServer:{Key2EndpointMap:{key:"kk_service_http/ss/go_user/128.2.2.3:8484"  value:{Addr:"128.2.2.3:8484"}}}
	var pBListServer kk_etcd_models.PBListServer
	for key, endpoint := range endpointMap {
		pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
			ServiceName: key,
			ServiceAddr: endpoint.Addr,
		})
	}
	sort.Slice(pBListServer.ListServer, func(i, j int) bool {
		return pBListServer.ListServer[i].ServiceName < pBListServer.ListServer[j].ServiceName
	})
	return 1, &pBListServer, nil
}
