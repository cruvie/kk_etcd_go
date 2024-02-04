package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_base_proto_type"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_response"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"

	"github.com/cruvie/kk_etcd_go/kk_etcd_models"

	"github.com/gin-gonic/gin"
	"log/slog"
)

// KVPut
//
//	@Description	Put config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	kk_etcd_models.PBKV	true	"Put config info"
//	@Router			/KVPut [post]
func KVPut(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	if !check_user.CheckWritePermission(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have write permission!"}, nil))
		return
	}
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res := service.KVPut(stage, pbKV.Key, pbKV.Value)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// KVGet
//
//	@Description	Get config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	kk_etcd_models.PBKV	true	"Get config info"
//	@Router			/KVGet [post]
func KVGet(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res, value := service.KVGet(stage, pbKV.Key)
	switch res {
	case 1:
		pbKV.Value = string(value)
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, &pbKV))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// KVList
//
//	@Description	Get kv list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/KVList [post]
func KVList(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	var prefix kk_base_proto_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &prefix); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res, list := service.KVList(stage, prefix.Value)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, list))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}

// KVDel
//
//	@Description	Del config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	kk_etcd_models.PBKV	true	"Del config info"
//	@Router			/KVDel [post]
func KVDel(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName(), config.Config.DebugMode)
	if !check_user.CheckWritePermission(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have write permission!"}, nil))
		return
	}
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		slog.Error("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	res := service.KVDel(stage, pbKV.Key)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
}
