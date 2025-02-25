package kk_etcd_models

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *SnapshotParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *SnapshotRestoreParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}

	return nil
}

func (x *SnapshotInfoParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *AllKVsBackupParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}

	return nil
}
func (x *AllKVsRestoreParam) BindCheck(stage *kk_stage.Stage) error {
	err := kk_http.ReadReq(stage, x)
	if err != nil {
		return err
	}

	return nil
}
