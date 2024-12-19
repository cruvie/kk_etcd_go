package service

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/etcd_ai/rag_server"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type SerAI struct{}

func (SerAI) Query(stage *kk_stage.Stage, param *kk_etcd_models.QueryParam) (string, error) {
	//contents, err := vector_store.SearchDocuments(param.GetQuestion())
	//if err != nil {
	//	return "", err
	//}
	answer, err := rag_server.Query(param.GetQuestion(), nil)
	return answer, err
}
