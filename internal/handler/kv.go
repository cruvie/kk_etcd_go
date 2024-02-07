package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"

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
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckWritePermission(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have write permission!"}, nil))
		return
	}
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	err := service.KVPut(stage, pbKV.GetKey(), pbKV.GetValue())
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// KVGet
//
//	@Description	Get config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	kk_etcd_models.PBKV	true	"Get config info"
//	@Router			/KVGet [post]
func KVGet(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res, value := service.KVGet(stage, pbKV.Key)
	switch res {
	case 1:
		pbKV.Value = string(value)
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, &pbKV))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// KVList
//
//	@Description	Get kv list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/KVList [post]
func KVList(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var prefix kk_pb_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &prefix); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res, list := service.KVList(stage, prefix.Value)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, list))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}

// KVDel
//
//	@Description	Del config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	kk_etcd_models.PBKV	true	"Del config info"
//	@Router			/KVDel [post]
func KVDel(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckWritePermission(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have write permission!"}, nil))
		return
	}
	var pbKV kk_etcd_models.PBKV
	if err := kk_http.ReadProtoBuf(stage, &pbKV); err != nil {
		slog.Error("failed to read proto buf", "err", err)
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	res := service.KVDel(stage, pbKV.Key)
	switch res {
	case 1:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
}
