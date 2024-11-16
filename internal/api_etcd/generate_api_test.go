package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestUser(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hUser",
		Tag:         "user",
		GroupUrl:    "/user/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "login",
				HandlerFuncName: "Login",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "logout",
				HandlerFuncName: "Logout",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "add user",
				HandlerFuncName: "UserAdd",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "delete user",
				HandlerFuncName: "UserDelete",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "get user",
				HandlerFuncName: "GetUser",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "get my info",
				HandlerFuncName: "MyInfo",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "list user",
				HandlerFuncName: "UserList",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "grant role",
				HandlerFuncName: "UserGrantRole",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestRole(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hRole",
		Tag:         "role",
		GroupUrl:    "/role/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//role
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "add role",
				HandlerFuncName: "RoleAdd",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "delete role",
				HandlerFuncName: "RoleDelete",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "list role",
				HandlerFuncName: "RoleList",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "get role",
				HandlerFuncName: "RoleGet",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "grant permission",
				HandlerFuncName: "RoleGrantPermission",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "revoke permission",
				HandlerFuncName: "RoleRevokePermission",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestKV(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hKV",
		Tag:         "kv",
		GroupUrl:    "/kv/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//kv
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "put kv",
				HandlerFuncName: "KVPut",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "get kv",
				HandlerFuncName: "KVGet",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "del kv",
				HandlerFuncName: "KVDel",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "list kv",
				HandlerFuncName: "KVList",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestServer(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hServer",
		Tag:         "server",
		GroupUrl:    "/server/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//server
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "list server",
				HandlerFuncName: "ServerList",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "deregister server",
				HandlerFuncName: "DeregisterServer",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
func TestBackup(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		ApiPkgName:  "kk_etcd_models",
		HandlerName: "hBackup",
		Tag:         "backup",
		GroupUrl:    "/backup/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//backup
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "snapshot",
				HandlerFuncName: "Snapshot",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "snapshot restore",
				HandlerFuncName: "SnapshotRestore",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "snapshot info",
				HandlerFuncName: "SnapshotInfo",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "all kvs backup",
				HandlerFuncName: "AllKVsBackup",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "all kvs restore",
				HandlerFuncName: "AllKVsRestore",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
