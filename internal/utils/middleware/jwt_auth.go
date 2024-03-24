package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"

	"github.com/gin-gonic/gin"
	"log/slog"

	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {
	stage := global_model.GetRequestStage(c)

	token := global_model.GetAuthorizationToken(c)
	if token == "" {
		slog.Error("token is empty")
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	myClaims, err := kk_jwt.VerifyToken[string](token)
	if err != nil {
		slog.Error("fail to verify user from etcd")
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	//get user from etcd
	user, res := service.GetUser(stage, myClaims.UserId)
	if res != 1 {
		slog.Error("fail to get user from etcd")
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	//store user to gin context
	global_model.SetLoginUser(c, user)
	c.Next()

}
