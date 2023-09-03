package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// ParseHeader parse header middleware
func ParseHeader(c *gin.Context) {

	var header global_model.RequestHeader
	// bind header
	err := c.ShouldBindHeader(&header)
	if err != nil {
		slog.Info("fail to bind header", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		c.Abort()
		return
	} else {
		// store header to gin context
		global_model.SetRequestHeader(c, header)
	}

	c.Next()

}
