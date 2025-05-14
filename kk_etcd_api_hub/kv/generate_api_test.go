package kv

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestKV(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
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
			//kk_api_gen.GenerateApi(apiGroupModel, api.apiModel)
			//kk_api_gen.GenerateApiDefProto(api.apiModel)
			kk_api_gen.GenerateDartApi(apiGroupModel, api.apiModel)
			//kk_api_gen.GenerateTypescriptApi(apiGroupModel, api.apiModel)
		})
	}
}
