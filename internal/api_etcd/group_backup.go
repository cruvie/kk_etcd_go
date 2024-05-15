package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_func"
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hBackup handler.HBackup

// snapshot
//
//	@Description	snapshot
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotParam	body		kk_etcd_models.SnapshotParam	true	"SnapshotParam"
//	@Success		200					{object}	kk_etcd_models.SnapshotResponse
//	@Router			/snapshot [post]
func snapshot(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.SnapshotParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hBackup.Snapshot(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// snapshotRestore
//
//	@Description	snapshotRestore
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotRestoreParam	body		kk_etcd_models.SnapshotRestoreParam	true	"SnapshotRestoreParam"
//	@Success		200					{object}	kk_etcd_models.SnapshotRestoreResponse
//	@Router			/snapshotRestore [post]
func snapshotRestore(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.SnapshotRestoreParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hBackup.SnapshotRestore(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// snapshotInfo
//
//	@Description	snapshotInfo
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotInfoParam	body		kk_etcd_models.SnapshotInfoParam	true	"SnapshotInfoParam"
//	@Success		200					{object}	kk_etcd_models.SnapshotInfoResponse
//	@Router			/snapshotInfo [post]
func snapshotInfo(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.SnapshotInfoParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hBackup.SnapshotInfo(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// allKVsBackup
//
//	@Description	allKVsBackup
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			AllKVsBackupParam	body		kk_etcd_models.AllKVsBackupParam	true	"AllKVsBackupParam"
//	@Success		200					{object}	kk_etcd_models.AllKVsBackupResponse
//	@Router			/allKVsBackup [post]
func allKVsBackup(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.AllKVsBackupParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hBackup.AllKVsBackup(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}

// allKVsRestore
//
//	@Description	allKVsRestore
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			AllKVsRestoreParam	body		kk_etcd_models.AllKVsRestoreParam	true	"AllKVsRestoreParam"
//	@Success		200					{object}	kk_etcd_models.AllKVsRestoreResponse
//	@Router			/allKVsRestore [post]
func allKVsRestore(c *gin.Context) {
	stage := global_model.GetRequestStage(c)
	span := stage.StartTrace(kk_func.GetCurrentFunctionName())
	defer span.End()
	var param kk_etcd_models.AllKVsRestoreParam
	if err := param.BindCheck(stage); err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	err, resp := hBackup.AllKVsRestore(stage, &param)
	if err != nil {
		kk_http.ResponseErrPB(stage, err)
		return
	}
	kk_http.ResponseSuccessPB(stage, nil, resp)
}
