// Code generated by kk_api_gen. DO NOT EDIT.

package kVGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/api_def"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	In    *api_def.KVGet_Input
}

func NewApi(stage *kk_stage.Stage) *api {
	return &api{
		stage: stage,
		In:    new(api_def.KVGet_Input),
	}
}

//  Handler
//	@Tags			kv
//	@ID				KVGet
//	@Description	get kv
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			api_def.KVGet_Input	body		api_def.KVGet_Input	true	"KVGet_Input"
//	@Success		200					{object}	api_def.KVGet_Output
//	@Router			/kv/kVGet [post]
func Handler(c *gin.Context) {
		x := NewApi(kk_global_stage.GetRequestStage(c))
		span := x.stage.StartTrace("kVGet")
		defer span.End()
	
		if err := x.bindCheck(); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		resp, err := x.Handler()
		kk_http.WriteResponse(x.stage, err, resp)
}