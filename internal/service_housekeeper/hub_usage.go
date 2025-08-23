package service_housekeeper

import (
	"log/slog"
	"sync"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var hub *serverHub

// todo 服务在删除时正在检查会有并发问题，检查后可能又被重新注册
var hubLock sync.RWMutex //nolint

type serverHub struct {
	aliveHub *hubT
	deadHub  *hubT
}

func InitServiceHub() {
	hub = &serverHub{
		aliveHub: newHubT(),
		deadHub:  newHubT(),
	}
	go hub.watch()
	time.Sleep(1 * time.Second)

	list, err := serServiceInstance.serviceList(kk_etcd_models.ServiceRegistrationKey)
	if err != nil {
		panic(err)
	}
	for _, registration := range list {
		hub.putToDeadHub(registration)
	}
}

func (x *serverHub) delFromDeadHub(item *kk_etcd_models.PBServiceRegistration) {
	delFromHub(x.deadHub, item)
}

func (x *serverHub) putToDeadHub(item *kk_etcd_models.PBServiceRegistration) {
	alreadyExists := putToHub(x.deadHub, item)
	if !alreadyExists {
		startNewCheck(item)
	}
}

func (x *serverHub) delFromAliveHub(item *kk_etcd_models.PBServiceRegistration) {
	delFromHub(x.aliveHub, item)
	putAliveHubToEtcd(x.aliveHub)
}

func (x *serverHub) putToAliveHub(item *kk_etcd_models.PBServiceRegistration) {
	alreadyExists := putToHub(x.aliveHub, item)
	if !alreadyExists {
		//after put into aliveHub, put all alive conn to etcd for App outside to get conn
		putAliveHubToEtcd(x.aliveHub)
	}
}

// need this for SDK use, which cannot access memory data of running kk_etcd
func (x *serverHub) watch() {
	watchGrpcCh := kc.Watch(kc.Ctx(), kk_etcd_models.ServiceRegistrationKey, clientv3.WithPrefix())
	for {
		select {
		case <-kc.Ctx().Done():
			return
		case watchResp := <-watchGrpcCh:
			for _, event := range watchResp.Events {
				switch event.Type {
				case clientv3.EventTypePut:
					item := new(kk_etcd_models.PBServiceRegistration)
					err := item.UnMarshal(event.Kv.Value)
					if err != nil {
						slog.Error("watch EventTypePut", kk_stage.NewLog(nil).Error(err).Args()...)
						continue
					}
					//new registration putToDeadHub first
					x.putToDeadHub(item)
				case clientv3.EventTypeDelete:
					//paser HubKey from UniqueKey(event.Kv.Key)
					item := new(kk_etcd_models.PBServiceRegistration)
					err := item.BuildFromUniqueKey(string(event.Kv.Key))
					if err != nil {
						slog.Error("watch EventTypeDelete BuildFromUniqueKey", kk_stage.NewLog(nil).Error(err).Args()...)
						continue
					}
					x.delFromDeadHub(item)
					x.delFromAliveHub(item)
					stopCheck(item.UniqueKey())
				}
			}
		}
	}
}

func getOneAliveService(connType kk_etcd_models.PBServiceType, serviceName string) (*kk_etcd_models.PBServiceRegistration, error) {
	aliveHub, err := getAliveHubFromEtcd()
	if err != nil {
		return nil, err
	}
	addresses, ok := aliveHub.Get(connType.HubKey(serviceName))
	if !ok {
		return nil, kk_etcd_error.ErrNoAvailableConn
	}
	shuffle := addresses.Shuffle()
	for _, item := range shuffle {
		switch item.CheckConfig.Type {
		case kk_etcd_models.PBServiceType_Grpc:
			err := checkGrpc(item.CheckConfig)
			if err != nil {
				continue
			}
			return item, nil
		case kk_etcd_models.PBServiceType_Http:
			err := checkHttp(item.CheckConfig)
			if err != nil {
				continue
			}
			return item, nil
		default:
			continue
		}
	}
	return nil, kk_etcd_error.ErrNoAvailableConn
}

func putAliveHubToEtcd(hub *hubT) {
	err := util_kv.PutExistUpdateJson(kc.Client, kk_etcd_models.ServiceAliveKey, hub)
	if err != nil {
		slog.Error("failed to put to etcd", kk_stage.NewLog(nil).Error(err).Args()...)
		return
	}
}

func getAliveHubFromEtcd() (*hubT, error) {
	newHub := newHubT()
	err := util_kv.GetJson(kc.Client, kk_etcd_models.ServiceAliveKey, newHub)
	if err != nil {
		slog.Error("failed to get from etcd", kk_stage.NewLog(nil).Error(err).Args()...)
		return nil, err
	}
	return newHub, nil
}

func delFromHub(hub *hubT, item *kk_etcd_models.PBServiceRegistration) {
	conns, ok := hub.Get(item.HubKey())
	if ok {
		conns.DeleteFunc(func(c *kk_etcd_models.PBServiceRegistration) bool {
			return c.UniqueKey() == item.UniqueKey()
		})
	}
}

func putToHub(hub *hubT, item *kk_etcd_models.PBServiceRegistration) (alreadyExists bool) {
	conns, ok := hub.Get(item.HubKey())
	if ok {
		firstIndex := conns.FirstIndexFunc(func(c *kk_etcd_models.PBServiceRegistration) bool {
			return c.Equal(item)
		})
		if firstIndex != -1 {
			return true
		} else {
			//replace if exists, append if not
			conns.ReplaceAllOrAppend(item, func(r *kk_etcd_models.PBServiceRegistration) bool {
				//check key only
				return r.HubKey() == item.HubKey()
			})
		}
	} else {
		newConns := kk_sync.NewSlice[*kk_etcd_models.PBServiceRegistration]()
		newConns.Append(item)
		hub.Add(item.HubKey(), newConns)
	}
	return false
}
