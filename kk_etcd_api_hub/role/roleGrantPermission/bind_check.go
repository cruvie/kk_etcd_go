package roleGrantPermission

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_http"
)

func (x *api) bindCheck() error {
	err := kk_http.ReadReq(x.stage, x.In)
	if err != nil {
		return err
	}

	if x.In.GetName() == "" {
		return errors.New("name is empty")
	}

	return nil
}
