// Code generated by kk_api_gen. DO NOT EDIT.

package snapshotInfo

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	in    *SnapshotInfo_Input
}

func newApi(c *gin.Context) *api {
	return &api{
		stage: kk_global_stage.GetRequestStage(c),
		in:    new(SnapshotInfo_Input),
	}
}

//  Handler
//	@Tags			backup
//	@ID				SnapshotInfo
//	@Description	snapshot info
//	@Accept			json,application/x-protobuf
//	@Produce		json,application/x-protobuf
//	@Param			SnapshotInfo_Input	body		SnapshotInfo_Input	true	"SnapshotInfo_Input"
//	@Success		200					{object}	SnapshotInfo_Output
//	@Router			/backup/snapshotInfo [post]
func Handler(c *gin.Context) {
		x := newApi(c)
		span := x.stage.StartTrace("snapshotInfo")
		defer span.End()
	
		if err := x.bindCheck(x.stage); err != nil {
			kk_http.WriteResponse(x.stage, err, nil)
			return
		}
		err, resp := x.handler()
		kk_http.WriteResponse(x.stage, err, resp)
}