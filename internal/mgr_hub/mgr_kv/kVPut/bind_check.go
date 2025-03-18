package kVPut

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *api) bindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x.in)
	if err != nil {
		return err
	}

	if x.in.GetKey() == "" {
		return errors.New("key is empty")
	}
	if x.in.GetValue() == "" {
		return errors.New("value is empty")
	}

	return nil
}
