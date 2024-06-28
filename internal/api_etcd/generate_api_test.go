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

				HandlerFuncName: "Login",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "Logout",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "UserAdd",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "UserDelete",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "GetUser",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "MyInfo",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "UserList",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "UserGrantRole",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
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

				HandlerFuncName: "RoleAdd",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "RoleDelete",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "RoleList",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "RoleGet",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "RoleGrantPermission",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
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

				HandlerFuncName: "KVPut",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "KVGet",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "KVDel",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "KVList",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
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

				HandlerFuncName: "ServerList",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
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

				HandlerFuncName: "Snapshot",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "SnapshotRestore",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "SnapshotInfo",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "AllKVsBackup",
			},
		},
		{
			apiModel: kk_http.ApiModel{

				HandlerFuncName: "AllKVsRestore",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(apiGroupModel, api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
		})
	}
}
