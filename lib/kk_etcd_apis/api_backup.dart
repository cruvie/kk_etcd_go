import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_backup_kk_etcd.pb.dart';

class ApiBackup {
  /// snapshot
  static Future<void> snapshot(
      SnapshotParam param,
      SnapshotResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/backup/snapshot", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// snapshot restore
  static Future<void> snapshotRestore(
      SnapshotRestoreParam param,
      SnapshotRestoreResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/backup/snapshotRestore", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// snapshot info
  static Future<void> snapshotInfo(
      SnapshotInfoParam param,
      SnapshotInfoResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/backup/snapshotInfo", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// all kvs backup
  static Future<void> allKVsBackup(
      AllKVsBackupParam param,
      AllKVsBackupResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/backup/allKVsBackup", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// all kvs restore
  static Future<void> allKVsRestore(
      AllKVsRestoreParam param,
      AllKVsRestoreResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/backup/allKVsRestore", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
