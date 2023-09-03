package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"

	"gopkg.in/yaml.v3"
)

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), consts.EtcdConfig+configKey)
	if err != nil {
		kku_log.SlogPanic("failed to get kv:", configKey, err)
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		kku_log.SlogPanic("映射配置信息到结构体失败=", err, "读取到的信息=", string(getResponse.Kvs[0].Value))
	}
}
