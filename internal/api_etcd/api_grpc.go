package api_etcd

import (
	"fmt"
	"net"

	"gitee.com/cruvie/kk_go_kit/kk_file"
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"gitee.com/cruvie/kk_go_kit/kk_grpc/interceptor"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	apiImplAI "github.com/cruvie/kk_etcd_go/internal/service_hub/ai/api_impl"
	apiImplBackup "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_impl"
	apiImplKv "github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_impl"
	apiImplRole "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_impl"
	apiImplService "github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_impl"
	apiImplUser "github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_impl"
	"github.com/cruvie/kk_etcd_go/internal/utils/middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc/reflection"

	"github.com/cruvie/kk_etcd_go/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer(stage *kk_stage.Stage) *kk_server.KKRunServer {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.MaxRecvMsgSize(kk_file.MB.Int()*10),
		grpc.MaxSendMsgSize(kk_file.MB.Int()*10),
		grpc.ChainUnaryInterceptor(
			middleware.UnaryLogging(),
			interceptor.UnaryStage(kk_grpc.GFileDescHub),
			middleware.UnaryEtcdClient(),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(middleware.PanicRecovery)),
		),
	)
	reflection.Register(grpcServer)
	//grpcurl -plaintext localhost:2333 list
	//grpcurl -plaintext localhost:2333 describe api_def.AI
	kk_grpc.RegisterKKHealthCheckServer(grpcServer)
	apiImplAI.RegisterServer(grpcServer)
	apiImplRole.RegisterServer(grpcServer)
	apiImplUser.RegisterServer(grpcServer)
	apiImplService.RegisterServer(grpcServer)
	apiImplKv.RegisterServer(grpcServer)
	apiImplBackup.RegisterServer(grpcServer)

	run := func() {
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}
	done := func(quitCh <-chan struct{}) {
		<-quitCh
		// 包含了listener.Close()
		grpcServer.GracefulStop()
	}
	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}
