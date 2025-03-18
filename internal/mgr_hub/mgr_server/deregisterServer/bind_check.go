package deregisterServer

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

	ser := x.in.GetServer()
	if ser == nil {
		return errors.New("server is empty")
	}

	if ser.GetEndpointKey() == "" {
		return errors.New("ServerName is empty")
	}

	return nil
}
