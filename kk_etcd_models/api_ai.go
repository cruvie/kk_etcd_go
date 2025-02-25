package kk_etcd_models

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *QueryParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}
	if x.Question == "" {
		return errors.New("question is empty")
	}

	return nil
}
