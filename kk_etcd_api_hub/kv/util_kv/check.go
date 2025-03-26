package util_kv

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
)

func CheckKeyExist(stage *kk_stage.Stage, key string) error {
	_, err := GetKV(stage, key)
	if err != nil {
		return err
	} else {
		return kk_etcd_error.ErrKeyAlreadyExists
	}
}
