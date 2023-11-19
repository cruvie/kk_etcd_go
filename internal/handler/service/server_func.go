package service

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
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

func (t *serverFunc) registerServer(stage *kku_stage.Stage, registration *kk_etcd_models.ServiceRegistration) error {
	key := registration.ServerType + "/" + registration.ServerName

	endpointManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, key)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		msg := "failed to create etcd manager"
		slog.Error(msg, logBody.GetLogArgs()...)
		return err
	}

	lease, err := kk_etcd_client.EtcdClient.Grant(registration.Context, registration.Check.TTL)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		msg := "failed to create lease"
		slog.Error(msg, logBody.GetLogArgs()...)
		return err
	}

	endpointKey := key + "/" + registration.Address
	//add endpoint to etcd with lease id
	err = endpointManager.AddEndpoint(registration.Context, endpointKey, endpoints.Endpoint{Addr: registration.Address}, clientv3.WithLease(lease.ID))
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		msg := "failed to add endpoint to etcd"
		slog.Error(msg, logBody.GetLogArgs()...)
		return err
	}

	//start a goroutine to keep alive
	go func(registration *kk_etcd_models.ServiceRegistration, endpointManager endpoints.Manager, endpointKey string, lease *clientv3.LeaseGrantResponse) {
		defer func() {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetAny("endpointKey", endpointKey)
			msg := "keep alive goroutine exit"
			slog.Info(msg, logBody.GetLogArgs()...)
		}()
		if err := t.keepAliveOnce(stage, registration.Context, lease.ID); err != nil {
			return
		}
		ticker := time.NewTicker(time.Duration(registration.Check.Interval) * time.Second)
		failCount := 0
		for {
			select {
			case <-ticker.C:
				if err := t.keepAliveOnce(stage, registration.Context, lease.ID); err != nil {
					return
				}
				if ok := t.checkHealth(stage, registration); ok {
					failCount = 0
				} else {
					failCount++
				}
				if failCount == 3 {
					slog.Warn("service not healthy, delete endpoint", "endpointKey", endpointKey)
					t.deleteEndpointAndRevokeLease(stage, registration.Context, endpointManager, endpointKey, lease.ID)
					return
				}
			case <-registration.Context.Done():
				slog.Warn("context done", "endpointKey", endpointKey)
				t.deleteEndpointAndRevokeLease(stage, registration.Context, endpointManager, endpointKey, lease.ID)
				return
			}
		}
	}(registration, endpointManager, endpointKey, lease)
	return nil
}

func (t *serverFunc) keepAliveOnce(stage *kku_stage.Stage, context context.Context, leaseID clientv3.LeaseID) error {
	_, err := kk_etcd_client.EtcdClient.KeepAliveOnce(context, leaseID)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		slog.Error("failed to set keep alive", logBody.GetLogArgs()...)
		return err
	}
	return nil
}

func (t *serverFunc) checkHealth(stage *kku_stage.Stage, registration *kk_etcd_models.ServiceRegistration) (ok bool) {
	if registration.Check.HTTP != "" {
		var httpClient = &http.Client{
			Timeout: time.Duration(registration.Check.Timeout) * time.Second,
		}
		defer httpClient.CloseIdleConnections()
		// send http request to addr
		req, err := http.NewRequest(registration.Check.HTTPMethod, registration.Check.HTTP, nil)
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
			msg := "failed to create http request"
			slog.Error(msg, logBody.GetLogArgs()...)
			return false
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
			msg := "failed to check http health"
			slog.Error(msg, logBody.GetLogArgs()...)
			return false
		} else if resp.StatusCode == http.StatusOK {
			return true
		} else {
			return false
		}
	} else if registration.Check.GRPC != "" {
		conn, err := grpc.Dial(registration.Check.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
			msg := "failed to dial grpc"
			slog.Error(msg, logBody.GetLogArgs()...)
			return false
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
				msg := "failed to close grpc connection"
				slog.Error(msg, logBody.GetLogArgs()...)
			}
		}(conn)
		healthClient := grpc_health_v1.NewHealthClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(registration.Check.Timeout)*time.Second)
		defer cancel()
		resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
			msg := "failed to check grpc health"
			slog.Error(msg, logBody.GetLogArgs()...)
			return false
		}
		status := resp.GetStatus()
		if status == grpc_health_v1.HealthCheckResponse_SERVING {
			return true
		}
	}
	return false
}

func (t *serverFunc) deleteEndpointAndRevokeLease(stage *kku_stage.Stage, ctx context.Context, endpointManager endpoints.Manager, endpointKey string, leaseID clientv3.LeaseID) {
	_, err := kk_etcd_client.EtcdClient.Revoke(ctx, leaseID)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		msg := "failed to revoke lease"
		slog.Error(msg, logBody.GetLogArgs()...)
	}
	err = endpointManager.DeleteEndpoint(ctx, endpointKey)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err)
		msg := "failed to delete endpoint"
		slog.Error(msg, logBody.GetLogArgs()...)
	}
}
