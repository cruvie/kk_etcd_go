package serviceList

import (
	"errors"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *Api) CheckInput(stage *kk_stage.Stage) error {

	if x.In.GetServiceType() == kk_etcd_models.PBServiceType_Unknown {
		return errors.New("ServiceType is Unknow")
	}

	return nil
}
