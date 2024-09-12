package kk_etcd_models

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *ServerListParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}

	if x.GetServerType() == "" {
		return errors.New("ServerType is empty")
	}

	return nil
}

func (x *DeregisterServerParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadProtoBuf(stage, x)
	if err != nil {
		return err
	}
	ser := x.GetServer()
	if ser == nil {
		return errors.New("server is empty")
	}

	if ser.GetEndpointKey() == "" {
		return errors.New("ServerName is empty")
	}

	return nil
}
