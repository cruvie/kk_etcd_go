package roleRevokePermission

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(x.stage).RoleRevokePermission(
		context.Background(),
		x.in.GetName(),
		x.in.GetKey(),
		x.in.GetRangeEnd())
	if err != nil {
		return err
	}
	return nil
}
