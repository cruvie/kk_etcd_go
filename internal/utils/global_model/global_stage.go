package global_model

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// GlobalStage for kk_etcd client sdk
var GlobalStage *kk_stage.Stage

func InitGlobalStage(stage *kk_stage.Stage) {
	GlobalStage = stage
}

const requestStage = "requestStage"

// SetRequestStage 将当前请求头存储到gin的context中
func SetRequestStage(stage *kk_stage.Stage) {
	stage.GinCtx.Set(requestStage, stage)
}

// GetRequestStage 从gin的context中获取当前stage
func GetRequestStage(c *gin.Context) *kk_stage.Stage {
	stageAny, ok := c.Get(requestStage)
	stage := stageAny.(*kk_stage.Stage)
	if !ok {
		slog.Error("get request stage error", kk_log.NewLog(stage.TraceId).Error(errors.New("request stage not found")).Args()...)
		return nil
	}
	return stage
}
