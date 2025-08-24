package middleware

import (
	"context"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"gitee.com/cruvie/kk_go_kit/kk_grpc/interceptor"
	"gitee.com/cruvie/kk_go_kit/kk_protobuf"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// EtcdClient set etcd client for current request
func etcdClient(userName, password string) (*clientv3.Client, error) {
	// todo use https://github.com/etcd-io/etcd/pull/16803
	// set verify jwt directly
	// slog.Info("url", c.Request.URL.Path)
	cfg := clientv3.Config{
		Endpoints:   config.Config.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
		Username:    userName,
		Password:    password,
	}
	return clientv3.New(cfg)
}

func UnaryEtcdClient() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		stage := kk_grpc.GetGrpcStage(ctx)

		list, err := interceptor.GetInterceptorTypeListFromStage(stage)
		if err != nil {
			return nil, err
		}
		if !kk_protobuf.InterceptorType_INTERCEPTOR_TYPE_AUTH.IsIn(list...) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "kk_grpc missing metadata")
		}

		names := md.Get("UserName")
		if len(names) == 0 {
			return nil, status.Error(codes.Unauthenticated, "kk_grpc missing names")
		}
		pwds := md.Get("Password")
		if len(pwds) == 0 {
			return nil, status.Error(codes.Unauthenticated, "kk_grpc missing pwds")
		}

		client, err := etcdClient(names[0], pwds[0])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "kk_grpc missing client")
		}

		global_model.SetClient(stage, client)

		user, err := util_user.GetUser(stage, names[0])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "kk_grpc missing user")
		}
		// store user to gin context
		global_model.SetLoginUser(stage, user)

		resp, err := handler(ctx, req)

		{
			// end of a request
			global_model.CloseClient(stage)
		}
		return resp, err
	}
}
