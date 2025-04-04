// Code generated by kk_api_gen. DO NOT EDIT.

package kVGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	In    *KVGet_Input
}

func NewApi(stage *kk_stage.Stage) *api {
	return &api{
		stage: stage,
		In:    new(KVGet_Input),
	}
}

//  Handler
//	@Tags			kv
//	@ID				KVGet
//	@Description	get kv
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			KVGet_Input	body		KVGet_Input	true	"KVGet_Input"
//	@Success		200			{object}	KVGet_Output
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