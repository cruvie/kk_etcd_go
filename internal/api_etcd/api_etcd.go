package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/middleware"
	"github.com/gin-gonic/gin"
)

func ApiEtcd(stage *kk_stage.Stage) {
	if !config.Config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middleware.Cors())

	//swagger
	kk_swagger.InitSwagger(stage, r, config.Config.ServerAddr)

	r.Use(middleware.ParseHeader)

	userAPI := r.Group("User")
	{
		userAPI.POST(kk_func.GetFunctionName(handler.Login), handler.Login)
		userAPI.Use(middleware.JWTAuth)
		userAPI.POST(kk_func.GetFunctionName(handler.Logout), handler.Logout)
		userAPI.POST(kk_func.GetFunctionName(handler.UserAdd), handler.UserAdd)
		userAPI.POST(kk_func.GetFunctionName(handler.UserDelete), handler.UserDelete)
		userAPI.POST(kk_func.GetFunctionName(handler.GetUser), handler.GetUser)
		userAPI.POST(kk_func.GetFunctionName(handler.MyInfo), handler.MyInfo)
		userAPI.POST(kk_func.GetFunctionName(handler.UserList), handler.UserList)
		userAPI.POST(kk_func.GetFunctionName(handler.UserGrantRole), handler.UserGrantRole)
	}
	r.Use(middleware.JWTAuth)
	roleAPI := r.Group("Role")
	{
		roleAPI.POST(kk_func.GetFunctionName(handler.RoleAdd), handler.RoleAdd)
		roleAPI.POST(kk_func.GetFunctionName(handler.RoleDelete), handler.RoleDelete)
		roleAPI.POST(kk_func.GetFunctionName(handler.RoleList), handler.RoleList)
		roleAPI.POST(kk_func.GetFunctionName(handler.RoleGet), handler.RoleGet)
		roleAPI.POST(kk_func.GetFunctionName(handler.RoleGrantPermission), handler.RoleGrantPermission)
	}
	kvAPI := r.Group("KV")
	{
		kvAPI.POST(kk_func.GetFunctionName(handler.KVPut), handler.KVPut)
		kvAPI.POST(kk_func.GetFunctionName(handler.KVGet), handler.KVGet)
		kvAPI.POST(kk_func.GetFunctionName(handler.KVList), handler.KVList)
		kvAPI.POST(kk_func.GetFunctionName(handler.KVDel), handler.KVDel)
	}
	serverAPI := r.Group("Server")
	{
		serverAPI.POST(kk_func.GetFunctionName(handler.ServerList), handler.ServerList)
	}
	managerAPI := r.Group("Backup")
	{
		managerAPI.POST(kk_func.GetFunctionName(handler.Snapshot), handler.Snapshot)
		managerAPI.POST(kk_func.GetFunctionName(handler.SnapshotDownload), handler.SnapshotDownload)
		managerAPI.POST(kk_func.GetFunctionName(handler.SnapshotRestore), handler.SnapshotRestore)
		managerAPI.POST(kk_func.GetFunctionName(handler.SnapshotInfo), handler.SnapshotInfo)
		managerAPI.POST(kk_func.GetFunctionName(handler.AllKVsBackup), handler.AllKVsBackup)
		managerAPI.POST(kk_func.GetFunctionName(handler.AllKVsRestore), handler.AllKVsRestore)
	}

	kk_http.ServerWithGracefulShutdown(stage, r, config.Config.ServerAddr)

}
