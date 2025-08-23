package api_etcd

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"gitee.com/cruvie/kk_go_kit/kk_file"
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"gitee.com/cruvie/kk_go_kit/kk_grpc/interceptor"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	apiImplBackup "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_impl"
	apiImplKv "github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_impl"
	apiImplRole "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_impl"
	apiImplService "github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_impl"
	apiImplUser "github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_impl"
	"github.com/cruvie/kk_etcd_go/internal/utils/middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	// grpcurl -plaintext localhost:2333 list
	// grpcurl -plaintext localhost:2333 describe api_def.AI
	kk_grpc.RegisterKKHealthCheckServer(grpcServer)
	apiImplRole.RegisterServer(grpcServer)
	apiImplUser.RegisterServer(grpcServer)
	apiImplService.RegisterServer(grpcServer)
	apiImplKv.RegisterServer(grpcServer)
	apiImplBackup.RegisterServer(grpcServer)

	// 使用 grpcweb 包装 gRPC 服务端
	grpcWebServer := grpcweb.WrapServer(grpcServer)

	// 创建 HTTP 服务器，用于处理 gRPC-Web 请求
	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 允许跨域请求（根据需要配置）
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Grpc-Web")

			// 如果是预检请求，直接返回
			if r.Method == http.MethodOptions {
				return
			}

			// 将请求交给 grpcWebServer 处理
			if grpcWebServer.IsGrpcWebRequest(r) {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				http.Error(w, "Not a valid gRPC-Web request", http.StatusBadRequest)
			}
		}),
	}

	run := func() {
		go func() {
			if err := grpcServer.Serve(listener); err != nil {
				panic(err)
			}
		}()
		if err := httpServer.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
	done := func(quitCh <-chan struct{}) {
		<-quitCh
		// 包含了listener.Close()
		grpcServer.GracefulStop()
		err := httpServer.Shutdown(stage.Ctx)
		if err != nil {
			slog.Error("httpServer.Shutdown", "err", err)
		}
	}
	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}
