package kVList

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"strings"
)

func (x *api) Handler() (*api_def.KVList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	list, err := x.service()
	if err != nil {
		return nil, err
	}
	list.ListKV = slices.DeleteFunc(list.GetListKV(), func(kv *kk_etcd_models.PBKV) bool {
		return strings.HasPrefix(kv.Key, kk_etcd_models.ServiceKey)
	})
	return &api_def.KVList_Output{KVList: list}, err

}
