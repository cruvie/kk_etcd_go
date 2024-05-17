package kk_etcd_models

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *KVPutParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}
	if x.GetKey() == "" {
		return errors.New("key is empty")
	}
	if x.GetValue() == "" {
		return errors.New("value is empty")
	}

	return nil
}
func (x *KVGetParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *KVDelParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}

func (x *KVListParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
