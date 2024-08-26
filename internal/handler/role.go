package handler

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_reflect"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HRole struct{}

var serRole service.SerRole

func (HRole) RoleAdd(stage *kk_stage.Stage, param *kk_etcd_models.RoleAddParam) (error, *kk_etcd_models.RoleAddResponse) {
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetName() == kk_etcd_const.RoleRoot {
		return errors.New("illegal add root role"), nil
	}
	err = serRole.RoleAdd(param)
	return err, &kk_etcd_models.RoleAddResponse{}
}

func (HRole) RoleDelete(stage *kk_stage.Stage, param *kk_etcd_models.RoleDeleteParam) (error, *kk_etcd_models.RoleDeleteResponse) {
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetName() == kk_etcd_const.RoleRoot {
		return errors.New("illegal delete root role"), nil
	}
	err = serRole.RoleDelete(param.GetName())
	return err, &kk_etcd_models.RoleDeleteResponse{}
}
func (HRole) RoleGet(stage *kk_stage.Stage, param *kk_etcd_models.RoleGetParam) (error, *kk_etcd_models.RoleGetResponse) {
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	role, err := serRole.RoleGet(param.GetName())
	return err, &kk_etcd_models.RoleGetResponse{
		Role: role,
	}
}

func (HRole) RoleList(stage *kk_stage.Stage, _ *kk_etcd_models.RoleListParam) (error, *kk_etcd_models.RoleListResponse) {
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	err, roles := serRole.RoleList()
	return err, &kk_etcd_models.RoleListResponse{
		ListRole: roles,
	}
}

func (HRole) RoleGrantPermission(stage *kk_stage.Stage, param *kk_etcd_models.RoleGrantPermissionParam) (error, *kk_etcd_models.RoleGrantPermissionResponse) {
	span := stage.StartTrace(kk_reflect.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetRole().GetName() == kk_etcd_const.RoleRoot {
		return errors.New("illegal change root role permission"), nil
	}
	err = serRole.RoleGrantPermission(param.GetRole())
	return err, &kk_etcd_models.RoleGrantPermissionResponse{}
}
