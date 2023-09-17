package service

import (
	"context"
	"errors"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
	"sort"
)

func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	if registration.ServerType != key_prefix.ServiceHttp && registration.ServerType != key_prefix.ServiceGrpc {
		slog.Error("server type is invalid")
		return errors.New("server type is invalid")
	}
	if registration.ServerName == "" {
		slog.Error("server name cannot be empty")
		return errors.New("server name cannot be empty")
	}
	if registration.Address == "" {
		slog.Error("server address cannot be empty")
		return errors.New("server address cannot be empty")
	}
	if registration.Check == nil {
		slog.Error("server Check cannot be empty")
		return errors.New("server Check cannot be empty")
	}
	switch registration.Check.Type {
	case kk_etcd_models.CheckTypeHttp:
		if registration.Check.HTTP == "" {
			registration.Check.HTTP = "http://" + registration.Address + "/KKHealthCheck"
		}
	case kk_etcd_models.CheckTypeGrpc:
		if registration.Check.GRPC == "" {
			registration.Check.GRPC = registration.Address
		}
	default:
		slog.Error("server Check Type is invalid")
		return errors.New("server Check Type is invalid")
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
	return registerServer(registration)
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (res int, serverList *kk_etcd_models.PBListServer, err error) {
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
