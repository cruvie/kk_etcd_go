package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hKV handler.HKV

// kVPut
//
//	@Tags			kv
//	@Description	put kv
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			KVPutParam	body		kk_etcd_models.KVPutParam	true	"KVPutParam"
//	@Success		200			{object}	kk_etcd_models.KVPutResponse
//	@Router			/kv/kVPut [post]
func kVPut(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("kVPut")
	defer span.End()
	var param kk_etcd_models.KVPutParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hKV.KVPut(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}

// kVGet
//
//	@Tags			kv
//	@Description	get kv
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			KVGetParam	body		kk_etcd_models.KVGetParam	true	"KVGetParam"
//	@Success		200			{object}	kk_etcd_models.KVGetResponse
//	@Router			/kv/kVGet [post]
func kVGet(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("kVGet")
	defer span.End()
	var param kk_etcd_models.KVGetParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hKV.KVGet(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}

// kVDel
//
//	@Tags			kv
//	@Description	del kv
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			KVDelParam	body		kk_etcd_models.KVDelParam	true	"KVDelParam"
//	@Success		200			{object}	kk_etcd_models.KVDelResponse
//	@Router			/kv/kVDel [post]
func kVDel(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("kVDel")
	defer span.End()
	var param kk_etcd_models.KVDelParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hKV.KVDel(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}

// kVList
//
//	@Tags			kv
//	@Description	list kv
//	@Accept			application/x-protobuf,json
//	@Produce		application/x-protobuf,json
//	@Param			KVListParam	body		kk_etcd_models.KVListParam	true	"KVListParam"
//	@Success		200			{object}	kk_etcd_models.KVListResponse
//	@Router			/kv/kVList [post]
func kVList(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("kVList")
	defer span.End()
	var param kk_etcd_models.KVListParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.WriteResponse(stage, err, nil)
		return
	}
	err, resp := hKV.KVList(stage, &param)
	kk_http.WriteResponse(stage, err, resp)
}
