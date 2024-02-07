package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"

	"github.com/gin-gonic/gin"
)

// ServerList
//
//	@Description	Get service list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/ServerList [post]
func ServerList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var prefix kk_pb_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &prefix); err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res, list, _ := service.ServerList(stage, prefix.Value)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, list))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}
