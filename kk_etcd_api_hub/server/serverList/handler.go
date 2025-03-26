package serverList

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

// ServerList
// serverName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service lists
func (x *api) Handler() (*ServerList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	serverList, err := server_hub.SerServer{}.ServerList(global_model.GetClient(x.stage),
		kk_etcd_models.ServerType(x.In.GetServerType()))
	return &ServerList_Output{
		ServerList: serverList,
	}, err
}
