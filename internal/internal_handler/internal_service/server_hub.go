package internal_service

import (
	"context"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
	"time"
)

func InitServiceHub() {
	hub = &serverHub{
		hub: kk_sync.NewMap[*serverStatus](),
	}
	hub.watchServiceChange()
}

var hub *serverHub

type serverHub struct {
	hub *kk_sync.Map[*serverStatus]
}

func (x *serverHub) register(server *serverStatus) {
	if v, ok := x.hub.Get(server.Key); ok {
		//stop checker
		v.stopCheck()
	}
	server.ctx, server.cancelFunc = context.WithCancel(context.Background())
	x.hub.Add(server.Key, server)
	err := server.PutExistUpdateJson()
	if err != nil {
		slog.Error("register server failed server.PutExistUpdateJson()", kk_log.NewLog(nil).
			Error(err).Any("server", server).Args()...)
		return
	}
	err = server.Check()
	if err != nil {
		slog.Error("server may not be registered by kk_etcd", kk_log.NewLog(nil).
			Error(err).Any("server", server).Args()...)
		//no need to run checker
		return
	}
	//start checker
	go server.runCheck()
}

func (x *serverHub) deregister(key string) {
	v, ok := x.hub.Get(key)
	if ok {
		x.hub.Remove(key)
		//stop checker
		v.stopCheck()
	}
	err := v.KVDel()
	if err != nil {
		slog.Error("deregister server failed server.KVDel()", kk_log.NewLog(nil).Error(err).Any("server", v).Args()...)
	}
}
func (x *serverHub) updateStatus(status kk_etcd_models.PBServer_ServerStatus, server *serverStatus) error {
	v, ok := x.hub.Get(server.Key)
	if ok {
		// update server status
		v.Status = status
		v.LastCheck = time.Now()
		v.Msg = server.Msg
		x.hub.Add(v.Key, v)
		err := v.PutExistUpdateJson()
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("server not found %s", server.Key)
	}
}

func (x *serverHub) services() (map[string]*serverStatus, error) {
	services := make(map[string]*serverStatus)
	err, v := service.SerKV{}.KVList(internal_client.GlobalStage, kk_etcd_models.InternalServerStatus)
	if err != nil {
		return nil, err
	}
	for _, kv := range v.GetListKV() {
		var info serverStatus
		err := info.FromJson(kv.GetValue())
		if err != nil {
			return nil, err
		}
		services[info.Key] = &info
	}

	return services, nil
}

// watchServiceChange watch server change
func (x *serverHub) watchServiceChange() {

	getChannel := func(serverType kk_etcd_models.ServerType) endpoints.WatchChannel {
		etcdManager, err := endpoints.NewManager(global_model.GetClient(internal_client.GlobalStage), serverType.String())
		if err != nil {
			panic(err)
		}
		channel, err := etcdManager.NewWatchChannel(internal_client.GlobalStage.Ctx)
		if err != nil {
			panic(err)
		}
		return channel
	}
	grpcChannel := getChannel(kk_etcd_models.Grpc)
	httpChannel := getChannel(kk_etcd_models.Http)

	updateEndpoint := func(updates []*endpoints.Update) {
		for _, update := range updates {
			go func() {
				switch update.Op {
				case endpoints.Delete:
					hub.deregister(update.Key)
				case endpoints.Add:
					var info serverStatus
					info.fromEndpoint(update.Endpoint)
					hub.register(&info)
				}
			}()
		}
	}
	go func() {
		for {
			select {
			case <-internal_client.GlobalStage.Ctx.Done():
				return
			case updates := <-grpcChannel:
				updateEndpoint(updates)
			case updates := <-httpChannel:
				updateEndpoint(updates)
			}
		}
	}()
}
