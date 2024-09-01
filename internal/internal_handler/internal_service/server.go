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

func (SerServer) RegisterService(stage *kk_stage.Stage, registration *kk_etcd_models.ServiceRegistration) error {
	switch registration.CheckConfig.Type {
	case kk_etcd_models.CheckTypeHttp:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = "http://" + registration.Addr + kk_etcd_models.HealthCheckPath
		}
	case kk_etcd_models.CheckTypeGrpc:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = registration.Addr
		}
	}
	err := registration.Check()
	if err != nil {
		return err
	}
	return toolServer.registerServer(stage, registration)
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func (SerServer) ServerList(client *clientv3.Client, serverType kk_etcd_models.ServiceType, serverName string) (serverList *kk_etcd_models.PBListServer, err error) {
	endpointMap, err := toolServer.serverList(client, serverType, serverName)
	if err != nil {
		return nil, err
	}
	//ListServer:{Key2EndpointMap:{key:"kk_service_http/ss/go_user/128.2.2.3:8484"  value:{Addr:"128.2.2.3:8484"}}}
	var pBListServer kk_etcd_models.PBListServer
	for key, endpoint := range endpointMap {
		var info checkInfo
		err := info.Convert(endpoint)
		if err != nil {
			pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
				ServiceName: key,
				ServiceAddr: endpoint.Addr,
				Status:      kk_etcd_models.PBServer_UnKnown,
				LastCheck:   timestamppb.New(kk_time.DefaultTime),
				Msg:         fmt.Sprintf("could not found server in check hub Err: %s", err.Error()),
			})
		} else {
			pBListServer.ListServer = append(pBListServer.ListServer, &kk_etcd_models.PBServer{
				ServiceName: key,
				ServiceAddr: endpoint.Addr,
				Status:      info.Status,
				LastCheck:   timestamppb.New(info.LastCheck),
				Msg:         info.Msg,
			})
		}
	}
	sort.Slice(pBListServer.ListServer, func(i, j int) bool {
		return pBListServer.ListServer[i].ServiceName < pBListServer.ListServer[j].ServiceName
	})
	return &pBListServer, nil
}

func (SerServer) DeregisterServer(stage *kk_stage.Stage, param *kk_etcd_models.DeregisterServerParam) error {
	key := param.GetServerType() + "/" + param.GetServerName()
	endpointManager, err := endpoints.NewManager(global_model.GetClient(stage), param.GetServerType())
	if err != nil {
		return err
	}

	err = endpointManager.DeleteEndpoint(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}
