package kk_etcd

//
//import (
//	api_def2 "github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"
//	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_handlers/allKVsBackup"
//	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_handlers/allKVsRestore"
//	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_handlers/snapshot"
//	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_handlers/snapshotInfo"
//	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_handlers/snapshotRestore"
//	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
//)
//
//type MgrBackup struct{}
//
//func NewMgrBackup() *MgrBackup {
//	return &MgrBackup{}
//}
//
//func (*MgrBackup) Snapshot() (*api_def2.Snapshot_Output, error) {
//	api := snapshot.NewApi(internal_client.GlobalStage)
//	api.In = &api_def2.Snapshot_Input{}
//	return api.Handler()
//}
//
//func (*MgrBackup) SnapshotRestore() (*api_def2.SnapshotRestore_Output, error) {
//	api := snapshotRestore.NewApi(internal_client.GlobalStage)
//	api.In = &api_def2.SnapshotRestore_Input{}
//	return api.Handler()
//}
//
//func (*MgrBackup) SnapshotInfo(param *api_def2.SnapshotInfo_Input) (*api_def2.SnapshotInfo_Output, error) {
//	api := snapshotInfo.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}
//
//func (*MgrBackup) AllKVsBackup() (*api_def2.AllKVsBackup_Output, error) {
//	api := allKVsBackup.NewApi(internal_client.GlobalStage)
//	api.In = &api_def2.AllKVsBackup_Input{}
//	return api.Handler()
//}
//
//// AllKVsRestore will overwrite exist kv
//func (*MgrBackup) AllKVsRestore(param *api_def2.AllKVsRestore_Input) (*api_def2.AllKVsRestore_Output, error) {
//	api := allKVsRestore.NewApi(internal_client.GlobalStage)
//	api.In = param
//	return api.Handler()
//}
