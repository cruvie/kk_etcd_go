package deregisterService

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}

	ser := x.In.GetService()
	if ser == nil {
		return errors.New("server is empty")
	}

	if ser.GetServiceRegistration() == nil {
		return errors.New("ServiceRegistration is empty")
	}

	return nil
}
