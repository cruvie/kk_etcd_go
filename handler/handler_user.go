package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/models"
	"github.com/cruvie/kk_etcd_go/models/base_proto_type"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/check_user"
	"github.com/cruvie/kk_etcd_go/utils/global_model"
	"github.com/gin-gonic/gin"

	"log"
)

// Login
//
//	@Description	Log in
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Login info"
//	@Router			/Login [post]
func Login(c *gin.Context) {
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	token, res := service.Login(&pbUser)
	switch res {
	case 1:
		tokenPB := base_proto_type.PBString{Value: token}
		kku_http.ResponseProtoBuf(c, api_resp.SuccessMsgData("Login Succeeded", &tokenPB))
		return
	case 2:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("User not exist"))
		return
	case 4:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("Wrong user name or password"))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// Logout
//
//	@Description	Log out
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Logout info"
//	@Router			/Logout [post]
func Logout(c *gin.Context) {
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.Logout(&pbUser)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}

	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// UserAdd
//
//	@Description	Add user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Add user info"
//	@Router			/UserAdd [post]
func UserAdd(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}

	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.UserAdd(&pbUser)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// UserDelete
//
//	@Description	Delete user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Delete user info"
//	@Router			/UserDelete [post]
func UserDelete(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.UserDelete(c, pbUser.UserName, false)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	case 2:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("illegal delete root admin or current logged in user!"))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// GetUser
//
//	@Description	Get user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Get user info"
//	@Router			/GetUser [post]
func GetUser(c *gin.Context) {
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	user, res := service.GetUser(pbUser.UserName)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(user))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// MyInfo
//
//	@Description	Get my info
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/MyInfo [post]
func MyInfo(c *gin.Context) {
	loginUser := global_model.GetLoginUser(c)
	user, res := service.GetUser(loginUser.UserName)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(user))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// UserList
//
//	@Description	Get user list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/UserList [post]
func UserList(c *gin.Context) {
	res, users := service.UserList()
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(users))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// UserGrantRole
//
//	@Description	Grant role to user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pBListString	body	base_proto_type.PBListString	true	"Grant role to user info"
//	@Router			/UserGrantRole [post]
func UserGrantRole(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}

	var user models.PBUser
	if err := kku_http.ReadProtoBuf(c, &user); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}

	if user.UserName == config.GlobalConfig.Admin.UserName {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("can't change Admin user's role!"))
		return
	}

	res := service.UserGrantRole(&user)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	case -1:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("Illegal modification of root user's role!"))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}
