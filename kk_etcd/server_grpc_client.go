package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
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
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	c.rwLock.RLock()
	if len(c.clients) == 0 {
		slog.Error("no grpc client available, attempt to proactively obtain service", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args()...)
		err, serverList := hServer.ServerList(stage, &kk_etcd_models.ServerListParam{Prefix: c.serviceType + "/" + c.serviceName})
		if err != nil {
			slog.Error("ServerList failed", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args()...)
			return nil
		} else {
			c.rwLock.RUnlock()
			c.refreshClients(stage, serverList.GetServerList())
			c.rwLock.RLock()
			if len(c.clients) == 0 {
				slog.Error("still no grpc client available", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args()...)
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
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	slog.Info("start watch server list", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args()...)
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	err := WatchServerList(ctx, c.serviceType+"/"+c.serviceName, serverListChan)
	if err != nil {
		slog.Error("WatchServerList failed", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Error(err).Args()...)
		return err
	}

	go func() {
		for {
			//slog.Info("server list changed", logBody.Args()...)
			select {
			case <-ctx.Done():
				slog.Info("ctx done stop ListenServerChange, close serverListChan", newLog.Any("serviceType", c.serviceType).Any("serviceName", c.serviceName).Args()...)
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
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	//slog.Info("serverListChan", "serverList", serverList)
	c.clients = make([]*T, 0)
	for _, server := range serverList.ListServer {
		// Set up a connection to the server_grpc_im_msg.
		grpcConn, err := grpc.NewClient(server.ServiceAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		)
		if err != nil {
			slog.Error("grpc client connect failed", newLog.Error(err).Any("target", server.ServiceAddr).Args()...)
			return
		}
		client := c.clientBuilder(grpcConn)
		c.clients = append(c.clients, &client)
	}
}
