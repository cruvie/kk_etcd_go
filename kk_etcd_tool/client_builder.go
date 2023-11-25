package kk_etcd_tool

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
	rwLock          sync.RWMutex
	clients         []*T
	serviceType     string
	serviceName     string
	grpcConnBuilder func(grpcConn grpc.ClientConnInterface) (client T)
}

func NewClientHub[T any](
	serviceType string,
	serviceName string,
	grpcConnBuilder func(grpcConn grpc.ClientConnInterface) (client T)) *ClientHub[T] {
	return &ClientHub[T]{
		serviceType:     serviceType,
		serviceName:     serviceName,
		grpcConnBuilder: grpcConnBuilder,
	}
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

func (c *ClientHub[T]) ListenServerChange(stage *kk_stage.Stage) {
	logBody := kk_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("serviceType", c.serviceType).SetAny("serviceName", c.serviceName)
	slog.Info("start watch server list", logBody.GetLogArgs()...)
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	defer close(serverListChan)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		err := kk_etcd.WatchServerList(ctx, c.serviceType+"/"+c.serviceName, serverListChan)
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
			c.rwLock.Lock()
			//slog.Info("serverListChan", "serverList", serverList)
			clear(c.clients)
			for _, server := range serverList.ListServer {
				client := c.buildGrpcClient(stage, server.ServiceAddr)
				if client != nil {
					c.clients = append(c.clients, client)
				}
			}
			c.rwLock.Unlock()
		}
	}
}

// buildGrpcClient target=ip+port
func (c *ClientHub[T]) buildGrpcClient(stage *kk_stage.Stage, target string) *T {
	// Set up a connection to the server_grpc_im_msg.
	grpcConn, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		logBody := kk_stage.NewLogBody().SetError(err).SetTraceId(stage.TraceId).SetAny("target", target)
		slog.Error("grpc client connect failed", logBody.GetLogArgs()...)
		return nil
	}
	client := c.grpcConnBuilder(grpcConn)
	return &client
}
