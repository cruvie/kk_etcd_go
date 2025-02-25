package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hAI handler.HAI

// query
//
//	@Tags			ai
//	@ID				Query
//	@Description	query ai
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			QueryParam	body		kk_etcd_models.QueryParam	true	"QueryParam"
//	@Success		200			{object}	kk_etcd_models.QueryResponse
//	@Router			/ai/query [post]
func query(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("query")
	defer span.End()

	var param kk_etcd_models.QueryParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hAI.Query(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}
