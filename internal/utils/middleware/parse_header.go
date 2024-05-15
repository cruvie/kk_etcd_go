package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"

	"github.com/gin-gonic/gin"
	"log/slog"
)

// ParseHeader parse header middleware
func ParseHeader(c *gin.Context) {
	stage := global_model.GetRequestStage(c)

	var header global_model.RequestHeader
	// bind header
	err := c.ShouldBindHeader(&header)
	if err != nil {
		slog.Error("fail to bind header", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		c.Abort()
		return
	} else {
		// store header to gin context
		global_model.SetRequestHeader(stage, header)
	}

	c.Next()

}
