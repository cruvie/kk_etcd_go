// Code generated by kk_api_gen. DO NOT EDIT.

package allKVsBackup

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
)

type api struct {
	stage *kk_stage.Stage
	In    *AllKVsBackup_Input
}

func NewApi(stage *kk_stage.Stage) *api {
	return &api{
		stage: stage,
		In:    new(AllKVsBackup_Input),
	}
}

//	 Handler
//		@Tags			backup
//		@ID				AllKVsBackup
//		@Description	all kvs backup
//		@Accept			json,application/x-protobuf
//		@Produce		json,application/x-protobuf
//		@Param			AllKVsBackup_Input	body		AllKVsBackup_Input	true	"AllKVsBackup_Input"
//		@Success		200					{object}	AllKVsBackup_Output
//		@Router			/backup/allKVsBackup [post]
func Handler(c *gin.Context) {
	x := NewApi(kk_global_stage.GetRequestStage(c))
	span := x.stage.StartTrace("allKVsBackup")
	defer span.End()

	if err := x.bindCheck(); err != nil {
		kk_http.WriteResponse(x.stage, err, nil)
		return
	}
	resp, err := x.Handler()
	kk_http.WriteResponse(x.stage, err, resp)
}
