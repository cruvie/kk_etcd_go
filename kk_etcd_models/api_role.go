package kk_etcd_models

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *RoleAddParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *RoleDeleteParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *RoleListParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *RoleGetParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *RoleGrantPermissionParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
