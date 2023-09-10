package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/models"
	"github.com/cruvie/kk_etcd_go/models/base_proto_type"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/check_user"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// KVPut
//
//	@Description	Put config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Put config info"
//	@Router			/KVPut [post]
func KVPut(c *gin.Context) {
	if !check_user.CheckWritePermission(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have write permission!"))
		return
	}
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		slog.Info("failed to read proto buf:", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.KVPut(pbKV.Key, pbKV.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVGet
//
//	@Description	Get config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Get config info"
//	@Router			/KVGet [post]
func KVGet(c *gin.Context) {
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		slog.Info("failed to read proto buf:", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res, value := service.KVGet(pbKV.Key)
	switch res {
	case 1:
		pbKV.Value = string(value)
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(&pbKV))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVList
//
//	@Description	Get kv list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/KVList [post]
func KVList(c *gin.Context) {
	var prefix base_proto_type.PBString
	if err := kku_http.ReadProtoBuf(c, &prefix); err != nil {
		slog.Info("failed to read proto buf:", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res, list := service.KVList(prefix.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(list))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVDel
//
//	@Description	Del config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Del config info"
//	@Router			/KVDel [post]
func KVDel(c *gin.Context) {
	if !check_user.CheckWritePermission(c) {
		kku_http.ResponseProtoBuf(c, api_resp.FailMsg("you don't have write permission!"))
		return
	}
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		slog.Info("failed to read proto buf:", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res := service.KVDel(pbKV.Key)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}
