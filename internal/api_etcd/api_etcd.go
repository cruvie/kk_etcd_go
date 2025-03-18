package api_etcd

import (
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_middleware"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_ai/query"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_backup/allKVsBackup"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_backup/allKVsRestore"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_backup/snapshot"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_backup/snapshotInfo"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_backup/snapshotRestore"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/kVDel"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/kVGet"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/kVList"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_kv/kVPut"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleAdd"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleDelete"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleGet"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleGrantPermission"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleList"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_role/roleRevokePermission"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_server/deregisterServer"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_server/serverList"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/getUser"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/login"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/logout"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/myInfo"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/userAdd"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/userDelete"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/userGrantRole"
	"github.com/cruvie/kk_etcd_go/internal/mgr_hub/mgr_user/userList"
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
		apiGroup.POST("login", login.Handler)
		apiGroup.POST("logout", logout.Handler)
		apiGroup.POST("userAdd", userAdd.Handler)
		apiGroup.POST("userDelete", userDelete.Handler)
		apiGroup.POST("getUser", getUser.Handler)
		apiGroup.POST("myInfo", myInfo.Handler)
		apiGroup.POST("userList", userList.Handler)
		apiGroup.POST("userGrantRole", userGrantRole.Handler)
	}

	{
		apiGroup := r.Group("role")
		apiGroup.POST("roleAdd", roleAdd.Handler)
		apiGroup.POST("roleDelete", roleDelete.Handler)
		apiGroup.POST("roleGet", roleGet.Handler)
		apiGroup.POST("roleList", roleList.Handler)
		apiGroup.POST("roleGrantPermission", roleGrantPermission.Handler)
		apiGroup.POST("roleRevokePermission", roleRevokePermission.Handler)
	}

	{
		apiGroup := r.Group("kv")
		apiGroup.POST("kVPut", kVPut.Handler)
		apiGroup.POST("kVGet", kVGet.Handler)
		apiGroup.POST("kVDel", kVDel.Handler)
		apiGroup.POST("kVList", kVList.Handler)
	}

	{
		apiGroup := r.Group("server")
		apiGroup.POST("serverList", serverList.Handler)
		apiGroup.POST("deregisterServer", deregisterServer.Handler)
	}

	{
		apiGroup := r.Group("backup")
		apiGroup.POST("snapshot", snapshot.Handler)
		apiGroup.POST("snapshotRestore", snapshotRestore.Handler)
		apiGroup.POST("snapshotInfo", snapshotInfo.Handler)
		apiGroup.POST("allKVsBackup", allKVsBackup.Handler)
		apiGroup.POST("allKVsRestore", allKVsRestore.Handler)

	}

	{
		apiGroup := r.Group("ai")
		apiGroup.POST("query", query.Handler)
	}

	return server.NewKKRunServer()
}
