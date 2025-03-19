package util_role

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_error"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

func AddRole(stage *kk_stage.Stage, roleName string) error {
	_, err := global_model.GetClient(stage).RoleAdd(context.Background(), roleName)
	if err != nil && !kk_etcd_error.ErrorIs(err, rpctypes.ErrGRPCRoleAlreadyExist) {
		return err
	}
	return nil
}
