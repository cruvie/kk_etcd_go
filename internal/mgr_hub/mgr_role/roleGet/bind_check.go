package roleGet

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

	if x.in.GetName() == "" {
		return errors.New("name is empty")
	}

	return nil
}
