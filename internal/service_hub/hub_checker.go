package service_hub

import (
	"context"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
	"net/http"
	"time"
)

var runningCheck = make(map[string] /*UniqueKey*/ *checkT)

type checkT struct {
	ctx          context.Context
	cancel       context.CancelFunc
	registration *kk_etcd_models.PBServiceRegistration
}

func startNewCheck(registration *kk_etcd_models.PBServiceRegistration) {
	ctx, cancel := context.WithCancel(context.Background())
	checker := &checkT{
		ctx:          ctx,
		cancel:       cancel,
		registration: registration,
	}
	stopCheck(registration.UniqueKey())
	runningCheck[registration.UniqueKey()] = checker
	go checker.runCheck()
}

func stopCheck(uniqueKey string) {
	t, ok := runningCheck[uniqueKey]
	if !ok {
		return
	}
	t.cancel()
	delete(runningCheck, uniqueKey)
}
func (x *checkT) update(err error) {
	if err == nil {
		hub.putToAliveHub(x.registration)
		hub.delFromDeadHub(x.registration)
	}
}

func (x *checkT) runCheck() {
	ticker := time.NewTicker(x.registration.CheckConfig.Interval.AsDuration())
	switch x.registration.CheckConfig.Type {
	case kk_etcd_models.PBServiceType_Grpc:
		//chack instance first
		err := checkGrpc(x.registration.CheckConfig)
		x.update(err)
		for {
			select {
			case <-x.ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				err := checkGrpc(x.registration.CheckConfig)
				x.update(err)
			}
		}
	case kk_etcd_models.PBServiceType_Http:
		err := checkHttp(x.registration.CheckConfig)
		x.update(err)
		for {
			select {
			case <-x.ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				err := checkHttp(x.registration.CheckConfig)
				x.update(err)
			}
		}
	}
}

func checkGrpc(checkConfig *kk_etcd_models.PBServiceRegistration_PBCheckConfig) (err error) {
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
			slog.Error(msg, kk_stage.NewLog(nil).Error(err).Args()...)
		}
	}(conn)
	healthClient := grpc_health_v1.NewHealthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), checkConfig.Timeout.AsDuration())
	defer cancel()
	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return fmt.Errorf("failed to checkT grpc health %w", err)
	}
	status := resp.GetStatus()
	if status == grpc_health_v1.HealthCheckResponse_SERVING {
		return nil
	} else {
		return fmt.Errorf("grpc status code is not ok, status:%d", status)
	}
}

func checkHttp(checkConfig *kk_etcd_models.PBServiceRegistration_PBCheckConfig) (err error) {
	var httpClient = &http.Client{
		Timeout: checkConfig.Timeout.AsDuration(),
	}
	defer httpClient.CloseIdleConnections()
	// send http request to addr
	req, err := http.NewRequest(http.MethodGet, checkConfig.Addr, nil)
	if err != nil {
		return fmt.Errorf("failed to new http request %w", err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to checkT http health %w", err)
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("http status code is not ok, code:%d", resp.StatusCode)
	}
}
