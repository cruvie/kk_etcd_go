package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_base_proto_type"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_response"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/gin-gonic/gin"
)

func Snapshot(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())

	err := service.Snapshot(stage, "")
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}

// todo cmk
func SnapshotDownload(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())
	if !check_user.CheckWritePermission(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have write permission!"}, nil))
		return
	}
	var fileName kk_base_proto_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &fileName); err != nil {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	_, err := service.SnapshotDownload(stage, fileName.Value)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}
func SnapshotRestore(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, &kk_response.KKResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}
	var fileName kk_base_proto_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &fileName); err != nil {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	cmdStr, err := service.SnapshotRestore(stage, fileName.Value)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, &kk_base_proto_type.PBString{Value: cmdStr}))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}

func SnapshotInfo(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())

	var fileName kk_base_proto_type.PBString
	if err := kk_http.ReadProtoBuf(stage, &fileName); err != nil {
		kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, nil))
		return
	}
	info, err := service.SnapshotInfo(stage, fileName.Value)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, &kk_base_proto_type.PBString{Value: info}))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}

func AllKVsBackup(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())

	err := service.AllKVsBackup(stage, "")
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}
func AllKVsRestore(c *gin.Context) {
	stage := kk_stage.NewStage(c, kk_func.GetCurrentFunctionName())

	err := service.AllKVsRestore(stage, "")
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_response.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_response.Fail(stage, nil, &kk_base_proto_type.PBString{Value: err.Error()}))
}
