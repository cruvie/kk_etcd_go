package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_http"
	"gitee.com/cruvie/kk_go_kit/kk_http/kk_global_stage"
	"github.com/cruvie/kk_etcd_go/internal/handler"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"github.com/gin-gonic/gin"
)

var hBackup handler.HBackup

// snapshot
//
//	@Tags			backup
//	@Description	snapshot
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotParam	body		kk_etcd_models.SnapshotParam	true	"SnapshotParam"
//	@Success		200				{object}	kk_etcd_models.SnapshotResponse
//	@Router			/backup/snapshot [post]
func snapshot(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("snapshot")
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
//	@Tags			backup
//	@Description	snapshot restore
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotRestoreParam	body		kk_etcd_models.SnapshotRestoreParam	true	"SnapshotRestoreParam"
//	@Success		200						{object}	kk_etcd_models.SnapshotRestoreResponse
//	@Router			/backup/snapshotRestore [post]
func snapshotRestore(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("snapshotRestore")
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
//	@Tags			backup
//	@Description	snapshot info
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			SnapshotInfoParam	body		kk_etcd_models.SnapshotInfoParam	true	"SnapshotInfoParam"
//	@Success		200					{object}	kk_etcd_models.SnapshotInfoResponse
//	@Router			/backup/snapshotInfo [post]
func snapshotInfo(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("snapshotInfo")
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
//	@Tags			backup
//	@Description	all kvs backup
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			AllKVsBackupParam	body		kk_etcd_models.AllKVsBackupParam	true	"AllKVsBackupParam"
//	@Success		200					{object}	kk_etcd_models.AllKVsBackupResponse
//	@Router			/backup/allKVsBackup [post]
func allKVsBackup(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("allKVsBackup")
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
//	@Tags			backup
//	@Description	all kvs restore
//	@Accept			octet-stream
//	@Produce		octet-stream
//	@Param			AllKVsRestoreParam	body		kk_etcd_models.AllKVsRestoreParam	true	"AllKVsRestoreParam"
//	@Success		200					{object}	kk_etcd_models.AllKVsRestoreResponse
//	@Router			/backup/allKVsRestore [post]
func allKVsRestore(c *gin.Context) {
	stage := kk_global_stage.GetRequestStage(c)
	span := stage.StartTrace("allKVsRestore")
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
