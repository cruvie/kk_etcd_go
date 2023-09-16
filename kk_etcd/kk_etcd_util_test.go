package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
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

type config struct {
	ServerAddr string `yaml:"ServerAddr"`
	Postgres   struct {
		Dsn string `yaml:"Dsn"`
	} `yaml:"Postgres"`
	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
	} `yaml:"Redis"`
	MinIO struct {
		AccessEndpoint string `yaml:"AccessEndpoint"`
	} `yaml:"MinIO"`
}

var GlobalConfig config

func TestGetConfig(t *testing.T) {

	var (
		endpoints = []string{"http://127.0.0.1:2379"} //http://etcd:2379  http://127.0.0.1:2379
		configKey = "my_config"

		userName = "admin"
		password = "admin"
	)

	InitEtcd(endpoints, userName, password)
	GetConfig(configKey, &GlobalConfig)

}

type server struct {
}

func (s *server) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	status := grpc_health_v1.HealthCheckResponse_SERVING
	// 检查gRPC服务的状态，并根据情况设置status
	return &grpc_health_v1.HealthCheckResponse{Status: status}, nil
}

func (s *server) Watch(request *grpc_health_v1.HealthCheckRequest, watchServer grpc_health_v1.Health_WatchServer) error {
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
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	//register grpc service
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	_ = RegisterService(&kk_etcd_models.ServiceRegistration{
		Context:    ctx,
		ServerType: key_prefix.ServiceGrpc,
		ServerName: "haha",
		Address:    "127.0.0.1:34844",
		Check: &kk_etcd_models.ServiceCheck{
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
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	_ = RegisterService(&kk_etcd_models.ServiceRegistration{
		Context:    ctx,
		ServerType: key_prefix.ServiceHttp,
		ServerName: "haha",
		Address:    "127.0.0.1:8848",
		Check: &kk_etcd_models.ServiceCheck{
			TTL:  15,
			HTTP: "http://127.0.0.1:8848/Check",
		},
	})
	w.Wait()
	cancelFunc()
}

func TestGetHttpServiceList(t *testing.T) {
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	for i := 0; i < 100; i++ {
		list, _ := ServerList(key_prefix.ServiceHttp)
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}
func TestGetGrpcServiceList(t *testing.T) {
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	for i := 0; i < 100; i++ {
		list, _ := ServerList(key_prefix.ServiceGrpc)
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}
