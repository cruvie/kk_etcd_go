package middleware

import (
	"context"
	"log/slog"

	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

// interceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		config.LogCfg.Logger.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func logTraceID(ctx context.Context) logging.Fields {
	traceId := kk_grpc.GetTraceIdFromGrpcCtx(ctx)
	return logging.Fields{kk_stage.TraceIdKey, traceId}
}

func UnaryLogging() grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(interceptorLogger(), logging.WithFieldsFromContext(logTraceID))
}

func StreamLogging() grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(interceptorLogger(), logging.WithFieldsFromContext(logTraceID))
}
