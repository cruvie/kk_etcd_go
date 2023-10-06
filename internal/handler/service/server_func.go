package service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
	"net/http"
	"time"
)

type serverFunc struct{}

var toolServer serverFunc

func (t *serverFunc) registerServer(registration *kk_etcd_models.ServiceRegistration) error {
	key := registration.ServerType + registration.ServerName

	endpointManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, key)
	if err != nil {
		slog.Error("failed to create etcd manager", "err", err)
		return err
	}

	lease, err := kk_etcd_client.EtcdClient.Grant(registration.Context, registration.Check.TTL)
	if err != nil {
		slog.Error("failed to create lease", "err", err)
		return err
	}

	endpointKey := key + "/" + registration.Address
	//add endpoint to etcd with lease id
	err = endpointManager.AddEndpoint(registration.Context, endpointKey, endpoints.Endpoint{Addr: registration.Address}, clientv3.WithLease(lease.ID))
	if err != nil {
		slog.Error("failed to add endpoint to etcd", "err", err)
		return err
	}

	//start a goroutine to keep alive
	go func(registration *kk_etcd_models.ServiceRegistration, endpointManager endpoints.Manager, endpointKey string, lease *clientv3.LeaseGrantResponse) {
		defer slog.Info("keep alive goroutine exit", "endpointKey", endpointKey)
		if err := t.keepAliveOnce(registration.Context, lease.ID); err != nil {
			return
		}
		ticker := time.NewTicker(time.Duration(registration.Check.Interval) * time.Second)
		failCount := 0
		for {
			select {
			case <-ticker.C:
				if err := t.keepAliveOnce(registration.Context, lease.ID); err != nil {
					return
				}
				if ok := t.checkHealth(registration); ok {
					failCount = 0
				} else {
					failCount++
				}
				if failCount == 3 {
					slog.Info("service not healthy, delete endpoint", "endpointKey", endpointKey)
					t.deleteEndpointAndRevokeLease(registration.Context, endpointManager, endpointKey, lease.ID)
					return
				}
			case <-registration.Context.Done():
				slog.Info("context done", "endpointKey", endpointKey)
				t.deleteEndpointAndRevokeLease(registration.Context, endpointManager, endpointKey, lease.ID)
				return
			}
		}
	}(registration, endpointManager, endpointKey, lease)
	return nil
}

func (t *serverFunc) keepAliveOnce(context context.Context, leaseID clientv3.LeaseID) error {
	_, err := kk_etcd_client.EtcdClient.KeepAliveOnce(context, leaseID)
	if err != nil {
		slog.Error("failed to set keep alive", "err", err)
		return err
	}
	return nil
}

func (t *serverFunc) checkHealth(registration *kk_etcd_models.ServiceRegistration) (ok bool) {
	if registration.Check.HTTP != "" {
		var httpClient = &http.Client{
			Timeout: time.Duration(registration.Check.Timeout) * time.Second,
		}
		defer httpClient.CloseIdleConnections()
		// send http request to addr
		req, err := http.NewRequest(registration.Check.HTTPMethod, registration.Check.HTTP, nil)
		if err != nil {
			return false
		}
		resp, err := httpClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			return true
		}
		if err != nil {
			slog.Error("failed to check http health", "err", err)
			return false
		}
	} else if registration.Check.GRPC != "" {
		conn, err := grpc.Dial(registration.Check.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			slog.Error("failed to dial grpc", "err", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				slog.Error("failed to close grpc connection", "err", err)
			}
		}(conn)
		healthClient := grpc_health_v1.NewHealthClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(registration.Check.Timeout)*time.Second)
		defer cancel()
		resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			slog.Error("failed to check grpc health", "err", err)
			return false
		}
		status := resp.GetStatus()
		if status == grpc_health_v1.HealthCheckResponse_SERVING {
			return true
		}
	}
	return false
}

func (t *serverFunc) deleteEndpointAndRevokeLease(ctx context.Context, endpointManager endpoints.Manager, endpointKey string, leaseID clientv3.LeaseID) {
	_, err := kk_etcd_client.EtcdClient.Revoke(ctx, leaseID)
	if err != nil {
		slog.Error("failed to revoke lease", "err", err)
	}
	err = endpointManager.DeleteEndpoint(ctx, endpointKey)
	if err != nil {
		slog.Error("failed to delete endpoint", "err", err)
	}
}
