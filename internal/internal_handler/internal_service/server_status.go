package internal_service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
	"net/http"
	"time"
)

type serverStatus struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	Key        string
	kk_etcd_models.CheckConfig
	Status    kk_etcd_models.PBServer_ServerStatus
	LastCheck time.Time
	Msg       string
	Metadata  any
}

func (x *serverStatus) KVKey() string {
	return kk_etcd_models.InternalServerStatus + x.Key
}
func (x *serverStatus) PutExistUpdateJson() error {
	return service.SerKV{}.PutExistUpdateJson(internal_client.GlobalStage, x.KVKey(), x)
}
func (x *serverStatus) GetJson() (err error) {
	err = service.SerKV{}.GetJson(internal_client.GlobalStage, x.KVKey(), x)
	return err
}
func (x *serverStatus) FromJson(data string) (err error) {
	err = json.Unmarshal([]byte(data), x)
	return err
}
func (x *serverStatus) KVDel() error {
	return service.SerKV{}.KVDel(internal_client.GlobalStage, x.KVKey())
}

// fromEndpoint endpoint to serverStatus
// todo https://github.com/etcd-io/etcd/issues/18520

// serverStatus, ok := endpoint.Metadata.(serverStatus)
func (x *serverStatus) fromEndpoint(endpoint endpoints.Endpoint) {
	x.Status = kk_etcd_models.PBServer_UnKnown
	bytes, err := json.Marshal(endpoint.Metadata)
	if err != nil {
		slog.Error("fromEndpoint Endpoint metadata", kk_log.NewLog(nil).Error(err).Args()...)
		return
	}
	err = json.Unmarshal(bytes, x)
	if err != nil {
		slog.Error("fromEndpoint Endpoint metadata", kk_log.NewLog(nil).Error(err).Args()...)
	}
}

func newServerStatus(key string,
	checkConfig kk_etcd_models.CheckConfig,
	metadata any) *serverStatus {
	return &serverStatus{
		Key:         key,
		CheckConfig: checkConfig,
		Status:      kk_etcd_models.PBServer_Init,
		LastCheck:   kk_time.DefaultTime,
		Msg:         "init checker",
		Metadata:    metadata,
	}
}

func (x *serverStatus) stopCheck() {
	if x.cancelFunc != nil {
		x.cancelFunc()
	}
}
func (x *serverStatus) runCheck() {
	defer func() {
		x.Msg = "close checker"
		err := hub.updateStatus(kk_etcd_models.PBServer_UnKnown, x)
		if err != nil {
			slog.Error("close checker", kk_log.NewLog(nil).Error(err).Args()...)
		}
	}()

	updateHealth := func(err error) error {
		if err != nil {
			x.Msg = err.Error()
			err := hub.updateStatus(kk_etcd_models.PBServer_Stop, x)
			if err != nil {
				return err
			}
		} else {
			x.Msg = "Health check success!"
			err := hub.updateStatus(kk_etcd_models.PBServer_Running, x)
			if err != nil {
				return err
			}
		}
		return nil
	}

	ticker := time.NewTicker(x.CheckConfig.Interval)
	switch x.CheckConfig.Type {
	case kk_etcd_models.Grpc:
		for {
			select {
			case <-x.ctx.Done():
				return
			case <-ticker.C:
				err := x.checkGrpc(&x.CheckConfig)
				if err != nil {
					slog.Warn("checkGrpc", kk_log.NewLog(nil).Error(err).Args()...)
				}
				err = updateHealth(err)
				if err != nil {
					slog.Error("updateHealth", kk_log.NewLog(nil).Error(err).Args()...)
					return
				}
			}
		}
	case kk_etcd_models.Http:
		for {
			select {
			case <-x.ctx.Done():
				return
			case <-ticker.C:
				err := x.checkHttp(&x.CheckConfig)
				if err != nil {
					slog.Warn("checkHttp", kk_log.NewLog(nil).Error(err).Args()...)
				}
				err = updateHealth(err)
				if err != nil {
					slog.Error("updateHealth", kk_log.NewLog(nil).Error(err).Args()...)
					return
				}
			}
		}
	}
}

func (x *serverStatus) checkGrpc(checkConfig *kk_etcd_models.CheckConfig) (err error) {
	conn, err := grpc.NewClient(checkConfig.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient failed")
	}
	defer func(conn *grpc.ClientConn) {
		if conn == nil {
			return
		}
		err := conn.Close()
		if err != nil {
			msg := "failed to close grpc connection"
			slog.Error(msg, kk_log.NewLog(nil).Error(err).Args()...)
		}
	}(conn)
	healthClient := grpc_health_v1.NewHealthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), checkConfig.Timeout)
	defer cancel()
	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return fmt.Errorf("failed to check grpc health %w", err)
	}
	status := resp.GetStatus()
	if status == grpc_health_v1.HealthCheckResponse_SERVING {
		return nil
	} else {
		return fmt.Errorf("grpc status code is not ok, status:%d", status)
	}
}

func (x *serverStatus) checkHttp(checkConfig *kk_etcd_models.CheckConfig) (err error) {
	var httpClient = &http.Client{
		Timeout: checkConfig.Timeout,
	}
	defer httpClient.CloseIdleConnections()
	// send http request to addr
	req, err := http.NewRequest("GET", checkConfig.Addr, nil)
	if err != nil {
		return fmt.Errorf("failed to new http request %w", err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to check http health %w", err)
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("http status code is not ok, code:%d", resp.StatusCode)
	}
}
