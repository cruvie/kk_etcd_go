package global_stage

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/gin-gonic/gin"
	"log/slog"
)

const requestStage = "requestStage"

// SetRequestStage 将当前请求头存储到gin的context中
func SetRequestStage(stage *kk_stage.Stage) {
	stage.GinCtx.Set(requestStage, stage)
}

// GetRequestStage 从gin的context中获取当前stage
func GetRequestStage(c *gin.Context) *kk_stage.Stage {
	stageAny, ok := c.Get(requestStage)
	stage := stageAny.(*kk_stage.Stage)
	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: stage.TraceId})
	if !ok {
		slog.Error("get request stage error", newLog.Error(errors.New("request stage not found")).Args()...)
		return nil
	}
	return stage
}
