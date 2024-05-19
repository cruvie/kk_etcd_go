package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hServer handler.HServer

// serverList
//
//	@Description	serverList
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			ServerListParam	body		kk_etcd_models.ServerListParam	true	"ServerListParam"
//	@Success		200				{object}	kk_etcd_models.ServerListResponse
//	@Router			/server/serverList [post]
func serverList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
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
