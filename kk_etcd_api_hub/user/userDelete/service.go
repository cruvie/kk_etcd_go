package userDelete

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *api) service(userName string) error {
	span := x.stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(x.stage).UserDelete(context.Background(), userName)
	if err != nil {
		return err
	}
	return err
}
