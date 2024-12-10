package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hRole handler.HRole

// roleAdd
//
//	@Tags			role
//	@Description	add role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleAddParam	body		kk_etcd_models.RoleAddParam	true	"RoleAddParam"
//	@Success		200				{object}	kk_etcd_models.RoleAddResponse
//	@Router			/role/roleAdd [post]
func roleAdd(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleAdd")
	defer span.End()
	var param kk_etcd_models.RoleAddParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleAdd(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}

// roleDelete
//
//	@Tags			role
//	@Description	delete role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleDeleteParam	body		kk_etcd_models.RoleDeleteParam	true	"RoleDeleteParam"
//	@Success		200				{object}	kk_etcd_models.RoleDeleteResponse
//	@Router			/role/roleDelete [post]
func roleDelete(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleDelete")
	defer span.End()
	var param kk_etcd_models.RoleDeleteParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleDelete(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}

// roleList
//
//	@Tags			role
//	@Description	list role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleListParam	body		kk_etcd_models.RoleListParam	true	"RoleListParam"
//	@Success		200				{object}	kk_etcd_models.RoleListResponse
//	@Router			/role/roleList [post]
func roleList(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleList")
	defer span.End()
	var param kk_etcd_models.RoleListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleList(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}

// roleGet
//
//	@Tags			role
//	@Description	get role
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleGetParam	body		kk_etcd_models.RoleGetParam	true	"RoleGetParam"
//	@Success		200				{object}	kk_etcd_models.RoleGetResponse
//	@Router			/role/roleGet [post]
func roleGet(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleGet")
	defer span.End()
	var param kk_etcd_models.RoleGetParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleGet(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}

// roleGrantPermission
//
//	@Tags			role
//	@Description	grant permission
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleGrantPermissionParam	body		kk_etcd_models.RoleGrantPermissionParam	true	"RoleGrantPermissionParam"
//	@Success		200							{object}	kk_etcd_models.RoleGrantPermissionResponse
//	@Router			/role/roleGrantPermission [post]
func roleGrantPermission(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleGrantPermission")
	defer span.End()
	var param kk_etcd_models.RoleGrantPermissionParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleGrantPermission(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}

// roleRevokePermission
//
//	@Tags			role
//	@Description	revoke permission
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleRevokePermissionParam	body		kk_etcd_models.RoleRevokePermissionParam	true	"RoleRevokePermissionParam"
//	@Success		200							{object}	kk_etcd_models.RoleRevokePermissionResponse
//	@Router			/role/roleRevokePermission [post]
func roleRevokePermission(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("roleRevokePermission")
	defer span.End()
	var param kk_etcd_models.RoleRevokePermissionParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponsePB(stage, err, nil)
		return
	}
	err, resp := hRole.RoleRevokePermission(stage, &param)
	kk_http.ResponsePB(stage, err, resp)
}
