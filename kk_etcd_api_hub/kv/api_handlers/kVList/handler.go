package kVList

import (
	"slices"
	"strings"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*api_def.KVList_Output, error) {

	list, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	list.ListKV = slices.DeleteFunc(list.GetListKV(), func(kv *kk_etcd_models.PBKV) bool {
		return strings.HasPrefix(kv.Key, kk_etcd_models.ServiceKey)
	})
	return &api_def.KVList_Output{KVList: list}, err

}
