package roleGrantPermission

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *Api) CheckInput(stage *kk_stage.Stage) error {

	if x.In.GetName() == "" {
		return errors.New("name is empty")
	}

	return nil
}
