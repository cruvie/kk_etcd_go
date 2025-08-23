package util_role

import (
	"context"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

func AddRole(stage *kk_stage.Stage, roleName string) error {
	_, err := global_model.GetClient(stage).RoleAdd(context.Background(), roleName)
	if err != nil && !utils.ErrorIs(err, rpctypes.ErrGRPCRoleAlreadyExist) {
		return err
	}
	return nil
}
