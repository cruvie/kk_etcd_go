package kVPut

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *Api) CheckInput(stage *kk_stage.Stage) error {

	if x.In.GetKey() == "" {
		return errors.New("key is empty")
	}
	if x.In.GetValue() == "" {
		return errors.New("value is empty")
	}

	return nil
}
