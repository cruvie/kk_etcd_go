package kVList

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"slices"
	"strings"
)

func (x *api) handler() (error, *KVList_Output) {
	span := x.stage.StartTrace("handler")
	defer span.End()

	err, list := x.service()
	list.ListKV = slices.DeleteFunc(list.GetListKV(), func(kv *kk_etcd_models.PBKV) bool {
		return strings.HasPrefix(kv.Key, kk_etcd_models.Server)
	})
	return err, &KVList_Output{KVList: list}

}
