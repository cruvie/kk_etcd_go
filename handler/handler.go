package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/gin-gonic/gin"
	"kk_etcd_go/consts"
	"kk_etcd_go/handler/service"
	"kk_etcd_go/models"
	"kk_etcd_go/models/base_proto_type"
	"kk_etcd_go/utils/api_resp"
	"kk_etcd_go/utils/global_model"
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

// AddUser
//
//	@Description	Add user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Add user info"
//	@Router			/AddUser [post]
func AddUser(c *gin.Context) {
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.AddUser(&pbUser)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// DeleteUser
//
//	@Description	Delete user
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbUser	body	models.PBUser	true	"Delete user info"
//	@Router			/DeleteUser [post]
func DeleteUser(c *gin.Context) {
	var pbUser models.PBUser
	if err := kku_http.ReadProtoBuf(c, &pbUser); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.DeleteUser(c, pbUser.UserName, false)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	case 2:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("illegal delete root admin or your own user!"))
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

// AddRole
//
//	@Description	Add role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	models.PBRole	true	"Add role info"
//	@Router			/AddRole [post]
func AddRole(c *gin.Context) {
	var pbRole models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.AddRole(&pbRole)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// DeleteRole
//
//	@Description	Delete role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	models.PBRole	true	"Delete role info"
//	@Router			/DeleteRole [post]
func DeleteRole(c *gin.Context) {
	var pbRole models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.DeleteRole(pbRole.Name)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// RoleList
//
//	@Description	Get role list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/RoleList [post]
func RoleList(c *gin.Context) {
	res, roles := service.RoleList()
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(roles))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// RoleGet
//
//	@Description	Get role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	models.PBRole	true	"Get role info"
//	@Router			/RoleGet [post]
func RoleGet(c *gin.Context) {
	var pbRole models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	role, res := service.RoleGet(pbRole.Name)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(role))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// RoleGrantPermission
//
//	@Description	Grant permission to role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	models.PBRole	true	"Grant permission to role info"
//	@Router			/RoleGrantPermission [post]
func RoleGrantPermission(c *gin.Context) {
	var pbRole models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.RoleGrantPermission(&pbRole)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
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
	var pBListString base_proto_type.PBListString
	if err := kku_http.ReadProtoBuf(c, &pBListString); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.UserGrantRole(pBListString.ListString[0], pBListString.ListString[1])
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVPutConfig
//
//	@Description	Put config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Put config info"
//	@Router			/KVPutConfig [post]
func KVPutConfig(c *gin.Context) {
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.KVPut(consts.EtcdConfig+pbKV.Key, pbKV.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVGetConfig
//
//	@Description	Get config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Get config info"
//	@Router			/KVGetConfig [post]
func KVGetConfig(c *gin.Context) {
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res, value := service.KVGet(consts.EtcdConfig + pbKV.Key)
	switch res {
	case 1:
		pbKV.Value = string(value)
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(&pbKV))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVGetConfigList
//
//	@Description	Get config list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/KVGetConfigList [post]
func KVGetConfigList(c *gin.Context) {
	res, value := service.KVGetConfigList()
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(value))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVDelConfig
//
//	@Description	Del config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Del config info"
//	@Router			/KVDelConfig [post]
func KVDelConfig(c *gin.Context) {
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		log.Println(err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.KVDel(consts.EtcdConfig + pbKV.Key)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}
