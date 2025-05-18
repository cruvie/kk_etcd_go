package allKVsBackup

import (
	"context"
	"encoding/json"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

func (x *api) service() (*api_def.AllKVsBackup_Output, error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	list := &kk_etcd_models.PBListKV{}
	getOutput, err := global_model.GetClient(x.stage).Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, kv := range getOutput.Kvs {
		pbKV := &kk_etcd_models.PBKV{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
		//todo only support backup normal kv
		if strings.HasPrefix(pbKV.Key, kk_etcd_models.PBServerType_Http.String()) ||
			strings.HasPrefix(pbKV.Key, kk_etcd_models.PBServerType_Grpc.String()) {
			continue
		}

		list.ListKV = append(list.ListKV, pbKV)
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	timeStr := time.Now().Format(time.DateTime)
	fileName := "etcd_all_kv_" + timeStr + ".json"
	return &api_def.AllKVsBackup_Output{
		Name: fileName,
		File: marshal,
	}, nil
}
