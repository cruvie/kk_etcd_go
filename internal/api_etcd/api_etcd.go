package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
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

	{
		apiGroup := r.Group("user")
		apiGroup.POST(kk_func.GetFunctionName(login), login)
		apiGroup.Use(middleware.JWTAuth)
		apiGroup.POST(kk_func.GetFunctionName(logout), logout)
		apiGroup.POST(kk_func.GetFunctionName(userAdd), userAdd)
		apiGroup.POST(kk_func.GetFunctionName(userDelete), userDelete)
		apiGroup.POST(kk_func.GetFunctionName(getUser), getUser)
		apiGroup.POST(kk_func.GetFunctionName(myInfo), myInfo)
		apiGroup.POST(kk_func.GetFunctionName(userList), userList)
		apiGroup.POST(kk_func.GetFunctionName(userGrantRole), userGrantRole)
	}
	r.Use(middleware.JWTAuth)
	{
		apiGroup := r.Group("role")
		apiGroup.POST(kk_func.GetFunctionName(roleAdd), roleAdd)
		apiGroup.POST(kk_func.GetFunctionName(roleDelete), roleDelete)
		apiGroup.POST(kk_func.GetFunctionName(roleGet), roleGet)
		apiGroup.POST(kk_func.GetFunctionName(roleList), roleList)
		apiGroup.POST(kk_func.GetFunctionName(roleGrantPermission), roleGrantPermission)
	}

	{
		apiGroup := r.Group("kv")
		apiGroup.POST(kk_func.GetFunctionName(kVPut), kVPut)
		apiGroup.POST(kk_func.GetFunctionName(kVGet), kVGet)
		apiGroup.POST(kk_func.GetFunctionName(kVDel), kVDel)
		apiGroup.POST(kk_func.GetFunctionName(kVList), kVList)
	}

	{
		apiGroup := r.Group("server")
		apiGroup.POST(kk_func.GetFunctionName(serverList), serverList)
	}

	{
		apiGroup := r.Group("backup")
		apiGroup.POST(kk_func.GetFunctionName(snapshot), snapshot)
		apiGroup.POST(kk_func.GetFunctionName(snapshotRestore), snapshotRestore)
		apiGroup.POST(kk_func.GetFunctionName(snapshotInfo), snapshotInfo)
		apiGroup.POST(kk_func.GetFunctionName(allKVsBackup), allKVsBackup)
		apiGroup.POST(kk_func.GetFunctionName(allKVsRestore), allKVsRestore)

	}

	kk_http.ServerWithGracefulShutdown(stage, r, config.Config.ServerAddr)

}
