package kk_etcd_test

import (
	"context"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_net"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/types/known/durationpb"
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

func serviceList(req *api_def.ServiceList_Input) (*api_def.ServiceList_Output, error) {
	header := http.Header{}
	header.Add("Content-Type", "application/x-protobuf")
	header.Add("Accept", "application/x-protobuf")
	header.Add("Username", "root")
	header.Add("Password", "root")
	out := &api_def.ServiceList_Output{}
	err := kk_http.SendPBRequest(5*time.Second,
		http.MethodPost, "http://127.0.0.1:2333/service/serviceList",
		header, req, out)

	return out, err
}

func TestServiceList(t *testing.T) {
	req := &api_def.ServiceList_Input{}
	out, err := serviceList(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Response: %+v\n", out)
}

func cleanup(t *testing.T) {
	client := global_model.GetClient(internal_client.GlobalStage)
	//kk_service/internal_service_status/kk_service/grpc/ss_msg/192.168.124.118:58754
	resp, err := client.Get(context.Background(), kk_etcd_models.ServiceKey, clientv3.WithPrefix())
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	_, err = global_model.GetClient(internal_client.GlobalStage).
		Delete(context.Background(),
			kk_etcd_models.ServiceKey,
			clientv3.WithPrefix(),
		)
	if err != nil {
		t.Error(err)
	}
}

func TestCleanUp(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	cleanup(t)
}
func runGrpcServiceWithPort(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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

func registerGrpcService(port int) {
	localIP, err := kk_net.GetLocalIP()
	if err != nil {
		panic(err)
	}
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()

	//register grpc service
	err = kk_etcd.RegisterService(&kk_etcd_models.PBServiceRegistration{
		ServiceType: kk_etcd_models.PBServiceType_Grpc,
		ServiceName: fmt.Sprintf("test_grpc_%d", port),
		ServiceAddr: fmt.Sprintf("%s:%d", localIP, port),
		CheckConfig: &kk_etcd_models.PBServiceRegistration_PBCheckConfig{
			Type:     kk_etcd_models.PBServiceType_Grpc,
			Timeout:  durationpb.New(5 * time.Second),
			Interval: durationpb.New(10 * time.Second),
			Addr:     fmt.Sprintf("%s:%d", localIP, port),
		},
	})
	if err != nil {
		slog.Error("failed to register service", "err", err)
	}
}

func TestStartGrpcServer(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	port := 34000
	for i := 0; i < 30; i++ {
		go runGrpcServiceWithPort(port)
		registerGrpcService(port)
		port++
	}
	select {}
}

func TestGetGrpcServiceList(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			t.Error(err)
		}
	}()
	for {
		list, err := serviceList(&api_def.ServiceList_Input{ServiceType: kk_etcd_models.PBServiceType_Grpc, ServiceName: "haha_grpc"})
		if err != nil {
			slog.Error("failed to list", "err", err)
		}
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}

func TestGetGrpcClient(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	conn, client, err := kk_etcd.GetGrpcClient[grpc_health_v1.HealthClient]("test_grpc_34004",
		grpc_health_v1.NewHealthClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close() //nolint
	response, err := client.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{
		Service: "empty kk",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println(response)
}

func runHttpServiceWithPort(port int) {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc(kk_etcd_models.HealthCheckPath, func(writer http.ResponseWriter, request *http.Request) {
		slog.Info("Check")
		writer.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), serveMux)
	if err != nil {
		slog.Error("failed to listen", "err", err)
	}
}

func registerHttpService(port int) {
	localIP, err := kk_net.GetLocalIP()
	if err != nil {
		panic(err)
	}

	err = kk_etcd.RegisterService(&kk_etcd_models.PBServiceRegistration{
		ServiceType: kk_etcd_models.PBServiceType_Http,
		ServiceName: fmt.Sprintf("test_http_%d", port),
		ServiceAddr: fmt.Sprintf("%s:%d", localIP, port),
		CheckConfig: &kk_etcd_models.PBServiceRegistration_PBCheckConfig{
			Type:     kk_etcd_models.PBServiceType_Http,
			Timeout:  durationpb.New(5 * time.Second),
			Interval: durationpb.New(10 * time.Second),
			Addr:     fmt.Sprintf("http://%s:%d%s", localIP, port, kk_etcd_models.HealthCheckPath),
		},
	})
	if err != nil {
		slog.Error("failed to register service", "err", err)
	}
}

func TestStartHttpServer(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	port := 8840
	for i := 0; i < 30; i++ {
		go runHttpServiceWithPort(port)
		registerHttpService(port)
		port++
	}
	select {}
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
		list, err := serviceList(&api_def.ServiceList_Input{ServiceType: kk_etcd_models.PBServiceType_Http, ServiceName: "haha_http"})
		if err != nil {
			slog.Error("failed to list", "err", err)
		}
		slog.Info("list", "list", list)
		time.Sleep(time.Second * 5)
	}
}

//func TestWatchServiceList(t *testing.T) {
//	closeFunc := initTestEnv()
//	defer func() {
//		err := closeFunc()
//		if err != nil {
//			log.Println(err)
//		}
//	}()
//	serverListChan := make(chan *kk_etcd_models.PBListService)
//	defer close(serverListChan)
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	err := WatchServiceList(ctx, kk_etcd_models.PBServiceType_Http, "haha_http", serverListChan)
//	if err != nil {
//		slog.Error("WatchServiceList failed", "err", err)
//		return
//	}
//
//	for {
//		slog.Info("watching for server list change")
//		select {
//		case <-ctx.Done():
//			slog.Info("ctx done")
//			return
//		case serverList := <-serverListChan:
//			slog.Info("serverListChan", "serverList", serverList)
//		}
//	}
//}
