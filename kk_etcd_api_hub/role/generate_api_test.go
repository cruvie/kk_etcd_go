package api_def

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestRole(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		Tag:      "role",
		GroupUrl: "/role/",
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
			kk_api_gen.GenerateApiDefProto(api.apiModel)
			//kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
