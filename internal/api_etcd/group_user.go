package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hUser handler.HUser

// login
//
//	@Description	login
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			LoginParam	body		kk_etcd_models.LoginParam	true	"LoginParam"
//	@Success		200			{object}	kk_etcd_models.LoginResponse
//	@Router			/user/login [post]
func login(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.LoginParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.Login(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// logout
//
//	@Description	logout
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			LogoutParam	body		kk_etcd_models.LogoutParam	true	"LogoutParam"
//	@Success		200			{object}	kk_etcd_models.LogoutResponse
//	@Router			/user/logout [post]
func logout(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.LogoutParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.Logout(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// userAdd
//
//	@Description	userAdd
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			UserAddParam	body		kk_etcd_models.UserAddParam	true	"UserAddParam"
//	@Success		200				{object}	kk_etcd_models.UserAddResponse
//	@Router			/user/userAdd [post]
func userAdd(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.UserAddParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.UserAdd(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// userDelete
//
//	@Description	userDelete
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			UserDeleteParam	body		kk_etcd_models.UserDeleteParam	true	"UserDeleteParam"
//	@Success		200				{object}	kk_etcd_models.UserDeleteResponse
//	@Router			/user/userDelete [post]
func userDelete(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.UserDeleteParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.UserDelete(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// getUser
//
//	@Description	getUser
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			GetUserParam	body		kk_etcd_models.GetUserParam	true	"GetUserParam"
//	@Success		200				{object}	kk_etcd_models.GetUserResponse
//	@Router			/user/getUser [post]
func getUser(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.GetUserParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.GetUser(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// myInfo
//
//	@Description	myInfo
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			MyInfoParam	body		kk_etcd_models.MyInfoParam	true	"MyInfoParam"
//	@Success		200			{object}	kk_etcd_models.MyInfoResponse
//	@Router			/user/myInfo [post]
func myInfo(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.MyInfoParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.MyInfo(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// userList
//
//	@Description	userList
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			UserListParam	body		kk_etcd_models.UserListParam	true	"UserListParam"
//	@Success		200				{object}	kk_etcd_models.UserListResponse
//	@Router			/user/userList [post]
func userList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.UserListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.UserList(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// userGrantRole
//
//	@Description	userGrantRole
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			UserGrantRoleParam	body		kk_etcd_models.UserGrantRoleParam	true	"UserGrantRoleParam"
//	@Success		200					{object}	kk_etcd_models.UserGrantRoleResponse
//	@Router			/user/userGrantRole [post]
func userGrantRole(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.UserGrantRoleParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hUser.UserGrantRole(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}
