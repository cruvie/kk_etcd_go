package roleGrantPermission

import (
	"context"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (x *api) service() error {
	span := x.stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(x.stage).RoleGrantPermission(
		context.Background(),
		x.in.GetName(),
		x.in.GetPerm().GetKey(),
		x.in.GetPerm().GetRangeEnd(),
		clientv3.PermissionType(x.in.GetPerm().GetPermissionType()))
	if err != nil {
		return err
	}
	return nil
}
