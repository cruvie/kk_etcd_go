package handler

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_models/kk_pb_type"
	"github.com/cruvie/kk_etcd_go/internal/handler/service"
	"github.com/cruvie/kk_etcd_go/internal/utils/check_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/gin-gonic/gin"
)

func Snapshot(c *gin.Context) {
	stage := global_model.GetRequestStage(c)

	pBFile, err := service.Snapshot(stage)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, pBFile))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, &kk_pb_type.PBString{Value: err.Error()}))
}

func SnapshotRestore(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	if !check_user.CheckRootRole(stage) {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, &kk_pb_type.PBResponse{
			Msg: "you don't have root role!"}, nil))
		return
	}

	cmdStr, err := service.SnapshotRestore(stage)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, &kk_pb_type.PBString{Value: cmdStr}))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, &kk_pb_type.PBString{Value: err.Error()}))
}

func SnapshotInfo(c *gin.Context) {
	stage := global_model.GetRequestStage(c)

	var pbFile kk_pb_type.PBFile
	if err := kk_http.ReadProtoBuf(stage, &pbFile); err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	info, err := service.SnapshotInfo(stage, &pbFile)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, &kk_pb_type.PBString{Value: info}))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, &kk_pb_type.PBString{Value: err.Error()}))
}

func AllKVsBackup(c *gin.Context) {
	stage := global_model.GetRequestStage(c)

	pbFile, err := service.AllKVsBackup(stage)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, pbFile))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, &kk_pb_type.PBString{Value: err.Error()}))
}
func AllKVsRestore(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	var pbFile kk_pb_type.PBFile
	if err := kk_http.ReadProtoBuf(stage, &pbFile); err != nil {
		kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, nil))
		return
	}
	err := service.AllKVsRestore(stage, &pbFile)
	switch err {
	case nil:
		kk_http.ResponseProtoBuf(c, kk_http.Success(stage, nil, nil))
		return
	}
	kk_http.ResponseProtoBuf(c, kk_http.Fail(stage, nil, &kk_pb_type.PBString{Value: err.Error()}))
}
