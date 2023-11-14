package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"

	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {
	stage := kku_stage.NewStage(c, kku_func.GetCurrentFunctionName())

	token := global_model.GetAuthorizationToken(c)
	if token == "" {
		slog.Info("token is empty")
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	myClaims := kku_jwt.VerifyToken[string](stage, token)
	//get user from etcd
	user, res := service.GetUser(stage, myClaims.UserId)
	if res != 1 {
		slog.Info("fail to get user from etcd")
		kku_http.ResponseProtoBuf(c, api_resp.Fail(stage, &api_resp.ApiResp{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	//store user to gin context
	global_model.SetLoginUser(c, user)
	c.Next()

}
