package kVList

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"strings"
)

func (x *api) Handler() (*KVList_Output, error) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	list, err := x.service()
	list.ListKV = slices.DeleteFunc(list.GetListKV(), func(kv *kk_etcd_models.PBKV) bool {
		return strings.HasPrefix(kv.Key, kk_etcd_models.Server)
	})
	return &KVList_Output{KVList: list}, err

}
