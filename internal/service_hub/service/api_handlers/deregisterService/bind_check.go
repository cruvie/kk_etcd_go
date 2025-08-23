package deregisterService

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *Api) CheckInput(stage *kk_stage.Stage) error {

	ser := x.In.GetService()
	if ser == nil {
		return errors.New("server is empty")
	}

	if ser.GetServiceRegistration() == nil {
		return errors.New("ServiceRegistration is empty")
	}

	return nil
}
