// Code generated by kk_api_gen. DO NOT EDIT.

package login

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	In    *Login_Input
}

func NewApi(stage *kk_stage.Stage) *api {
	return &api{
		stage: stage,
		In:    new(Login_Input),
	}
}

//  Handler
//	@Tags			user
//	@ID				Login
//	@Description	login
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			Login_Input	body		Login_Input	true	"Login_Input"
//	@Success		200			{object}	Login_Output
//	@Router			/user/login [post]
func Handler(c *gin.Context) {
		x := NewApi(kk_global_stage.GetRequestStage(c))
		span := x.stage.StartTrace("login")
		defer span.End()
	
		if err := x.bindCheck(); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		resp, err := x.Handler()
		kk_http.WriteResponse(x.stage, err, resp)
}