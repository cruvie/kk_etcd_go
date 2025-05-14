package kVPut

import "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/api_def"

func (x *api) Handler() (*api_def.KVPut_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err := x.service()
	return &api_def.KVPut_Output{}, err
}
