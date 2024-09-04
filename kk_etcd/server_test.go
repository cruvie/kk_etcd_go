package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"log/slog"
	"net"
	"net/http"
	"testing"
	"time"
)

type server struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (s *server) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status := grpc_health_v1.HealthCheckResponse_SERVING
	slog.Info("receive check", "in", in)
	return &grpc_health_v1.HealthCheckResponse{Status: status}, nil
}

func (s *server) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	slog.Info("Watch")
	return nil
}

func TestStartGrpcServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":34844")
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			slog.Error("failed to close listener", "err", err)
		}
	}(listener)
	if err != nil {
		slog.Error("failed to listen", "err", err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	grpc_health_v1.RegisterHealthServer(grpcServer, &server{})
	slog.Info("grpc_rec listening at", kk_log.NewLog(nil).Any("addr", listener.Addr()).Args()...)
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("failed to serve", "err", err)
	}
}
func TestRegisterGrpcService(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()

	//register grpc service
	err := RegisterService(&kk_etcd_models.ServerRegistration{
		ServerType: kk_etcd_models.Grpc,
		ServerName: "haha_grpc",
		ServerAddr: "127.0.0.1:34844",
		Metadata:   "meta",
		CheckConfig: kk_etcd_models.CheckConfig{
			Type:     kk_etcd_models.Grpc,
			Timeout:  10 * time.Second,
			Interval: 5 * time.Second,
			Addr:     "127.0.0.1:34844",
		},
	})
	if err != nil {
		slog.Error("failed to register service", "err", err)
	}
}
func TestGetGrpcServiceList(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	for {
		list, err := ServerList(&kk_etcd_models.ServerListParam{ServerType: kk_etcd_models.Grpc.String(), ServerName: "haha_grpc"})
		if err != nil {
			slog.Error("failed to list", "err", err)
		}
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}

func TestStartHttpServer(t *testing.T) {
	http.HandleFunc(kk_etcd_models.HealthCheckPath, func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Check")
		writer.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":8848", nil)
	if err != nil {
		slog.Error("failed to listen", "err", err)
	}
}
func TestRegisterHttpService(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	err := RegisterService(&kk_etcd_models.ServerRegistration{
		ServerType: kk_etcd_models.Http,
		ServerName: "haha_http",
		ServerAddr: "127.0.0.1:8848",
		CheckConfig: kk_etcd_models.CheckConfig{
			Type:     kk_etcd_models.Http,
			Timeout:  10 * time.Second,
			Interval: 5 * time.Second,
			Addr:     "http://127.0.0.1:8848" + kk_etcd_models.HealthCheckPath,
		},
	})
	if err != nil {
		slog.Error("failed to register service", "err", err)
	}
}

func TestGetHttpServiceList(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	for {
		list, err := ServerList(&kk_etcd_models.ServerListParam{ServerType: kk_etcd_models.Http.String(), ServerName: "haha_http"})
		if err != nil {
			slog.Error("failed to list", "err", err)
		}
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}

func TestWatchServerList(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	defer close(serverListChan)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := WatchServerList(ctx, kk_etcd_models.Http, "haha_http", serverListChan)
	if err != nil {
		slog.Error("WatchServerList failed", "err", err)
		return
	}

	for {
		slog.Info("watching for server list change")
		select {
		case <-ctx.Done():
			slog.Info("ctx done")
			return
		case serverList := <-serverListChan:
			slog.Info("serverListChan", "serverList", serverList)
		}
	}
}
