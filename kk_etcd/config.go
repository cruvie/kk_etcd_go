package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"gopkg.in/yaml.v3"
	"log/slog"
)

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) error {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName())
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), kk_etcd_const.Config+configKey)
	if err != nil {
		slog.Error("failed to get config", kk_stage.NewLog(stage).Error(err).
			Any("configKey", configKey).Args()...)
		return err
	}
	if getResponse.Kvs == nil {
		slog.Warn("failed to get kv, getResponse.Kvs is nil", kk_stage.NewLog(stage).Args()...)
		return nil
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		slog.Error("failed Unmarshal config", kk_stage.NewLog(stage).Error(err).Any("config", string(getResponse.Kvs[0].Value)).Args()...)
		return err
	}
	return nil
}

// SetConfig set config to etcd
func SetConfig(configKey string, config string) error {
	stage := kk_stage.NewStage(nil, kk_func.GetCurrentFunctionName())

	_, err := kk_etcd_client.EtcdClient.Put(context.Background(), kk_etcd_const.Config+configKey, config)
	if err != nil {
		slog.Error("failed to put config", kk_stage.NewLog(stage).Error(err).
			Any("config", config).Args()...)
		return err
	}
	return nil
}
