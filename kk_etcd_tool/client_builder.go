package kk_grpc

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"math/rand"
	"sync"
)

type ClientHub[T any] struct {
	rwLock  sync.RWMutex
	clients []*T
}

func NewClientHub[T any]() *ClientHub[T] {
	return &ClientHub[T]{}
}

func (c *ClientHub[T]) GetGrpcClient() *T {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if len(c.clients) == 0 {
		return nil
	}
	// get a random client
	randomIndex := rand.Intn(len(c.clients))
	client := c.clients[randomIndex]
	return client
}

func ListenServerChange[T any](stage *kk_stage.Stage, serviceType string, serviceName string, clientHub *ClientHub[T], buildFunc func(grpcConn grpc.ClientConnInterface) (client T)) {
	logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("serviceType", serviceType).SetAny("serviceName", serviceName)
	slog.Info("start watch server list", logBody.GetLogArgs()...)
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	defer close(serverListChan)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		err := kk_etcd.WatchServerList(ctx, serviceType+"/"+serviceName, serverListChan)
		if err != nil {
			logBody.SetError(err)
			slog.Error("WatchServerList failed", logBody.GetLogArgs()...)
			panic(nil)
		}
	}()
	for {
		//slog.Info("server list changed", logBody.GetLogArgs()...)
		select {
		case <-ctx.Done():
			slog.Info("ctx done")
			return
		case serverList := <-serverListChan:
			clientHub.rwLock.Lock()
			//slog.Info("serverListChan", "serverList", serverList)
			clear(clientHub.clients)
			for _, server := range serverList.ListServer {
				client := buildGrpcClient[T](stage, server.ServiceAddr, buildFunc)
				if client != nil {
					clientHub.clients = append(clientHub.clients, client)
				}
			}
			clientHub.rwLock.Unlock()
		}
	}
}

// buildGrpcClient target=ip+port
func buildGrpcClient[T any](stage *kk_stage.Stage, target string, buildFunc func(grpcConn grpc.ClientConnInterface) (client T)) *T {
	// Set up a connection to the server_grpc_im_msg.
	grpcConn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		logBody := kk_stage.NewLogBody().SetError(err).SetTraceId(stage.TraceId)
		slog.Error("grpc client connect failed", logBody.GetLogArgs()...)
		return nil
	}
	client := buildFunc(grpcConn)
	return &client
}
