package kVPut

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}

	if x.In.GetKey() == "" {
		return errors.New("key is empty")
	}
	if x.In.GetValue() == "" {
		return errors.New("value is empty")
	}

	return nil
}
