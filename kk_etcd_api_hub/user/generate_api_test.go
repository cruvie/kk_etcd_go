package user

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestUser(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		Tag:      "user",
		GroupUrl: "/user/",
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
			kk_api_gen.GenerateApiDefProto(api.apiModel)
			//kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
