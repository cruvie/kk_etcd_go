package server

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestServer(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
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
			//kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			//kk_api_gen.GenerateApiDefProto(api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			//kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
