package middleware

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_http"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_jwt"
	"github.com/cruvie/kk_etcd_go/handler/service"
	"github.com/cruvie/kk_etcd_go/utils/api_resp"
	"github.com/cruvie/kk_etcd_go/utils/global_model"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

// JWTAuth jwt auth middleware
func JWTAuth(c *gin.Context) {

	token := global_model.GetAuthorizationToken(c)
	if token == "" {
		log.Println("token is empty")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		c.Abort()
		return
	}
	myClaims := kku_jwt.VerifyToken[string](token)
	//get user from etcd
	user, res := service.GetUser(myClaims.UserId)
	if res != 1 {
		log.Println("fail to get user from etcd")
		kku_http.ResponseProtoBuf(c, api_resp.FailCodeMsg(http.StatusUnauthorized, "LogIn again"))
		c.Abort()
		return
	}
	//store user to gin context
	global_model.SetLoginUser(c, user)
	c.Next()

}
