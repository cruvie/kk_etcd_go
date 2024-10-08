package handler

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HRole struct{}

var serRole service.SerRole

func (HRole) RoleAdd(stage *kk_stage.Stage, param *kk_etcd_models.RoleAddParam) (error, *kk_etcd_models.RoleAddResponse) {
	span := stage.StartTrace("RoleAdd")
	defer span.End()

	err := serRole.RoleAdd(stage, param)
	return err, &kk_etcd_models.RoleAddResponse{}
}

func (HRole) RoleDelete(stage *kk_stage.Stage, param *kk_etcd_models.RoleDeleteParam) (error, *kk_etcd_models.RoleDeleteResponse) {
	span := stage.StartTrace("RoleDelete")
	defer span.End()

	if param.GetName() == consts.RoleRoot {
		return errors.New("illegal delete root role"), nil
	}
	err := serRole.RoleDelete(stage, param.GetName())
	return err, &kk_etcd_models.RoleDeleteResponse{}
}
func (HRole) RoleGet(stage *kk_stage.Stage, param *kk_etcd_models.RoleGetParam) (error, *kk_etcd_models.RoleGetResponse) {
	span := stage.StartTrace("RoleGet")
	defer span.End()
	role, err := serRole.RoleGet(stage, param.GetName())
	return err, &kk_etcd_models.RoleGetResponse{
		Role: role,
	}
}

func (HRole) RoleList(stage *kk_stage.Stage, _ *kk_etcd_models.RoleListParam) (error, *kk_etcd_models.RoleListResponse) {
	span := stage.StartTrace("RoleList")
	defer span.End()
	err, roles := serRole.RoleList(stage)
	return err, &kk_etcd_models.RoleListResponse{
		ListRole: roles,
	}
}

func (HRole) RoleGrantPermission(stage *kk_stage.Stage, param *kk_etcd_models.RoleGrantPermissionParam) (error, *kk_etcd_models.RoleGrantPermissionResponse) {
	span := stage.StartTrace("RoleGrantPermission")
	defer span.End()

	err := serRole.RoleGrantPermission(stage, param)
	return err, &kk_etcd_models.RoleGrantPermissionResponse{}
}

func (HRole) RoleRevokePermission(stage *kk_stage.Stage, param *kk_etcd_models.RoleRevokePermissionParam) (error, *kk_etcd_models.RoleRevokePermissionResponse) {
	span := stage.StartTrace("RoleRevokePermission")
	defer span.End()
	err := serRole.RoleRevokePermission(stage, param)
	if err != nil {
		return err, nil
	}
	return nil, &kk_etcd_models.RoleRevokePermissionResponse{}
}
