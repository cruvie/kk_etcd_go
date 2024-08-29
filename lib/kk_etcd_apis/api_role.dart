import 'dart:typed_data';

import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:kk_go_kit/kk_util/kk_log.dart';

import '../kk_etcd_models/api_role_kk_etcd.pb.dart';

class ApiRole {
  /// add role
  static Future<void> roleAdd(
      RoleAddParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleAddResponse)? okFunc,
      Function(RoleAddResponse)? errorFunc}) async {
    RoleAddResponse response = RoleAddResponse();
    PBResponse res = await requestFunc("/role/roleAdd", param.writeToBuffer());
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

  /// delete role
  static Future<void> roleDelete(
      RoleDeleteParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleDeleteResponse)? okFunc,
      Function(RoleDeleteResponse)? errorFunc}) async {
    RoleDeleteResponse response = RoleDeleteResponse();
    PBResponse res =
        await requestFunc("/role/roleDelete", param.writeToBuffer());
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

  /// list role
  static Future<void> roleList(
      RoleListParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleListResponse)? okFunc,
      Function(RoleListResponse)? errorFunc}) async {
    RoleListResponse response = RoleListResponse();
    PBResponse res = await requestFunc("/role/roleList", param.writeToBuffer());
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

  /// get role
  static Future<void> roleGet(
      RoleGetParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleGetResponse)? okFunc,
      Function(RoleGetResponse)? errorFunc}) async {
    RoleGetResponse response = RoleGetResponse();
    PBResponse res = await requestFunc("/role/roleGet", param.writeToBuffer());
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

  /// grant permission
  static Future<void> roleGrantPermission(
      RoleGrantPermissionParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleGrantPermissionResponse)? okFunc,
      Function(RoleGrantPermissionResponse)? errorFunc}) async {
    RoleGrantPermissionResponse response = RoleGrantPermissionResponse();
    PBResponse res =
        await requestFunc("/role/roleGrantPermission", param.writeToBuffer());
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

  /// revoke permission
  static Future<void> roleRevokePermission(
      RoleRevokePermissionParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(RoleRevokePermissionResponse)? okFunc,
      Function(RoleRevokePermissionResponse)? errorFunc}) async {
    RoleRevokePermissionResponse response = RoleRevokePermissionResponse();
    PBResponse res =
        await requestFunc("/role/roleRevokePermission", param.writeToBuffer());
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
