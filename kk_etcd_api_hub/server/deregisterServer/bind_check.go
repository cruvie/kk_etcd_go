package deregisterServer

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}

	ser := x.In.GetServer()
	if ser == nil {
		return errors.New("server is empty")
	}

	if ser.GetServerRegistration() == nil {
		return errors.New("ServerRegistration is empty")
	}

	return nil
}
