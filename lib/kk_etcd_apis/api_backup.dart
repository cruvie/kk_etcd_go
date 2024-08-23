import 'dart:typed_data';

import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:kk_go_kit/kk_util/kk_log.dart';

import '../kk_etcd_models/api_manage_kk_etcd.pb.dart';

class ApiBackup {
  /// snapshot
  static Future<void> snapshot(
      SnapshotParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(SnapshotResponse)? okFunc,
      Function(SnapshotResponse)? errorFunc}) async {
    SnapshotResponse response = SnapshotResponse();
    PBResponse res =
        await requestFunc("/backup/snapshot", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }

  /// snapshot restore
  static Future<void> snapshotRestore(
      SnapshotRestoreParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(SnapshotRestoreResponse)? okFunc,
      Function(SnapshotRestoreResponse)? errorFunc}) async {
    SnapshotRestoreResponse response = SnapshotRestoreResponse();
    PBResponse res =
        await requestFunc("/backup/snapshotRestore", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }

  /// snapshot info
  static Future<void> snapshotInfo(
      SnapshotInfoParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(SnapshotInfoResponse)? okFunc,
      Function(SnapshotInfoResponse)? errorFunc}) async {
    SnapshotInfoResponse response = SnapshotInfoResponse();
    PBResponse res =
        await requestFunc("/backup/snapshotInfo", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }

  /// all kvs backup
  static Future<void> allKVsBackup(
      AllKVsBackupParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(AllKVsBackupResponse)? okFunc,
      Function(AllKVsBackupResponse)? errorFunc}) async {
    AllKVsBackupResponse response = AllKVsBackupResponse();
    PBResponse res =
        await requestFunc("/backup/allKVsBackup", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }

  /// all kvs restore
  static Future<void> allKVsRestore(
      AllKVsRestoreParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(AllKVsRestoreResponse)? okFunc,
      Function(AllKVsRestoreResponse)? errorFunc}) async {
    AllKVsRestoreResponse response = AllKVsRestoreResponse();
    PBResponse res =
        await requestFunc("/backup/allKVsRestore", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }
}
