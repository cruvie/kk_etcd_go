package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_response"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
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
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kk_http.ReadProtoBuf(stage, &pbRole); err != nil {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res := service.RoleAdd(stage, &pbRole)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// RoleDelete
//
//	@Description	Delete role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Delete role info"
//	@Router			/RoleDelete [post]
func RoleDelete(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kk_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res := service.RoleDelete(stage, pbRole.Name)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	case -1:
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "illegal delete root role!"}, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// RoleList
//
//	@Description	Get role list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/RoleList [post]
func RoleList(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	res, roles := service.RoleList(stage)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, roles))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// RoleGet
//
//	@Description	Get role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Get role info"
//	@Router			/RoleGet [post]
func RoleGet(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	var pbRole kk_etcd_models.PBRole
	if err := kk_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	role, res := service.RoleGet(stage, pbRole.Name)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, role))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// RoleGrantPermission
//
//	@Description	Grant permission to role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbRole	body	kk_etcd_models.PBRole	true	"Grant permission to role info"
//	@Router			/RoleGrantPermission [post]
func RoleGrantPermission(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var pbRole kk_etcd_models.PBRole
	if err := kk_http.ReadProtoBuf(stage, &pbRole); err != nil {
		slog.Info("failed to read protobuf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	if pbRole.Name == kk_etcd_const.RoleRoot {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "illegal change root role permission!"}, nil))
		return
	}
	res := service.RoleGrantPermission(stage, &pbRole)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}
