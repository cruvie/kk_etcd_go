package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/handler"
	"github.com/cruvie/kk_etcd_go/utils/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func ApiEtcd() {
	r := gin.Default()
	r.Use(middleware.Cors())

	//swagger
	port := strings.Split(config.GlobalConfig.ServerAddr, ":")[1]
	kk_swagger.InitSwagger(r, port)

	r.Use(middleware.ParseHeader)

	r.POST("/User/Login", handler.Login)

	r.Use(middleware.JWTAuth)

	userAPI := r.Group("/User")
	{
		userAPI.POST("/Logout", handler.Logout)
		userAPI.POST("/AddUser", handler.AddUser)
		userAPI.POST("/DeleteUser", handler.DeleteUser)
		userAPI.POST("/GetUser", handler.GetUser)
		userAPI.POST("/MyInfo", handler.MyInfo)
		userAPI.POST("/UserList", handler.UserList)
		userAPI.POST("/UserGrantRole", handler.UserGrantRole)
	}
	roleAPI := r.Group("/Role")
	{
		roleAPI.POST("/AddRole", handler.AddRole)
		roleAPI.POST("/DeleteRole", handler.DeleteRole)
		roleAPI.POST("/RoleList", handler.RoleList)
		roleAPI.POST("/RoleGet", handler.RoleGet)
		roleAPI.POST("/RoleGrantPermission", handler.RoleGrantPermission)
	}
	kvAPI := r.Group("/KV")
	{
		kvAPI.POST("/KVPutConfig", handler.KVPutConfig)
		kvAPI.POST("/KVGetConfig", handler.KVGetConfig)
		kvAPI.POST("/KVGetConfigList", handler.KVGetConfigList)
		kvAPI.POST("/KVDelConfig", handler.KVDelConfig)
	}

	err := r.Run(config.GlobalConfig.ServerAddr)
	if err != nil {
		log.Panicln("err to run server", err)
		return
	}

}
