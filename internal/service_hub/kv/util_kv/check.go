package util_kv

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func CheckKeyExist(stage *kk_stage.Stage, key string) error {
	_, err := GetKV(global_model.GetClient(stage), key)
	if err != nil {
		return err
	} else {
		return utils.ErrKeyAlreadyExists
	}
}
