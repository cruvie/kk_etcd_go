package service

import (
	"context"
	"encoding/json"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"gopkg.in/yaml.v3"
)

type kvTool struct{}

var toolKV kvTool

func (kvTool) putYaml(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := yaml.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = global_model.GetClient(stage).Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}

func (kvTool) putJson(stage *kk_stage.Stage, key string, structPtr any) error {
	value, err := json.Marshal(structPtr)
	if err != nil {
		return err
	}
	_, err = global_model.GetClient(stage).Put(context.Background(), key, string(value))
	if err != nil {
		return err
	}
	return nil
}
