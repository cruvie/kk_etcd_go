package server_hub

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"sort"
)

type SerServer struct{}

var serServerInstance = &SerServer{}

func (x *SerServer) RegisterService(registration *kk_etcd_models.PBServerRegistration) error {
	switch registration.CheckConfig.Type {
	case kk_etcd_models.PBServerType_Http:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = "http://" + registration.ServerAddr + kk_etcd_models.HealthCheckPath
		}
	case kk_etcd_models.PBServerType_Grpc:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = registration.ServerAddr
		}
	}
	err := registration.Check()
	if err != nil {
		return err
	}
	return x.registerServer([]*kk_etcd_models.PBServerRegistration{registration})
}

func (x *SerServer) DeregisterServer(stage *kk_stage.Stage, registration *kk_etcd_models.PBServerRegistration) error {
	return x.deRegisterServer(stage, registration)
}

func (x *SerServer) ServerList(serverType kk_etcd_models.PBServerType, serverName string) (serverList *kk_etcd_models.PBListServer, err error) {
	preFix := kk_etcd_models.ServerRegistrationKey + serverType.String()
	if serverName != "" {
		preFix += "/" + serverName
	}

	list, err := x.serverList(preFix)
	if err != nil {
		return nil, err
	}

	aliveStatuMap := hub.aliveHub.Data()
	deadStatuMap := hub.deadHub.Data()

	getStatus := func(item *kk_etcd_models.PBServerRegistration) kk_etcd_models.PBServer_ServerStatus {
		serverStatus, ok := aliveStatuMap[item.ServerName]
		switch {
		case ok:
			{
				index := slices.IndexFunc(serverStatus.Slice(), func(c *kk_etcd_models.PBServerRegistration) bool {
					return c.Key() == item.Key()
				})
				if index != -1 {
					return kk_etcd_models.PBServer_Running
				}
			}
		default:
			{
				serverStatus, ok = deadStatuMap[item.ServerName]
				if !ok {
					return kk_etcd_models.PBServer_UnKnown
				} else {
					index := slices.IndexFunc(serverStatus.Slice(), func(c *kk_etcd_models.PBServerRegistration) bool {
						return c.Key() == item.Key()
					})
					if index != -1 {
						return kk_etcd_models.PBServer_Stop
					}
				}
			}
		}
		return kk_etcd_models.PBServer_UnKnown
	}
	var pBListServer kk_etcd_models.PBListServer
	for _, item := range list {
		server := &kk_etcd_models.PBServer{
			ServerRegistration: item,
			Status:             getStatus(item),
		}
		pBListServer.ListServer = append(pBListServer.ListServer, server)
	}

	sort.Slice(pBListServer.ListServer, func(i, j int) bool {
		return pBListServer.ListServer[i].ServerRegistration.Key() < pBListServer.ListServer[j].ServerRegistration.Key()
	})
	return &pBListServer, nil
}

func (x *SerServer) GetConns(connType kk_etcd_models.PBServerType, serverName string) (string, error) {
	server, err := getOneAliveServer(connType, serverName)
	if err != nil {
		return "", err
	}
	if server != nil {
		return server.ServerAddr, nil
	}
	return "", kk_etcd_error.ErrNoAvailableConn
}
