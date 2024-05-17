package handler

import (
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"

	"github.com/cruvie/kk_etcd_go/internal/handler/service"

	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type HUser struct{}

var serUser service.SerUser

func (HUser) Login(stage *kk_stage.Stage, param *kk_etcd_models.LoginParam) (error, *kk_etcd_models.LoginResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	if param.UserName == kk_etcd_const.UserRoot {
		return errors.New("illegal login root user"), nil
	}
	tokenString, err := serUser.Login(stage, param)
	return err, &kk_etcd_models.LoginResponse{
		Token: tokenString,
	}
}

func (HUser) Logout(stage *kk_stage.Stage, param *kk_etcd_models.LogoutParam) (error, *kk_etcd_models.LogoutResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serUser.Logout(stage, param)
	return err, &kk_etcd_models.LogoutResponse{}
}

func (HUser) UserAdd(stage *kk_stage.Stage, param *kk_etcd_models.UserAddParam) (error, *kk_etcd_models.UserAddResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetUserName() == kk_etcd_const.UserRoot {
		return errors.New("illegal add root user"), nil
	}
	err = serUser.UserAdd(stage, &kk_etcd_models.PBUser{
		UserName: param.GetUserName(),
		Password: param.GetPassword(),
		Roles:    param.GetRoles(),
	})
	return err, &kk_etcd_models.UserAddResponse{}
}

func (HUser) UserDelete(stage *kk_stage.Stage, param *kk_etcd_models.UserDeleteParam) (error, *kk_etcd_models.UserDeleteResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetUserName() == kk_etcd_const.UserRoot ||
		param.GetUserName() == config.Config.Admin.UserName ||
		param.GetUserName() == global_model.GetLoginUser(stage).UserName {
		return errors.New("illegal delete root admin or current logged in user"), nil
	}
	err = serUser.UserDelete(stage, param.GetUserName())
	return err, &kk_etcd_models.UserDeleteResponse{}
}

func (HUser) GetUser(stage *kk_stage.Stage, param *kk_etcd_models.GetUserParam) (error, *kk_etcd_models.GetUserResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	user, err := serUser.GetUser(param.GetUserName())
	return err, &kk_etcd_models.GetUserResponse{
		User: user,
	}
}

func (HUser) MyInfo(stage *kk_stage.Stage, _ *kk_etcd_models.MyInfoParam) (error, *kk_etcd_models.MyInfoResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	loginUser := global_model.GetLoginUser(stage)
	return nil, &kk_etcd_models.MyInfoResponse{
		UserName: loginUser.GetUserName(),
		Roles:    loginUser.GetRoles(),
	}
}

func (HUser) UserList(stage *kk_stage.Stage, _ *kk_etcd_models.UserListParam) (error, *kk_etcd_models.UserListResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	err, users := serUser.UserList()
	return err, &kk_etcd_models.UserListResponse{
		ListUser: users,
	}
}

func (HUser) UserGrantRole(stage *kk_stage.Stage, param *kk_etcd_models.UserGrantRoleParam) (error, *kk_etcd_models.UserGrantRoleResponse) {
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()

	err := serUser.CheckRootRole(stage)
	if err != nil {
		return err, nil
	}
	if param.GetUserName() == config.Config.Admin.UserName {
		return errors.New("can't change Admin user's role"), nil
	}
	if param.GetUserName() == kk_etcd_const.UserRoot {
		return errors.New("illegal modification of root user's role"), nil
	}
	err = serUser.UserGrantRole(&kk_etcd_models.PBUser{
		UserName: param.GetUserName(),
		Roles:    param.GetRoles(),
	})
	return err, &kk_etcd_models.UserGrantRoleResponse{}
}
