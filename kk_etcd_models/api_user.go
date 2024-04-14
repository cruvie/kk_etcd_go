package kk_etcd_models

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *LoginParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *LogoutParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *UserAddParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *UserDeleteParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *GetUserParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *MyInfoParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *UserListParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *UserGrantRoleParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	return nil
}
