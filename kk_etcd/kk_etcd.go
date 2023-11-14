package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_func"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"gopkg.in/yaml.v3"
	"log/slog"
)

func InitEtcd(endpoints []string, userName string, password string) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	return service.InitEtcd(stage, endpoints, userName, password)
}

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), kk_etcd_const.Config+configKey)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("configKey", configKey)
		slog.Error("failed to get config", logBody.GetLogArgs()...)
		return err
	}
	if getResponse.Kvs == nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId)
		slog.Warn("failed to get kv, getResponse.Kvs is nil", logBody.GetLogArgs()...)
		return nil
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetTraceId(stage.TraceId).SetError(err).
			SetAny("config", string(getResponse.Kvs[0].Value))
		slog.Error("failed Unmarshal config", logBody.GetLogArgs()...)
		return err
	}
	return nil
}

func RegisterService(registration *kk_etcd_models.ServiceRegistration) error {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	err := service.RegisterService(stage, registration)
	return err
}

// ServerList
// serviceName, should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func ServerList(serviceName string) (serverList *kk_etcd_models.PBListServer, err error) {
	stage := kku_stage.NewStage(nil, kku_func.GetCurrentFunctionName())
	_, servers, err := service.ServerList(stage, serviceName)
	return servers, err
}
