package allKVsRestore

import (
	"bytes"
	"encoding/json"
	"log/slog"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

// SnapshotRestore todo migrate to etcd cluster manager
func (x *Api) Service(stage *kk_stage.Stage) (err error) {
	span := stage.StartTrace("service")
	defer span.End()

	newLog := kk_stage.NewLog(stage)
	list := &kk_etcd_models.PBListKV{}
	d := json.NewDecoder(bytes.NewReader(x.In.GetFile()))
	d.UseNumber()
	err = d.Decode(list)
	if err != nil {
		slog.Error("failed to Unmarshal kv", newLog.Error(err).Args()...)
		return err
	}
	for _, pbkv := range list.ListKV {
		err := util_kv.PutKV(stage, pbkv.GetKey(), pbkv.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}
