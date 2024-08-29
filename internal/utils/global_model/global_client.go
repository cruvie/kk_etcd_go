package global_model

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"go.etcd.io/etcd/client/v3"
	"log/slog"
)

const clientKey = "clientKey"

// SetClient set etcd client
func SetClient(stage *kk_stage.Stage, client *clientv3.Client) {
	stage.Set(clientKey, client)
}

// GetClient get etcd client
func GetClient(stage *kk_stage.Stage) *clientv3.Client {
	client, ok := stage.Get(clientKey)
	if !ok {
		return nil
	}
	return client.(*clientv3.Client)
}

// CloseClient reset etcd client after request
func CloseClient(stage *kk_stage.Stage) {
	client := GetClient(stage)
	if client != nil {
		err := client.Close()
		if err != nil {
			slog.Error("etcd client close error", "err", err)
		}
	}
}
