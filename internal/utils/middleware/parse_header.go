package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_response"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"

	"github.com/gin-gonic/gin"
	"log/slog"
)

// ParseHeader parse header middleware
func ParseHeader(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)

	var header global_model.RequestHeader
	// bind header
	err := c.ShouldBindHeader(&header)
	if err != nil {
		slog.Error("fail to bind header", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		c.Abort()
		return
	} else {
		// store header to gin context
		global_model.SetRequestHeader(c, header)
	}

	c.Next()

}
