package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"

	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"

	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"

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
	stage := global_model.GetRequestStage(c)

	var pbUser kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &pbUser); err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	token, res := service.Login(stage, &pbUser)
	switch res {
	case 1:
		tokenPB := kk_pb_type.PBString{Value: token}
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, &kk_pb_type.PBResponse{
			Msg: "Login Succeeded"}, &tokenPB))
		return
	case 2:
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "User not exist"}, nil))
		return
	case 4:
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "Wrong user name or password"}, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// Logout
//
//	@Description	Log out
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Logout info"
//	@Router			/Logout [post]
func Logout(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var pbUser kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res := service.Logout(stage, &pbUser)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	}

	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// UserAdd
//
//	@Description	Add user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Add user info"
//	@Router			/UserAdd [post]
func UserAdd(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}

	var pbUser kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res := service.UserAdd(stage, &pbUser)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// UserDelete
//
//	@Description	Delete user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Delete user info"
//	@Router			/UserDelete [post]
func UserDelete(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbUser kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res := service.UserDelete(stage, pbUser.UserName, false)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	case 2:
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "illegal delete root admin or current logged in user!"}, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// GetUser
//
//	@Description	Get user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	kk_etcd_models.PBUser	true	"Get user info"
//	@Router			/GetUser [post]
func GetUser(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var pbUser kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &pbUser); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	user, res := service.GetUser(stage, pbUser.UserName)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, user))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// MyInfo
//
//	@Description	Get my info
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/MyInfo [post]
func MyInfo(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	loginUser := global_model.GetLoginUser(c)
	user, res := service.GetUser(stage, loginUser.UserName)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, user))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// UserList
//
//	@Description	Get user list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/UserList [post]
func UserList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	res, users := service.UserList(stage)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, users))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// UserGrantRole
//
//	@Description	Grant role to user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			PBUser	body	kk_etcd_models.PBUser	true	"Grant role to user info"
//	@Router			/UserGrantRole [post]
func UserGrantRole(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}

	var user kk_etcd_models.PBUser
	if err := kk_http.ReadProtoBuf(stage, &user); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}

	if user.UserName == config.Config.Admin.UserName {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "can't change Admin user's role!"}, nil))
		return
	}

	res := service.UserGrantRole(stage, &user)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	case -1:
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "Illegal modification of root user's role!"}, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}
