package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hKV handler.HKV

// kVPut
//
//	@Description	kVPut
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			KVPutParam	body		kk_etcd_models.KVPutParam	true	"KVPutParam"
//	@Success		200					{object}	kk_etcd_models.KVPutResponse
//	@Router			/kVPut [post]
func kVPut(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.KVPutParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hKV.KVPut(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// kVGet
//
//	@Description	kVGet
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			KVGetParam	body		kk_etcd_models.KVGetParam	true	"KVGetParam"
//	@Success		200					{object}	kk_etcd_models.KVGetResponse
//	@Router			/kVGet [post]
func kVGet(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.KVGetParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hKV.KVGet(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// kVDel
//
//	@Description	kVDel
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			KVDelParam	body		kk_etcd_models.KVDelParam	true	"KVDelParam"
//	@Success		200					{object}	kk_etcd_models.KVDelResponse
//	@Router			/kVDel [post]
func kVDel(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.KVDelParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hKV.KVDel(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// kVList
//
//	@Description	kVList
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			KVListParam	body		kk_etcd_models.KVListParam	true	"KVListParam"
//	@Success		200					{object}	kk_etcd_models.KVListResponse
//	@Router			/kVList [post]
func kVList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.KVListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hKV.KVList(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}
