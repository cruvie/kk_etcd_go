package query

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/etcd_ai/rag_server"
)

func (x *Api) Service(stage *kk_stage.Stage) (string, error) {
	span := stage.StartTrace("service")
	defer span.End()
	//contents, err := vector_store.SearchDocuments(x.In.GetQuestion())
	//if err != nil {
	//	return "", err
	//}
	answer, err := rag_server.Query(x.In.GetQuestion(), nil)
	return answer, err
}
