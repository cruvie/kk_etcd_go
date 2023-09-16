package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log/slog"
)

// registerServer
// serviceType: key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// addr: 192.123.32.11:8080
// serverName:
// ttl:lease time, default 15 seconds, service will be deleted from etcd after ttl seconds
func registerServer(ctx context.Context, serviceType, addr, serverName string, ttl ...int64) {
	key := serviceType + serverName
	if serverName == "" {
		slog.Error("serverName can not be empty")
	}
	if len(ttl) == 0 {
		ttl = append(ttl, 15)
	}
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, key)
	if err != nil {
		slog.Error("failed to create etcd manager", "err", err)
	}

	lease, err := kk_etcd_client.EtcdClient.Grant(ctx, ttl[0])
	if err != nil {
		slog.Error("failed to create lease", "err", err)
	}

	//add endpoint to etcd with lease id
	err = etcdManager.AddEndpoint(ctx, key+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		slog.Error("failed to add endpoint to etcd", "err", err)
	}

	//every ttl/3 seconds to renew lease
	respChan, err := kk_etcd_client.EtcdClient.KeepAlive(ctx, lease.ID)
	if err != nil {
		slog.Error("failed to set keep alive", "err", err)
	}

	//start a goroutine to listen keep alive result until context done
	go func() {
		for {
			select {
			case resp := <-respChan:
				if resp == nil {
					//slog.Error("keep alive channel closed")
					return
				} else {
					//slog.Info("keep alive", "resp",resp)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}
