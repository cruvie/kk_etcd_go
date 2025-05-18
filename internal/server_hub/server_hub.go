package server_hub

import (
	"bytes"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"sync"
	"time"
)

type hubT struct {
	*kk_sync.Map[string, *kk_sync.Slice[*kk_etcd_models.PBServerRegistration]]
}

func newHubT() *hubT {
	return &hubT{
		Map: kk_sync.NewMap[string, *kk_sync.Slice[*kk_etcd_models.PBServerRegistration]](),
	}
}

func (x *hubT) MarshalJSON() ([]byte, error) {
	m := make(map[string][]*kk_etcd_models.PBServerRegistration)
	data := x.Data()
	for k, v := range data {
		m[k] = v.Slice()
	}

	return json.Marshal(m)
}

func (x *hubT) UnmarshalJSON(b []byte) error {
	tmp := make(map[string][]*kk_etcd_models.PBServerRegistration)
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&tmp)
	if err != nil {
		return err
	}

	for k, v := range tmp {
		slice := kk_sync.NewSlice[*kk_etcd_models.PBServerRegistration]()
		slice.Append(v...)
		x.Add(k, slice)
	}

	return nil
}

var hub *serverHub

// todo 服务在删除时正在检查会有并发问题，检查后可能又被重新注册
var hubLock sync.RWMutex //nolint

type serverHub struct {
	aliveHub *hubT
	deadHub  *hubT
}

func InitConnHub() {
	hub = &serverHub{
		aliveHub: newHubT(),
		deadHub:  newHubT(),
	}
	hub.watch()
	time.Sleep(1 * time.Second)

	list, err := serServerInstance.serverList(kk_etcd_models.ServerRegistrationKey)
	if err != nil {
		panic(err)
	}
	for _, registration := range list {
		hub.putToDeadHub(registration)
	}
}

func (x *serverHub) delFromDeadHub(item *kk_etcd_models.PBServerRegistration) {
	delFromHub(x.deadHub, item)
}

func (x *serverHub) putToDeadHub(item *kk_etcd_models.PBServerRegistration) {
	putToHub(x.deadHub, item)
	startNewCheck(item)
}

func (x *serverHub) delFromAliveHub(item *kk_etcd_models.PBServerRegistration) {
	delFromHub(x.aliveHub, item)
}

func (x *serverHub) putToAliveHub(item *kk_etcd_models.PBServerRegistration) {
	putToHub(x.aliveHub, item)
	//after put into aliveHub, put all alive conn to etcd for App outside to get conn
	putHubToEtcd(kk_etcd_models.ServerAliveKey, x.aliveHub)
}

func putHubToEtcd(key string, hub *hubT) {
	err := util_kv.PutExistUpdateJson(kc.Client, key, hub)
	if err != nil {
		slog.Error("failed to put to etcd", kk_log.NewLog(nil).Error(err).Args()...)
		return
	}
}

func getHubFromEtcd(key string) (*hubT, error) {
	newHub := newHubT()
	err := util_kv.GetJson(kc.Client, key, newHub)
	if err != nil {
		slog.Error("failed to get from etcd", kk_log.NewLog(nil).Error(err).Args()...)
		return nil, err
	}
	return newHub, nil
}

func delFromHub(hub *hubT, item *kk_etcd_models.PBServerRegistration) {
	conns, ok := hub.Get(item.ServerName)
	if ok {
		conns.DeleteFunc(func(c *kk_etcd_models.PBServerRegistration) bool {
			return c.Key() == item.Key()
		})
	}
}

func putToHub(hub *hubT, item *kk_etcd_models.PBServerRegistration) {
	conns, ok := hub.Get(item.ServerName)
	if ok {
		//replace if exists, append if not
		conns.ReplaceOrAppend(item, func(r *kk_etcd_models.PBServerRegistration) bool {
			return r.Key() == item.Key()
		})
	} else {
		newConns := kk_sync.NewSlice[*kk_etcd_models.PBServerRegistration]()
		newConns.Append(item)
		hub.Add(item.ServerName, newConns)
	}
}

// need this for SDK use, which cannot access memory data of running kk_etcd
func (x *serverHub) watch() {
	watchGrpcCh := kc.Watch(kc.Ctx(), kk_etcd_models.ServerRegistrationKey, clientv3.WithPrefix())
	go x.watchCh(watchGrpcCh)
}

func (x *serverHub) watchCh(watchCh clientv3.WatchChan) {
	for {
		select {
		case <-kc.Ctx().Done():
			return
		case watchResp := <-watchCh:
			for _, event := range watchResp.Events {
				item := new(kk_etcd_models.PBServerRegistration)
				err := item.UnMarshal(event.Kv.Value)
				if err != nil {
					slog.Error("watch EventTypePut", kk_log.NewLog(nil).Error(err).Args()...)
					continue
				}
				switch event.Type {
				case clientv3.EventTypePut:
					//new registration putToDeadHub first
					x.putToDeadHub(item)
				case clientv3.EventTypeDelete:
					x.delFromDeadHub(item)
					x.delFromAliveHub(item)
					stopCheck(item.Key())
				}
			}
		}
	}
}

func getOneAliveServer(connType kk_etcd_models.PBServerType, serverName string) (*kk_etcd_models.PBServerRegistration, error) {
	aliveHub, err := getHubFromEtcd(kk_etcd_models.ServerAliveKey)
	if err != nil {
		return nil, err
	}
	addresses, ok := aliveHub.Get(serverName)
	if !ok {
		return nil, kk_etcd_error.ErrNoAvailableConn
	}
	//filter connType
	addresses.DeleteFunc(func(registration *kk_etcd_models.PBServerRegistration) bool {
		return registration.ServerType != connType
	})
	shuffle := addresses.Shuffle()
	for _, item := range shuffle {
		switch item.CheckConfig.Type {
		case kk_etcd_models.PBServerType_Grpc:
			err := checkGrpc(item.CheckConfig)
			if err != nil {
				continue
			}
			return item, nil
		case kk_etcd_models.PBServerType_Http:
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
