package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// ParseHeader parse header middleware
func ParseHeader(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())

	var header global_model.RequestHeader
	// bind header
	err := c.ShouldBindHeader(&header)
	if err != nil {
		slog.Error("fail to bind header", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		c.Abort()
		return
	} else {
		// store header to gin context
		global_model.SetRequestHeader(c, header)
	}

	c.Next()

}
