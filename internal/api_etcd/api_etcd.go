package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_reflect"
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
		apiGroup.POST(kk_reflect.GetFunctionName(login), login)
		apiGroup.Use(middleware.JWTAuth)
		apiGroup.POST(kk_reflect.GetFunctionName(logout), logout)
		apiGroup.POST(kk_reflect.GetFunctionName(userAdd), userAdd)
		apiGroup.POST(kk_reflect.GetFunctionName(userDelete), userDelete)
		apiGroup.POST(kk_reflect.GetFunctionName(getUser), getUser)
		apiGroup.POST(kk_reflect.GetFunctionName(myInfo), myInfo)
		apiGroup.POST(kk_reflect.GetFunctionName(userList), userList)
		apiGroup.POST(kk_reflect.GetFunctionName(userGrantRole), userGrantRole)
	}
	r.Use(middleware.JWTAuth)
	{
		apiGroup := r.Group("role")
		apiGroup.POST(kk_reflect.GetFunctionName(roleAdd), roleAdd)
		apiGroup.POST(kk_reflect.GetFunctionName(roleDelete), roleDelete)
		apiGroup.POST(kk_reflect.GetFunctionName(roleGet), roleGet)
		apiGroup.POST(kk_reflect.GetFunctionName(roleList), roleList)
		apiGroup.POST(kk_reflect.GetFunctionName(roleGrantPermission), roleGrantPermission)
	}

	{
		apiGroup := r.Group("kv")
		apiGroup.POST(kk_reflect.GetFunctionName(kVPut), kVPut)
		apiGroup.POST(kk_reflect.GetFunctionName(kVGet), kVGet)
		apiGroup.POST(kk_reflect.GetFunctionName(kVDel), kVDel)
		apiGroup.POST(kk_reflect.GetFunctionName(kVList), kVList)
	}

	{
		apiGroup := r.Group("server")
		apiGroup.POST(kk_reflect.GetFunctionName(serverList), serverList)
	}

	{
		apiGroup := r.Group("backup")
		apiGroup.POST(kk_reflect.GetFunctionName(snapshot), snapshot)
		apiGroup.POST(kk_reflect.GetFunctionName(snapshotRestore), snapshotRestore)
		apiGroup.POST(kk_reflect.GetFunctionName(snapshotInfo), snapshotInfo)
		apiGroup.POST(kk_reflect.GetFunctionName(allKVsBackup), allKVsBackup)
		apiGroup.POST(kk_reflect.GetFunctionName(allKVsRestore), allKVsRestore)

	}

	kk_http.ServerWithGracefulShutdown(stage, r, config.Config.ServerAddr)

}
