import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_role_kk_etcd.pb.dart';

class ApiRole {
  /// add role
  static Future<void> roleAdd(
      RoleAddParam param,
      RoleAddResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/role/roleAdd", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// delete role
  static Future<void> roleDelete(
      RoleDeleteParam param,
      RoleDeleteResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/role/roleDelete", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// list role
  static Future<void> roleList(
      RoleListParam param,
      RoleListResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/role/roleList", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// get role
  static Future<void> roleGet(
      RoleGetParam param,
      RoleGetResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/role/roleGet", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// grant permission
  static Future<void> roleGrantPermission(
      RoleGrantPermissionParam param,
      RoleGrantPermissionResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest(
        "/role/roleGrantPermission", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// revoke permission
  static Future<void> roleRevokePermission(
      RoleRevokePermissionParam param,
      RoleRevokePermissionResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest(
        "/role/roleRevokePermission", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
