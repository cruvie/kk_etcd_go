package service_housekeeper

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"sort"
)

type SerService struct{}

var serServiceInstance = &SerService{}

func (x *SerService) RegisterService(registration *kk_etcd_models.PBServiceRegistration) error {
	switch registration.CheckConfig.Type {
	case kk_etcd_models.PBServiceType_Http:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = "http://" + registration.ServiceAddr + kk_etcd_models.HealthCheckPath
		}
	case kk_etcd_models.PBServiceType_Grpc:
		if registration.CheckConfig.Addr == "" {
			registration.CheckConfig.Addr = registration.ServiceAddr
		}
	}
	err := registration.Check()
	if err != nil {
		return err
	}
	return x.registerService([]*kk_etcd_models.PBServiceRegistration{registration})
}

func (x *SerService) DeregisterService(stage *kk_stage.Stage, registration *kk_etcd_models.PBServiceRegistration) error {
	return x.deRegisterService(stage, registration)
}

func (x *SerService) ServiceList(serviceType kk_etcd_models.PBServiceType, serviceName string) (serviceList *kk_etcd_models.PBListService, err error) {
	preFix := kk_etcd_models.ServiceRegistrationKey + serviceType.String()
	if serviceName != "" {
		preFix += "/" + serviceName
	}

	list, err := x.serviceList(preFix)
	if err != nil {
		return nil, err
	}

	aliveStatuMap := hub.aliveHub.Data() //todo 死锁！！
	deadStatuMap := hub.deadHub.Data()

	getStatus := func(item *kk_etcd_models.PBServiceRegistration) kk_etcd_models.PBService_ServiceStatus {
		serviceStatus, ok := aliveStatuMap[item.HubKey()]
		switch {
		case ok:
			{
				index := slices.IndexFunc(serviceStatus.Slice(), func(c *kk_etcd_models.PBServiceRegistration) bool {
					return c.Equal(item)
				})
				if index != -1 {
					return kk_etcd_models.PBService_Running
				}
			}
		default:
			{
				serviceStatus, ok = deadStatuMap[item.HubKey()]
				if !ok {
					return kk_etcd_models.PBService_UnKnown
				} else {
					index := slices.IndexFunc(serviceStatus.Slice(), func(c *kk_etcd_models.PBServiceRegistration) bool {
						return c.Equal(item)
					})
					if index != -1 {
						return kk_etcd_models.PBService_Stop
					}
				}
			}
		}
		return kk_etcd_models.PBService_UnKnown
	}
	var pBListService kk_etcd_models.PBListService
	for _, item := range list {
		service := &kk_etcd_models.PBService{
			ServiceRegistration: item,
			Status:              getStatus(item),
		}
		pBListService.ListService = append(pBListService.ListService, service)
	}

	sort.Slice(pBListService.ListService, func(i, j int) bool {
		return pBListService.ListService[i].ServiceRegistration.UniqueKey() < pBListService.ListService[j].ServiceRegistration.UniqueKey()
	})
	return &pBListService, nil
}

func (x *SerService) GetServiceAddr(connType kk_etcd_models.PBServiceType, serviceName string) (string, error) {
	service, err := getOneAliveService(connType, serviceName)
	if err != nil {
		return "", err
	}
	if service != nil {
		return service.ServiceAddr, nil
	}
	return "", kk_etcd_error.ErrNoAvailableConn
}
