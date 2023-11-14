package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models/base_proto_type"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// Login
//
//	@Description	Log in
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Login info"
//	@Router			/Login [post]
func Login(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())

	var pbUser kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &pbUser); err != nil {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	token, res := service.Login(stage, &pbUser)
	switch res {
	case 1:
		tokenPB := base_proto_type.PBString{Value: token}
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, &api_resp.ApiResp{
			Msg: "Login Succeeded"}, &tokenPB))
		return
	case 2:
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "User not exist"}, nil))
		return
	case 4:
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "Wrong user name or password"}, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// Logout
//
//	@Description	Log out
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Logout info"
//	@Router			/Logout [post]
func Logout(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	var pbUser kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	res := service.Logout(stage, &pbUser)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	}

	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// UserAdd
//
//	@Description	Add user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Add user info"
//	@Router			/UserAdd [post]
func UserAdd(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(stage) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}

	var pbUser kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	res := service.UserAdd(stage, &pbUser)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// UserDelete
//
//	@Description	Delete user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Delete user info"
//	@Router			/UserDelete [post]
func UserDelete(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(stage) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbUser kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	res := service.UserDelete(stage, pbUser.UserName, false)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	case 2:
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "illegal delete root admin or current logged in user!"}, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// GetUser
//
//	@Description	Get user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Get user info"
//	@Router			/GetUser [post]
func GetUser(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	var pbUser kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	user, res := service.GetUser(stage, pbUser.UserName)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, user))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// MyInfo
//
//	@Description	Get my info
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/MyInfo [post]
func MyInfo(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	loginUser := global_model.GetLoginUser(c)
	user, res := service.GetUser(stage, loginUser.UserName)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, user))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// UserList
//
//	@Description	Get user list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/UserList [post]
func UserList(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	res, users := service.UserList(stage)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, users))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// UserGrantRole
//
//	@Description	Grant role to user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pBListString	body	base_proto_type.PBListString	true	"Grant role to user info"
//	@Router			/UserGrantRole [post]
func UserGrantRole(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(stage) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}

	var user kk_etcd_models.PBUser
	if err := kku_http.ReadProtoBuf(stage, &user); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}

	if user.UserName == config.Config.Admin.UserName {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "can't change Admin user's role!"}, nil))
		return
	}

	res := service.UserGrantRole(stage, &user)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	case -1:
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "Illegal modification of root user's role!"}, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}
