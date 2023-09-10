package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/models"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/check_user"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// KVPutConfig
//
//	@Description	Put config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Put config info"
//	@Router			/KVPutConfig [post]
func KVPutConfig(c *gin.Context) {
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
	res := service.KVPut(key_prefix.Config+pbKV.Key, pbKV.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVGetConfig
//
//	@Description	Get config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Get config info"
//	@Router			/KVGetConfig [post]
func KVGetConfig(c *gin.Context) {
	var pbKV models.PBKV
	if err := kku_http.ReadProtoBuf(c, &pbKV); err != nil {
		slog.Info("failed to read proto buf:", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res, value := service.KVGet(key_prefix.Config + pbKV.Key)
	switch res {
	case 1:
		pbKV.Value = string(value)
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(&pbKV))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVGetConfigList
//
//	@Description	Get config list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/KVGetConfigList [post]
func KVGetConfigList(c *gin.Context) {
	res, list := service.KVGetConfigList()
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(list))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}

// KVDelConfig
//
//	@Description	Del config
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			pbKV	body	models.PBKV	true	"Del config info"
//	@Router			/KVDelConfig [post]
func KVDelConfig(c *gin.Context) {
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
	res := service.KVDel(key_prefix.Config + pbKV.Key)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.Success())
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}
