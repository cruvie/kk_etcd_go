package internal_service

import (
	"context"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sort"
)

type SerServer struct{}

var serInternalServer SerServer

func (SerServer) RegisterService(stage *kk_stage.Stage, registration *kk_etcd_models.ServerRegistration) error {
	switch registration.CheckConfig.Type {
	case kk_etcd_models.Http:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = "http://" + registration.ServerAddr + kk_etcd_models.HealthCheckPath
		}
	case kk_etcd_models.Grpc:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = registration.ServerAddr
		}
	}
	err := registration.Check()
	if err != nil {
		return err
	}
	return toolServer.registerServer(stage, registration)
}

// ServerList
// serverName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func (SerServer) ServerList(client *clientv3.Client, serverType kk_etcd_models.ServerType, serverName string) (serverList *kk_etcd_models.PBListServer, err error) {
	endpointMap, err := toolServer.serverList(client, serverType, serverName)
	if err != nil {
		return nil, err
	}
	var pBListServer kk_etcd_models.PBListServer
	serviceStatus, err := hub.services()
	if err != nil {
		return nil, err
	}
	for key, endpoint := range endpointMap {
		status, ok := serviceStatus[key]
		if !ok {
			pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
				EndpointManagerTarget: status.EndpointManagerTarget(),
				EndpointKey:           status.EndpointKey(),
				ServerName:            status.ServerName,
				ServerAddr:            endpoint.Addr,
				Status:                kk_etcd_models.PBServer_UnKnown,
				LastCheck:             timestamppb.New(kk_time.DefaultTime),
				Msg:                   fmt.Sprintf("could not found sevice %s in service hub \n may not be registered by kk_etcd", key),
			})
		} else {
			pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
				EndpointManagerTarget: status.EndpointManagerTarget(),
				EndpointKey:           status.EndpointKey(),
				ServerName:            status.ServerName,
				ServerAddr:            endpoint.Addr,
				Status:                status.Status,
				LastCheck:             timestamppb.New(status.LastCheck),
				Msg:                   status.Msg,
			})
		}
	}
	sort.Slice(pBListServer.ListServer, func(i, j int) bool {
		return pBListServer.ListServer[i].ServerName < pBListServer.ListServer[j].ServerName
	})
	return &pBListServer, nil
}

func (SerServer) DeregisterServer(stage *kk_stage.Stage, param *kk_etcd_models.DeregisterServerParam) error {
	endpointManager, err := endpoints.NewManager(global_model.GetClient(stage), param.GetServer().GetEndpointManagerTarget())
	if err != nil {
		return err
	}

	err = endpointManager.DeleteEndpoint(context.Background(), param.GetServer().GetEndpointKey())
	if err != nil {
		return err
	}
	return nil
}
