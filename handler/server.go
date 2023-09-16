package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/models/base_proto_type"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// ServerList
//
//	@Description	Get service list
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Router			/ServerList [post]
func ServerList(c *gin.Context) {
	var prefix base_proto_type.PBString
	if err := kku_http.ReadProtoBuf(c, &prefix); err != nil {
		slog.Info("failed to read proto buf", "err", err)
		kku_http.ResponseProtoBuf(c, api_resp.Fail())
		return
	}
	res, list, _ := service.ServerList(prefix.Value)
	switch res {
	case 1:
		kku_http.ResponseProtoBuf(c, api_resp.SuccessData(list))
		return
	}
	kku_http.ResponseProtoBuf(c, api_resp.Fail())
	return
}
