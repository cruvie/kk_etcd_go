package server_hub

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
	"time"
)

func InitServiceHub() {
	hub = &serverHub{
		hub: kk_sync.NewMap[string, *serverStatus](),
	}
	hub.watchServiceChange()
}

var hub *serverHub

type serverHub struct {
	hub *kk_sync.Map[string, *serverStatus]
}

func (x *serverHub) register(server *serverStatus) {
	server.ctx, server.cancelFunc = context.WithCancel(context.Background())
	x.hub.Add(server.kVKey(), server)
	err := server.putExistUpdateJson()
	if err != nil {
		if kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCNoSpace) {
			//https://etcd.io/docs/v3.5/op-guide/maintenance/#space-quota
			slog.Warn("register server success but has no space warning", kk_log.NewLog(nil).
				String("func", " PutExistUpdateJson()").
				Error(err).Args()...)
		} else {
			slog.Error("register server failed PutExistUpdateJson()", kk_log.NewLog(nil).
				Error(err).Any("server", server).Args()...)
			return
		}
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

func (x *serverHub) deregister(endpointKey string) {
	key := kk_etcd_models.InternalServerStatus + endpointKey
	oldStatus, ok := x.hub.Get(key)
	if ok {
		x.hub.Remove(key)
		//stop check
		oldStatus.stopCheck()
	}
	err := oldStatus.kVDelWithKey(key)
	if err != nil {
		slog.Error("deregister server failed server.KVDel()", kk_log.NewLog(nil).Error(err).Any("server", oldStatus).Args()...)
	}
}
func (x *serverHub) updateStatus(status kk_etcd_models.PBServer_ServerStatus, server *serverStatus) {
	v, ok := x.hub.Get(server.kVKey())
	if ok {
		// update server status
		v.Status = status
		v.LastCheck = time.Now()
		v.Msg = server.Msg
		x.hub.Add(v.kVKey(), v)
		err := v.putExistUpdateJson()
		if err != nil {
			if kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCNoSpace) {
				//https://etcd.io/docs/v3.5/op-guide/maintenance/#space-quota
				slog.Warn("register server success but has no space warning", kk_log.NewLog(nil).
					String("func", " PutExistUpdateJson()").
					Error(err).Args()...)
			} else {
				slog.Error("register server failed PutExistUpdateJson()", kk_log.NewLog(nil).
					Error(err).Any("server", server).Args()...)
				return
			}
		}
	} else {
		slog.Error("update server status failed server not found",
			kk_log.NewLog(nil).Any("server", server).Args()...)
	}
}

// watchServiceChange watch server change
func (x *serverHub) watchServiceChange() {

	getChannel := func(serverType kk_etcd_models.ServerType) endpoints.WatchChannel {
		etcdManager, err := serverType.NewEndpointManager(global_model.GetClient(internal_client.GlobalStage))
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
				var status serverStatus
				status.fromEndpoint(update.Endpoint)
				switch update.Op {
				case endpoints.Delete:
					//delete op only has key info
					hub.deregister(update.Key)
				case endpoints.Add:
					hub.register(&status)
				}
			}()
		}
	}
	go func() {
		for {
			select {
			case <-internal_client.GlobalStage.Ctx.Done():
				return
			case updates, ok := <-grpcChannel:
				if !ok {
					return
				}
				updateEndpoint(updates)
			case updates, ok := <-httpChannel:
				if !ok {
					return
				}
				updateEndpoint(updates)
			}
		}
	}()
}
