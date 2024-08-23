package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"testing"
)

func TestUser(t *testing.T) {
	apiGroupModel := kk_http.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hUser",
		Tag:         "user",
		GroupUrl:    "/user/",
	}
	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		{
			apiModel: kk_http.ApiModel{
				Description:     "login",
				HandlerFuncName: "Login",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "logout",
				HandlerFuncName: "Logout",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "add user",
				HandlerFuncName: "UserAdd",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "delete user",
				HandlerFuncName: "UserDelete",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "get user",
				HandlerFuncName: "GetUser",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "get my info",
				HandlerFuncName: "MyInfo",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "list user",
				HandlerFuncName: "UserList",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "grant role",
				HandlerFuncName: "UserGrantRole",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			kk_http.GenerateDartApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestRole(t *testing.T) {
	apiGroupModel := kk_http.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hRole",
		Tag:         "role",
		GroupUrl:    "/role/",
	}
	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		//role
		{
			apiModel: kk_http.ApiModel{
				Description:     "add role",
				HandlerFuncName: "RoleAdd",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "delete role",
				HandlerFuncName: "RoleDelete",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "list role",
				HandlerFuncName: "RoleList",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "get role",
				HandlerFuncName: "RoleGet",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "grant permission",
				HandlerFuncName: "RoleGrantPermission",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			kk_http.GenerateDartApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestKV(t *testing.T) {
	apiGroupModel := kk_http.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hKV",
		Tag:         "kv",
		GroupUrl:    "/kv/",
	}
	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		//kv
		{
			apiModel: kk_http.ApiModel{
				Description:     "put kv",
				HandlerFuncName: "KVPut",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "get kv",
				HandlerFuncName: "KVGet",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "del kv",
				HandlerFuncName: "KVDel",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "list kv",
				HandlerFuncName: "KVList",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			kk_http.GenerateDartApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestServer(t *testing.T) {
	apiGroupModel := kk_http.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hServer",
		Tag:         "server",
		GroupUrl:    "/server/",
	}
	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		//server
		{
			apiModel: kk_http.ApiModel{
				Description:     "list server",
				HandlerFuncName: "ServerList",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			kk_http.GenerateDartApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestBackup(t *testing.T) {
	apiGroupModel := kk_http.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hBackup",
		Tag:         "backup",
		GroupUrl:    "/backup/",
	}
	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		//backup
		{
			apiModel: kk_http.ApiModel{
				Description:     "snapshot",
				HandlerFuncName: "Snapshot",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "snapshot restore",
				HandlerFuncName: "SnapshotRestore",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "snapshot info",
				HandlerFuncName: "SnapshotInfo",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "all kvs backup",
				HandlerFuncName: "AllKVsBackup",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				Description:     "all kvs restore",
				HandlerFuncName: "AllKVsRestore",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			kk_http.GenerateDartApi(apiGroupModel, api.apiModel)
		})
	}
}
