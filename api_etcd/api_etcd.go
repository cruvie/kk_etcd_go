package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/config"
	"github.com/cruvie/kk_etcd_go/handler"
	"github.com/cruvie/kk_etcd_go/utils/middleware"
	"github.com/gin-gonic/gin"
	"strings"
)

func ApiEtcd() {
	r := gin.Default()
	r.Use(middleware.Cors())

	//swagger
	port := strings.Split(config.Config.ServerAddr, ":")[1]
	kk_swagger.InitSwagger(r, port)

	r.Use(middleware.ParseHeader)

	userAPI := r.Group("User")
	{
		userAPI.POST(kku_func.GetFunctionName(handler.Login), handler.Login)
		userAPI.Use(middleware.JWTAuth)
		userAPI.POST(kku_func.GetFunctionName(handler.Logout), handler.Logout)
		userAPI.POST(kku_func.GetFunctionName(handler.UserAdd), handler.UserAdd)
		userAPI.POST(kku_func.GetFunctionName(handler.UserDelete), handler.UserDelete)
		userAPI.POST(kku_func.GetFunctionName(handler.GetUser), handler.GetUser)
		userAPI.POST(kku_func.GetFunctionName(handler.MyInfo), handler.MyInfo)
		userAPI.POST(kku_func.GetFunctionName(handler.UserList), handler.UserList)
		userAPI.POST(kku_func.GetFunctionName(handler.UserGrantRole), handler.UserGrantRole)
	}
	r.Use(middleware.JWTAuth)
	roleAPI := r.Group("Role")
	{
		roleAPI.POST(kku_func.GetFunctionName(handler.RoleAdd), handler.RoleAdd)
		roleAPI.POST(kku_func.GetFunctionName(handler.RoleDelete), handler.RoleDelete)
		roleAPI.POST(kku_func.GetFunctionName(handler.RoleList), handler.RoleList)
		roleAPI.POST(kku_func.GetFunctionName(handler.RoleGet), handler.RoleGet)
		roleAPI.POST(kku_func.GetFunctionName(handler.RoleGrantPermission), handler.RoleGrantPermission)
	}
	kvAPI := r.Group("KV")
	{
		kvAPI.POST(kku_func.GetFunctionName(handler.KVPutConfig), handler.KVPutConfig)
		kvAPI.POST(kku_func.GetFunctionName(handler.KVGetConfig), handler.KVGetConfig)
		kvAPI.POST(kku_func.GetFunctionName(handler.KVGetConfigList), handler.KVGetConfigList)
		kvAPI.POST(kku_func.GetFunctionName(handler.KVDelConfig), handler.KVDelConfig)
	}

	err := r.Run(config.Config.ServerAddr)
	if err != nil {
		kku_log.SlogPanic("err to run server", err)
		return
	}

}
