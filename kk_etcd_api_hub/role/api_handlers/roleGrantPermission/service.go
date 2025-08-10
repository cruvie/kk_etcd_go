package roleGrantPermission

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("service")
	defer span.End()

	_, err := global_model.GetClient(stage).RoleGrantPermission(
		context.Background(),
		x.In.GetName(),
		x.In.GetPerm().GetKey(),
		x.In.GetPerm().GetRangeEnd(),
		clientv3.PermissionType(x.In.GetPerm().GetPermissionType()))
	if err != nil {
		return err
	}
	return nil
}
