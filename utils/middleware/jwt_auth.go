package middleware

import (
	"gitee.com/gkdgo/kk_go_kit/kk_utils/kku_http"
	"gitee.com/gkdgo/kk_go_kit/kk_utils/kku_jwt"
	"github.com/gin-gonic/gin"
	"kk_etcd_go/handler/service"
	"kk_etcd_go/utils/api_resp"
	"kk_etcd_go/utils/global_model"
	"log"
	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {

	global_model.ClearLoginUser(c)

	if c.Request.URL.Path == "/User/Login" {
		c.Next()
		return
	}
	token := global_model.GetAuthorizationToken(c)
	if token == "" {
		log.Println("token is empty")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		return
	}
	myClaims := kku_jwt.VerifyToken(token)
	//get user from etcd
	user, res := service.GetUser(myClaims.UserId.(string))
	if res != 1 {
		log.Println("fail to get user from etcd")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		return
	}
	//store user to gin context
	global_model.SetLoginUser(c, user)
	c.Next()

}
