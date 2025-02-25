package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_models"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// ParseHeader parse header middleware
func ParseHeader(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)

	var header global_model.RequestHeader
	// bind header
	err := c.ShouldBindHeader(&header)
	if err != nil {
		slog.Error("fail to bind header", "err", err)
		kk_http.WriteCustomResponse(stage, &kk_models.PBResponse{
			Code: http.StatusBadRequest,
			Msg:  "fail to bind header",
			Data: nil,
		})
		c.Abort()
		return
	} else {
		// store header to gin context
		global_model.SetRequestHeader(stage, header)
	}

	c.Next()

}
