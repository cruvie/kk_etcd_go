package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
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
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	res := service.RoleAdd(&pbRole)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// RoleDelete
//
//	@Description	Delete role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Delete role info"
//	@Router			/RoleDelete [post]
func RoleDelete(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	res := service.RoleDelete(pbRole.Name)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	case -1:
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "illegal delete root role!"}, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// RoleList
//
//	@Description	Get role list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/RoleList [post]
func RoleList(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	res, roles := service.RoleList()
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, roles))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// RoleGet
//
//	@Description	Get role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Get role info"
//	@Router			/RoleGet [post]
func RoleGet(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	role, res := service.RoleGet(pbRole.Name)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, role))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}

// RoleGrantPermission
//
//	@Description	Grant permission to role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Grant permission to role info"
//	@Router			/RoleGrantPermission [post]
func RoleGrantPermission(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(c) {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kku_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
		return
	}
	if pbRole.Name == kk_etcd_const.RoleRoot {
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Msg: "illegal change root role permission!"}, nil))
		return
	}
	res := service.RoleGrantPermission(&pbRole)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success(stage, nil, nil))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, nil, nil))
}
