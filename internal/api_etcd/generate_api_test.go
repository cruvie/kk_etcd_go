package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"testing"
)

func TestGenerateApi(t *testing.T) {

	apis := []struct {
		apiModel kk_http.ApiModel
	}{
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "login",
				HandlerName:     "hUser",
				HandlerFuncName: "Login",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "logout",
				HandlerName:     "hUser",
				HandlerFuncName: "Logout",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "userAdd",
				HandlerName:     "hUser",
				HandlerFuncName: "UserAdd",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "userDelete",
				HandlerName:     "hUser",
				HandlerFuncName: "UserDelete",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "getUser",
				HandlerName:     "hUser",
				HandlerFuncName: "GetUser",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "myInfo",
				HandlerName:     "hUser",
				HandlerFuncName: "MyInfo",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "userList",
				HandlerName:     "hUser",
				HandlerFuncName: "UserList",
				PreUrl:          "/user/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "userGrantRole",
				HandlerName:     "hUser",
				HandlerFuncName: "UserGrantRole",
				PreUrl:          "/user/",
			},
		},
		//role
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "roleAdd",
				HandlerName:     "hRole",
				HandlerFuncName: "RoleAdd",
				PreUrl:          "/role/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "roleDelete",
				HandlerName:     "hRole",
				HandlerFuncName: "RoleDelete",
				PreUrl:          "/role/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "roleList",
				HandlerName:     "hRole",
				HandlerFuncName: "RoleList",
				PreUrl:          "/role/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "roleGet",
				HandlerName:     "hRole",
				HandlerFuncName: "RoleGet",
				PreUrl:          "/role/",
			},
		},
		{
			apiModel: kk_http.ApiModel{
				ApiPkgName:      "kk_etcd_models",
				GroupFuncName:   "roleGrantPermission",
				HandlerName:     "hRole",
				HandlerFuncName: "RoleGrantPermission",
				PreUrl:          "/role/",
			},
		},
		//kv
	}
	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_http.GenerateApi(api.apiModel)
			//kk_http.GenerateDartApi(api.apiModel)
		})
	}
}
