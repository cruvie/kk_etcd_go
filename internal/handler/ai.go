package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HAI struct{}

var serAI service.SerAI

func (HAI) Query(stage *kk_stage.Stage, param *kk_etcd_models.QueryParam) (error, *kk_etcd_models.QueryResponse) {
	span := stage.StartTrace("Query")
	defer span.End()
	answer, err := serAI.Query(stage, param)
	if err != nil {
		return err, nil
	}
	return nil, &kk_etcd_models.QueryResponse{
		Answer: answer,
	}
}
