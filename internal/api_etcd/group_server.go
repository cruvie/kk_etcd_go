package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_reflect"
	"github.com/cruvie/kk_etcd_go/internal/internal_handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hServer internal_handler.HServer

// serverList
//
//	@Tags			server
//	@Description	list server
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			ServerListParam	body		kk_etcd_models.ServerListParam	true	"ServerListParam"
//	@Success		200				{object}	kk_etcd_models.ServerListResponse
//	@Router			/server/serverList [post]
func serverList(c *gin.Context) {
	stage := global_stage.GetRequestStage(c)
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.ServerListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hServer.ServerList(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}
