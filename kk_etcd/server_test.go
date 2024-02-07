package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

type server struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (s *server) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status := grpc_health_v1.HealthCheckResponse_SERVING
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
	slog.Info("grpc_rec listening at ", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("failed to serve", "err", err)
	}
}
func TestRegisterGrpcService(t *testing.T) {
	// run TestStartGrpcServer first
	var w sync.WaitGroup
	w.Add(1)
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		return
	}

	//register grpc service
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	_ = RegisterService(&kk_etcd_models.ServiceRegistration{
		Context:    ctx,
		ServerType: kk_etcd_const.ServiceGrpc,
		ServerName: "haha",
		Address:    "127.0.0.1:34844",
		Check: &kk_etcd_models.ServiceCheck{
			Type:    kk_etcd_models.CheckTypeGrpc,
			TTL:     15,
			Timeout: 10,
			GRPC:    "127.0.0.1:34844",
		},
	})
	w.Wait()
}

func TestStartHttpServer(t *testing.T) {
	http.HandleFunc("/Check", func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Check")
		writer.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":8848", nil)
	if err != nil {
		slog.Error("failed to listen", "err", err)
	}
}
func TestRegisterHttpService(t *testing.T) {
	// run TestStartHttpServer first
	var w sync.WaitGroup
	w.Add(1)
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		return
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	_ = RegisterService(&kk_etcd_models.ServiceRegistration{
		Context:    ctx,
		ServerType: kk_etcd_const.ServiceHttp,
		ServerName: "haha",
		Address:    "127.0.0.1:8848",
		Check: &kk_etcd_models.ServiceCheck{
			Type: kk_etcd_models.CheckTypeHttp,
			TTL:  15,
			HTTP: "http://127.0.0.1:8848/Check",
		},
	})
	w.Wait()
	cancelFunc()
}

func TestGetHttpServiceList(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		return
	}
	for i := 0; i < 100; i++ {
		list, _ := ServerList(kk_etcd_const.ServiceHttp)
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}
func TestGetGrpcServiceList(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		return
	}
	for i := 0; i < 100; i++ {
		list, _ := ServerList(kk_etcd_const.ServiceGrpc)
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}

func TestWatchServerList(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		return
	}
	serverListChan := make(chan *kk_etcd_models.PBListServer)
	defer close(serverListChan)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = WatchServerList(ctx, kk_etcd_const.ServiceHttp, serverListChan)
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
