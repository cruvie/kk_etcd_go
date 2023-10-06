package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/api_resp"
	global_model2 "github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/gin-gonic/gin"
	"log/slog"

	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {

	token := global_model2.GetAuthorizationToken(c)
	if token == "" {
		slog.Info("token is empty")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		c.Abort()
		return
	}
	myClaims := kku_jwt.VerifyToken[string](token)
	//get user from etcd
	user, res := service.GetUser(myClaims.UserId)
	if res != 1 {
		slog.Info("fail to get user from etcd")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		c.Abort()
		return
	}
	//store user to gin context
	global_model2.SetLoginUser(c, user)
	c.Next()

}
