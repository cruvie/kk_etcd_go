package roleDelete

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(x.stage).RoleDelete(context.Background(), x.in.GetName())
	if err != nil {
		return err
	}
	return nil
}
