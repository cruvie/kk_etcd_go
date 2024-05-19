package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hRole handler.HRole

// roleAdd
//
//	@Description	roleAdd
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleAddParam	body		kk_etcd_models.RoleAddParam	true	"RoleAddParam"
//	@Success		200				{object}	kk_etcd_models.RoleAddResponse
//	@Router			/role/roleAdd [post]
func roleAdd(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.RoleAddParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hRole.RoleAdd(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// roleDelete
//
//	@Description	roleDelete
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleDeleteParam	body		kk_etcd_models.RoleDeleteParam	true	"RoleDeleteParam"
//	@Success		200				{object}	kk_etcd_models.RoleDeleteResponse
//	@Router			/role/roleDelete [post]
func roleDelete(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.RoleDeleteParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hRole.RoleDelete(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// roleList
//
//	@Description	roleList
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleListParam	body		kk_etcd_models.RoleListParam	true	"RoleListParam"
//	@Success		200				{object}	kk_etcd_models.RoleListResponse
//	@Router			/role/roleList [post]
func roleList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.RoleListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hRole.RoleList(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// roleGet
//
//	@Description	roleGet
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleGetParam	body		kk_etcd_models.RoleGetParam	true	"RoleGetParam"
//	@Success		200				{object}	kk_etcd_models.RoleGetResponse
//	@Router			/role/roleGet [post]
func roleGet(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.RoleGetParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hRole.RoleGet(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// roleGrantPermission
//
//	@Description	roleGrantPermission
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			RoleGrantPermissionParam	body		kk_etcd_models.RoleGrantPermissionParam	true	"RoleGrantPermissionParam"
//	@Success		200							{object}	kk_etcd_models.RoleGrantPermissionResponse
//	@Router			/role/roleGrantPermission [post]
func roleGrantPermission(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.RoleGrantPermissionParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hRole.RoleGrantPermission(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}
