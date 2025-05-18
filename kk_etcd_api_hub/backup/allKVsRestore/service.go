package allKVsRestore

import (
	"bytes"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/util_kv"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"log/slog"
)

// SnapshotRestore todo migrate to etcd cluster manager
func (x *api) service() (err error) {
	span := x.stage.StartTrace("service")
	defer span.End()

	newLog := kk_log.NewLog(&kk_log.LogOption{TraceId: x.stage.TraceId})
	list := &kk_etcd_models.PBListKV{}
	d := json.NewDecoder(bytes.NewReader(x.In.GetFile()))
	d.UseNumber()
	err = d.Decode(list)
	if err != nil {
		slog.Error("failed to Unmarshal kv", newLog.Error(err).Args()...)
		return err
	}
	for _, pbkv := range list.ListKV {
		err := util_kv.PutKV(x.stage, pbkv.GetKey(), pbkv.GetValue())
		if err != nil {
			return err
		}
	}
	return nil
}
