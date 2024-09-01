package internal_service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
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

var Checker checkInfo

type checkInfo struct {
	Key string
	kk_etcd_models.CheckConfig
	ServiceType kk_etcd_models.ServiceType
	// Status
	Status    kk_etcd_models.PBServer_ServerStatus
	LastCheck time.Time
	Msg       string
	//endpoints.Endpoint.Metadata
	Metadata any
}

var runningGo = kk_sync.NewMap[struct{}]()

func (x *checkInfo) checkHealth(info checkInfo) {
	_, exists := runningGo.Get(info.Key)
	if exists {
		return
	}
	runningGo.Add(info.Key, struct{}{})
	defer func() {
		runningGo.Remove(info.Key)
		info.Msg = "close checker"
		err := x.updateHealth(kk_etcd_models.PBServer_UnKnown, info)
		if err != nil {
			slog.Error("close checker", kk_log.NewLog(nil).Error(err).Args()...)
		}
	}()

	updateHealth := func(err error) error {
		if err != nil {
			info.Msg = err.Error()
			err := x.updateHealth(kk_etcd_models.PBServer_Stop, info)
			if err != nil {
				return err
			}
		} else {
			info.Msg = "Health check success!"
			err := x.updateHealth(kk_etcd_models.PBServer_Running, info)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := info.Check()
	if err != nil {
		err := updateHealth(fmt.Errorf("info.Check() service may not register by kk_etcd %w", err))
		if err != nil {
			slog.Error("updateHealth", kk_log.NewLog(nil).Error(err).Args()...)
			return
		}
		return
	}
	ticker := time.NewTicker(info.CheckConfig.Interval)
	switch info.CheckConfig.Type {
	case kk_etcd_models.CheckTypeGrpc:
		for {
			select {
			case <-ticker.C:
				err := x.checkGrpc(&info.CheckConfig)
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
	case kk_etcd_models.CheckTypeHttp:
		for {
			select {
			case <-ticker.C:
				err := x.checkHttp(&info.CheckConfig)
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

func (x *checkInfo) updateHealth(status kk_etcd_models.PBServer_ServerStatus, info checkInfo) error {
	info.Status = status
	info.LastCheck = time.Now()
	endpointManager, err := endpoints.NewManager(global_model.GetClient(internal_client.GlobalStage), info.ServiceType.String())
	if err != nil {
		return err
	}
	err = endpointManager.DeleteEndpoint(context.Background(), info.Key)
	if err != nil {
		return err
	}
	err = endpointManager.AddEndpoint(context.Background(), info.Key,
		endpoints.Endpoint{
			Addr:     info.Addr,
			Metadata: info,
		})
	return err
}

func (x *checkInfo) checkGrpc(checkConfig *kk_etcd_models.CheckConfig) (err error) {
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

func (x *checkInfo) checkHttp(checkConfig *kk_etcd_models.CheckConfig) (err error) {
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

// WatchServerChange watch server change
func (x *checkInfo) WatchServerChange() {

	getChannel := func(serviceGrpc kk_etcd_models.ServiceType) endpoints.WatchChannel {
		etcdManager, err := endpoints.NewManager(global_model.GetClient(internal_client.GlobalStage), serviceGrpc.String())
		if err != nil {
			panic(err)
		}
		channel, err := etcdManager.NewWatchChannel(internal_client.GlobalStage.Ctx)
		if err != nil {
			panic(err)
		}
		return channel
	}
	grpcChannel := getChannel(kk_etcd_models.ServiceGrpc)
	httpChannel := getChannel(kk_etcd_models.ServiceHttp)

	updateEndpoint := func(updates []*endpoints.Update) {
		for _, update := range updates {
			//filter close or open a health check
			if update.Op == endpoints.Delete {
				//ignore delete event
				continue
			}
			go func() {
				var info checkInfo
				err := info.Convert(update.Endpoint)
				if err != nil {
					slog.Error("Convert Endpoint metadata", kk_log.NewLog(nil).Error(err).Args()...)
					return
				}
				if info.Status == kk_etcd_models.PBServer_UnKnown {
					return
				}
				Checker.checkHealth(info)
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

// Convert endpoint to checkInfo
// todo https://github.com/etcd-io/etcd/issues/18520
//
//	checkInfo, ok := endpoint.Metadata.(checkInfo)
func (x *checkInfo) Convert(endpoint endpoints.Endpoint) (err error) {
	bytes, err := json.Marshal(endpoint.Metadata)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, x)
	return err
}
