package api_etcd

import (
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_middleware"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/middleware"
	"time"
)

func ApiEtcd(stage *kk_stage.Stage) kk_server.KKRunServer {
	server := kk_http.KKHttpServer{
		Stage:           stage,
		ServerAddr:      fmt.Sprintf(":%d", config.Config.Port),
		ShutdownTimeout: 10 * time.Second,
	}
	server.InitEngine()
	r := server.Engine()
	r.Use(kk_middleware.DefaultCors())
	r.Use(kk_middleware.StageInit(stage))
	//swagger
	kk_swagger.InitSwagger(kk_swagger.Config{
		Options:   nil,
		R:         r,
		Port:      config.Config.Port,
		UrlPrefix: "",
		DebugMode: stage.DebugMode,
	})

	r.Use(middleware.ParseHeader)
	r.Use(middleware.EtcdClient)

	{
		apiGroup := r.Group("user")
		apiGroup.POST("login", login)
		apiGroup.POST("logout", logout)
		apiGroup.POST("userAdd", userAdd)
		apiGroup.POST("userDelete", userDelete)
		apiGroup.POST("getUser", getUser)
		apiGroup.POST("myInfo", myInfo)
		apiGroup.POST("userList", userList)
		apiGroup.POST("userGrantRole", userGrantRole)
	}

	{
		apiGroup := r.Group("role")
		apiGroup.POST("roleAdd", roleAdd)
		apiGroup.POST("roleDelete", roleDelete)
		apiGroup.POST("roleGet", roleGet)
		apiGroup.POST("roleList", roleList)
		apiGroup.POST("roleGrantPermission", roleGrantPermission)
		apiGroup.POST("roleRevokePermission", roleRevokePermission)
	}

	{
		apiGroup := r.Group("kv")
		apiGroup.POST("kVPut", kVPut)
		apiGroup.POST("kVGet", kVGet)
		apiGroup.POST("kVDel", kVDel)
		apiGroup.POST("kVList", kVList)
	}

	{
		apiGroup := r.Group("server")
		apiGroup.POST("serverList", serverList)
		apiGroup.POST("deregisterServer", deregisterServer)
	}

	{
		apiGroup := r.Group("backup")
		apiGroup.POST("snapshot", snapshot)
		apiGroup.POST("snapshotRestore", snapshotRestore)
		apiGroup.POST("snapshotInfo", snapshotInfo)
		apiGroup.POST("allKVsBackup", allKVsBackup)
		apiGroup.POST("allKVsRestore", allKVsRestore)

	}

	{
		apiGroup := r.Group("ai")
		apiGroup.POST("query", query)
	}

	return server.NewKKRunServer()
}
