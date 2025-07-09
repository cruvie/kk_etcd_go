package backup

import (
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_api_gen"
	"testing"
)

func TestBackup(t *testing.T) {
	apiGroupModel := kk_api_gen.ApiGroupModel{
		Tag:      "backup",
		GroupUrl: "/backup/",
	}
	apis := []struct {
		apiModel kk_api_gen.ApiModel
	}{
		//backup
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "snapshot",
				HandlerFuncName: "Snapshot",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "snapshot restore",
				HandlerFuncName: "SnapshotRestore",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "snapshot info",
				HandlerFuncName: "SnapshotInfo",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "all kvs backup",
				HandlerFuncName: "AllKVsBackup",
			},
		},
		{
			apiModel: kk_api_gen.ApiModel{
				Summary:         "all kvs restore",
				HandlerFuncName: "AllKVsRestore",
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
