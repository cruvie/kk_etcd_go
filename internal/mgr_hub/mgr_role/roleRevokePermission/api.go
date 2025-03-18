// Code generated by kk_api_gen. DO NOT EDIT.

package roleRevokePermission

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	in    *RoleRevokePermission_Input
}

func newApi(c *gin.Context) *api {
	return &api{
		stage: kk_global_stage.GetRequestStage(c),
		in:    new(RoleRevokePermission_Input),
	}
}

//  Handler
//	@Tags			role
//	@ID				RoleRevokePermission
//	@Description	revoke permission
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			RoleRevokePermission_Input	body		RoleRevokePermission_Input	true	"RoleRevokePermission_Input"
//	@Success		200							{object}	RoleRevokePermission_Output
//	@Router			/role/roleRevokePermission [post]
func Handler(c *gin.Context) {
		x := newApi(c)
		span := x.stage.StartTrace("roleRevokePermission")
		defer span.End()
	
		if err := x.bindCheck(x.stage); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		err, resp := x.handler()
		kk_http.WriteResponse(x.stage, err, resp)
}