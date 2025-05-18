package serverList

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

	if x.In.GetServerType() == kk_etcd_models.PBServerType_Unknown {
		return errors.New("ServerType is Unknow")
	}

	return nil
}
