package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_base_proto_type"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_response"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"

	"github.com/gin-gonic/gin"
)

// ServerList
//
//	@Description	Get service list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/ServerList [post]
func ServerList(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	var prefix kk_base_proto_type.PBString
	if err := kku_http.ReadProtoBuf(stage, &prefix); err != nil {
		kku_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res, list, _ := service.ServerList(stage, prefix.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, list))
		return
	}
	kku_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}
