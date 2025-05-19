package serviceList

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}

	if x.In.GetServiceType() == kk_etcd_models.PBServiceType_Unknown {
		return errors.New("ServiceType is Unknow")
	}

	return nil
}
