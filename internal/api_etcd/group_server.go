package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hServer handler.HServer

// serverList
//
//	@Tags			server
//	@Description	list server
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			ServerListParam	body		kk_etcd_models.ServerListParam	true	"ServerListParam"
//	@Success		200				{object}	kk_etcd_models.ServerListResponse
//	@Router			/server/serverList [post]
func serverList(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("serverList")
	defer span.End()
	var param kk_etcd_models.ServerListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hServer.ServerList(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}

// deregisterServer
//
//	@Tags			server
//	@Description	deregister server
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			DeregisterServerParam	body		kk_etcd_models.DeregisterServerParam	true	"DeregisterServerParam"
//	@Success		200						{object}	kk_etcd_models.DeregisterServerResponse
//	@Router			/server/deregisterServer [post]
func deregisterServer(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("deregisterServer")
	defer span.End()
	var param kk_etcd_models.DeregisterServerParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hServer.DeregisterServer(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}
