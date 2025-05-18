package serverList

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/server/api_def"
)

func (x *api) Handler() (*api_def.ServerList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	server := &server_hub.SerServer{}
	serverList, err := server.ServerList(x.In.GetServerType(), x.In.GetServerName())
	return &api_def.ServerList_Output{
		ServerList: serverList,
	}, err
}
