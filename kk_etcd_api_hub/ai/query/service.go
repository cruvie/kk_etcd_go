package query

import "github.com/cruvie/kk_etcd_go/internal/etcd_ai/rag_server"

func (x *api) service() (string, error) {
	span := x.stage.StartTrace("service")
	defer span.End()
	//contents, err := vector_store.SearchDocuments(x.In.GetQuestion())
	//if err != nil {
	//	return "", err
	//}
	answer, err := rag_server.Query(x.In.GetQuestion(), nil)
	return answer, err
}
