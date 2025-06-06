// Code generated by kk_api_gen. DO NOT EDIT.

package myInfo

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/api_def"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	In    *api_def.MyInfo_Input
}

func NewApi(stage *kk_stage.Stage) *api {
	return &api{
		stage: stage,
		In:    new(api_def.MyInfo_Input),
	}
}

//  Handler
//	@Tags			user
//	@ID				MyInfo
//	@Description	get my info
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			api_def.MyInfo_Input	body		api_def.MyInfo_Input	true	"MyInfo_Input"
//	@Success		200						{object}	api_def.MyInfo_Output
//	@Router			/user/myInfo [post]
func Handler(c *gin.Context) {
		x := NewApi(kk_global_stage.GetRequestStage(c))
		span := x.stage.StartTrace("myInfo")
		defer span.End()
	
		if err := x.bindCheck(); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		resp, err := x.Handler()
		kk_http.WriteResponse(x.stage, err, resp)
}