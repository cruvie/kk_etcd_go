package ai

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestEtcdAI(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		Tag:      "ai",
		GroupUrl: "/ai/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "query ai",
				HandlerFuncName: "Query",
			},
		},
	}

	for _, api := range apis {
		t.Run(api.apiModel.HandlerFuncName, func(t *testing.T) {
			kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			kk_api_gen.GenerateApiDefProto(apiGroupModel, api.apiModel)
			//kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
