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
	rwLock        sync.RWMutex
	clients       []*T
	serviceType   string
	serviceName   string
	clientBuilder func(grpcConn grpc.ClientConnInterface) (client T)
}

// NewClientHub create a new ClientHub
func NewClientHub[T any](
	serviceType string,
	serviceName string,
	clientBuilder func(grpcConn grpc.ClientConnInterface) (client T)) *ClientHub[T] {
	return &ClientHub[T]{
		clients:       make([]*T, 0),
		serviceType:   serviceType,
		serviceName:   serviceName,
		clientBuilder: clientBuilder,
	}
}

// GetClient get a random grpc client
func (c *ClientHub[T]) GetClient(stage *kk_stage.Stage) *T {
	c.rwLock.RLock()
	if len(c.clients) == 0 {
		slog.Error("no grpc client available, attempt to proactively obtain service", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args...)
		serverList, err := kk_etcd.ServerList(c.serviceType + "/" + c.serviceName)
		if err != nil {
			slog.Error("ServerList failed", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args...)
			return nil
		} else {
			c.rwLock.RUnlock()
			c.refreshClients(stage, serverList)
			c.rwLock.RLock()
			if len(c.clients) == 0 {
				slog.Error("still no grpc client available", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args...)
				return nil
			}
		}
	}
	// get a random client
	randomIndex := rand.Intn(len(c.clients))
	client := c.clients[randomIndex]
	c.rwLock.RUnlock()
	return client
}

// ListenServerChange use ctx cancelFunc to stop Listening
func (c *ClientHub[T]) ListenServerChange(ctx context.Context, stage *kk_stage.Stage) error {

	slog.Info("start watch server list", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args...)
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	err := kk_etcd.WatchServerList(ctx, c.serviceType+"/"+c.serviceName, serverListChan)
	if err != nil {
		slog.Error("WatchServerList failed", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args...)
		return err
	}

	go func() {
		for {
			//slog.Info("server list changed", logBody.Args...)
			select {
			case <-ctx.Done():
				slog.Info("ctx done stop ListenServerChange, close serverListChan", kk_stage.NewLogArgs(stage).Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args...)
				close(serverListChan)
				return
			case serverList := <-serverListChan:
				go c.refreshClients(stage, serverList)
			}
		}
	}()
	return nil
}

// refreshClients refresh grpc clients
func (c *ClientHub[T]) refreshClients(stage *kk_stage.Stage, serverList *kk_etcd_models.PBListServer) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	//slog.Info("serverListChan", "serverList", serverList)
	c.clients = make([]*T, 0)
	for _, server := range serverList.ListServer {
		// Set up a connection to the server_grpc_im_msg.
		grpcConn, err := grpc.Dial(server.ServiceAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		)
		if err != nil {
			slog.Error("grpc client connect failed", kk_stage.NewLogArgs(stage).Error(err).Any("target", server.ServiceAddr).Args...)
			return
		}
		client := c.clientBuilder(grpcConn)
		c.clients = append(c.clients, &client)
	}
}
