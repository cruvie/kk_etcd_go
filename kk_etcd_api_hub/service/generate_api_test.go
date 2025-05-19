package service

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestService(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		Tag:      "service",
		GroupUrl: "/service/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//service
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "list service",
				HandlerFuncName: "ServiceList",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Description:     "deregister service",
				HandlerFuncName: "DeregisterService",
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
