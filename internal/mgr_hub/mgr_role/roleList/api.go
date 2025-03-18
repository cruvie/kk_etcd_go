// Code generated by kk_api_gen. DO NOT EDIT.

package roleList

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	in    *RoleList_Input
}

func newApi(c *gin.Context) *api {
	return &api{
		stage: kk_global_stage.GetRequestStage(c),
		in:    new(RoleList_Input),
	}
}

//  Handler
//	@Tags			role
//	@ID				RoleList
//	@Description	list role
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			RoleList_Input	body		RoleList_Input	true	"RoleList_Input"
//	@Success		200				{object}	RoleList_Output
//	@Router			/role/roleList [post]
func Handler(c *gin.Context) {
		x := newApi(c)
		span := x.stage.StartTrace("roleList")
		defer span.End()
	
		if err := x.bindCheck(x.stage); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		err, resp := x.handler()
		kk_http.WriteResponse(x.stage, err, resp)
}