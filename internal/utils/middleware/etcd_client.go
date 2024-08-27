package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"net/http"
	"time"
)

// EtcdClient set etcd client for current request
func EtcdClient(c *gin.Context) {
	stage := global_stage.GetRequestStage(c)
	//todo use https://github.com/etcd-io/etcd/pull/16803
	// set verify jwt directly
	//slog.Info("url", c.Request.URL.Path)
	cfg := clientv3.Config{
		Endpoints:   []string{config.Config.Etcd.Endpoint},
		DialTimeout: 5 * time.Second,
		Username:    global_model.GetRequestHeader(stage).UserName,
		Password:    global_model.GetRequestHeader(stage).Password,
	}
	var err error
	kk_etcd_client.EtcdClient, err = clientv3.New(cfg)
	if err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_models.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  err.Error()}, nil))
		c.Abort()
		return
	}
	var serUser service.SerUser
	user, err := serUser.GetUser(global_model.GetRequestHeader(stage).UserName)
	if err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_models.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	//store user to gin context
	global_model.SetLoginUser(stage, user)
	c.Next()
	//reset etcd client after request
	err = kk_etcd_client.EtcdClient.Close()
	if err != nil {
		slog.Error("etcd client close error", "err", err)
	}
	kk_etcd_client.EtcdClient = nil
}
