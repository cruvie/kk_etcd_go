package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_jwt"
	"gitee.com/cruvie/kk_go_kit/kk_models"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model/global_stage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {
	stage := global_stage.GetRequestStage(c)

	token := global_model.GetAuthorizationToken(stage)
	if token == "" {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_models.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	myClaims, err := kk_jwt.VerifyToken[string](token)
	if err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_models.PBResponse{
			Code: http.StatusUnauthorized,
			Msg:  "LogIn again"}, nil))
		c.Abort()
		return
	}
	//get user from etcd
	var serUser service.SerUser
	user, err := serUser.GetUser(myClaims.Payload.Username)
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

}
