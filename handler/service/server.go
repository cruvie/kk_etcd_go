package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
	"strings"
)

// RegisterHttpService
// addr: 192.123.32.11:8080
// serverName:
// ttl:lease time, default 15 seconds, service will be deleted from etcd after ttl seconds
func RegisterHttpService(ctx context.Context, addr, serverName string, ttl ...int64) {
	registerServer(ctx, key_prefix.ServiceHttp, addr, serverName, ttl...)
}

// RegisterGrpcService
// addr: 192.123.32.11:8080
// serverName:
// ttl:lease time, default 15 seconds, service will be deleted from etcd after ttl seconds
func RegisterGrpcService(ctx context.Context, addr, serverName string, ttl ...int64) {
	registerServer(ctx, key_prefix.ServiceGrpc, addr, serverName, ttl...)
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (res int, serverList *models.PBListServer, err error) {
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, serviceName)
	if err != nil {
		slog.Error("failed to create etcd manager", "err", err)
		return -1, nil, err
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		slog.Error("failed to list endpoints", "err", err)
		return -1, nil, err
	}
	//ListServer:{Key2EndpointMap:{key:"kk_service_http/haha/128.2.2.3:8484"  value:{Addr:"128.2.2.3:8484"}}}
	var pBListServer models.PBListServer
	for key, endpoint := range endpointMap {
		split := strings.Split(key, "/")
		var name string
		if len(split) >= 2 {
			name = split[1]
		} else {
			name = key
		}
		pBListServer.ListServer = append(pBListServer.ListServer, &models.PBServer{
			ServiceName: name,
			ServiceAddr: endpoint.Addr,
		})
	}
	return 1, &pBListServer, nil
}
