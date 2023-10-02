package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/check_user"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// RoleAdd
//
//	@Description	Add role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Add role info"
//	@Router			/RoleAdd [post]
func RoleAdd(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.RoleAdd(&pbRole)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// RoleDelete
//
//	@Description	Delete role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Delete role info"
//	@Router			/RoleDelete [post]
func RoleDelete(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.RoleDelete(pbRole.Name)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	case -1:
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("illegal delete root role!"))
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
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Get role info"
//	@Router			/RoleGet [post]
func RoleGet(c *gin.Context) {
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
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
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Grant permission to role info"
//	@Router			/RoleGrantPermission [post]
func RoleGrantPermission(c *gin.Context) {
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have root role!"))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(c, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	if pbRole.Name == consts.RoleRoot {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("illegal change root role permission!"))
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
