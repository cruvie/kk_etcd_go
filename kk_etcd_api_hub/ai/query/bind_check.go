package query

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}
	if x.In.Question == "" {
		return errors.New("question is empty")
	}

	return nil
}
