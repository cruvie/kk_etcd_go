import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_user_kk_etcd.pb.dart';

class ApiUser {
  /// login
  static Future<void> login(
      LoginParam param,
      LoginResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/login", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// logout
  static Future<void> logout(
      LogoutParam param,
      LogoutResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/logout", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// add user
  static Future<void> userAdd(
      UserAddParam param,
      UserAddResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/userAdd", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// delete user
  static Future<void> userDelete(
      UserDeleteParam param,
      UserDeleteResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/userDelete", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// get user
  static Future<void> getUser(
      GetUserParam param,
      GetUserResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/getUser", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// get my info
  static Future<void> myInfo(
      MyInfoParam param,
      MyInfoResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/myInfo", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// list user
  static Future<void> userList(
      UserListParam param,
      UserListResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/userList", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// grant role
  static Future<void> userGrantRole(
      UserGrantRoleParam param,
      UserGrantRoleResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/user/userGrantRole", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
