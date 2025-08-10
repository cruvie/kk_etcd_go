package query

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *Api) CheckInput(stage *kk_stage.Stage) error {

	if x.In.Question == "" {
		return errors.New("question is empty")
	}

	return nil
}
