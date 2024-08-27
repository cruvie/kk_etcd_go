package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/gin-gonic/gin"
)

// RequestInterceptor the first middleware to intercept request
func RequestInterceptor(stage *kk_stage.Stage) func(c *gin.Context) {
	name := "RequestInterceptor"
	return func(c *gin.Context) {

		//slog.Info("url", c.Request.URL.Path)
		stage := kk_stage.NewStage(c.Request.Context(), name, stage.DebugMode).SetGinCtx(c)
		global_stage.SetRequestStage(stage)

		c.Next()
	}
}
