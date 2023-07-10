package kk_etcd

import (
	"context"
	"gopkg.in/yaml.v3"
	"kk_etcd_go/consts"
	"log"
)

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.GlobalConfig)
func GetConfig(configKey string, configStruct any) {
	getResponse, err := EtcdClient.Get(context.Background(), consts.EtcdConfig+configKey)
	if err != nil {
		log.Panicln("failed to get kv:", configKey, err)
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		log.Panicln("映射配置信息到结构体失败=", err, "读取到的信息=", string(getResponse.Kvs[0].Value))
	}
}
