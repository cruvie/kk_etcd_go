package service

import (
	"context"
	"encoding/json"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"gopkg.in/yaml.v3"
)

type kvTool struct{}

var toolKV kvTool

func (kvTool) putYaml(key string, structPtr any) error {
	value, err := yaml.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}

func (kvTool) putJson(key string, structPtr any) error {
	value, err := json.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = kk_etcd_client.EtcdClient.Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}
