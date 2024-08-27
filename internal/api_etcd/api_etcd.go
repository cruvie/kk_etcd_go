package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/utils/middleware"
)

func ApiEtcd(stage *kk_stage.Stage) {
	r := kk_http.KKGin(stage)
	r.Use(middleware.Cors())
	r.Use(middleware.RequestInterceptor(stage))
	//swagger
	kk_http.InitSwagger(stage, r)

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
	}

	{
		apiGroup := r.Group("backup")
		apiGroup.POST("snapshot", snapshot)
		apiGroup.POST("snapshotRestore", snapshotRestore)
		apiGroup.POST("snapshotInfo", snapshotInfo)
		apiGroup.POST("allKVsBackup", allKVsBackup)
		apiGroup.POST("allKVsRestore", allKVsRestore)

	}

	kk_http.ServerWithGracefulShutdown(stage, r, config.Config.ServerAddr)

}
