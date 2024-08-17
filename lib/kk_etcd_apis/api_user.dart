import 'dart:typed_data';

import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:kk_go_kit/kk_util/kk_log.dart';

import '../kk_etcd_models/api_user_kk_etcd.pb.dart';

class ApiUser {
  /// login
  static Future<void> login(
      LoginParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(LoginResponse)? okFunc,
      Function(LoginResponse)? errorFunc}) async {
    LoginResponse response = LoginResponse();
    PBResponse res = await requestFunc("/user/login", param.writeToBuffer());
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

  /// logout
  static Future<void> logout(
      LogoutParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(LogoutResponse)? okFunc,
      Function(LogoutResponse)? errorFunc}) async {
    LogoutResponse response = LogoutResponse();
    PBResponse res = await requestFunc("/user/logout", param.writeToBuffer());
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

  /// add user
  static Future<void> userAdd(
      UserAddParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(UserAddResponse)? okFunc,
      Function(UserAddResponse)? errorFunc}) async {
    UserAddResponse response = UserAddResponse();
    PBResponse res = await requestFunc("/user/userAdd", param.writeToBuffer());
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

  /// delete user
  static Future<void> userDelete(
      UserDeleteParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(UserDeleteResponse)? okFunc,
      Function(UserDeleteResponse)? errorFunc}) async {
    UserDeleteResponse response = UserDeleteResponse();
    PBResponse res =
        await requestFunc("/user/userDelete", param.writeToBuffer());
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

  /// get user
  static Future<void> getUser(
      GetUserParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(GetUserResponse)? okFunc,
      Function(GetUserResponse)? errorFunc}) async {
    GetUserResponse response = GetUserResponse();
    PBResponse res = await requestFunc("/user/getUser", param.writeToBuffer());
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

  /// get my info
  static Future<void> myInfo(
      MyInfoParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(MyInfoResponse)? okFunc,
      Function(MyInfoResponse)? errorFunc}) async {
    MyInfoResponse response = MyInfoResponse();
    PBResponse res = await requestFunc("/user/myInfo", param.writeToBuffer());
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

  /// list user
  static Future<void> userList(
      UserListParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(UserListResponse)? okFunc,
      Function(UserListResponse)? errorFunc}) async {
    UserListResponse response = UserListResponse();
    PBResponse res = await requestFunc("/user/userList", param.writeToBuffer());
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

  /// grant role
  static Future<void> userGrantRole(
      UserGrantRoleParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(UserGrantRoleResponse)? okFunc,
      Function(UserGrantRoleResponse)? errorFunc}) async {
    UserGrantRoleResponse response = UserGrantRoleResponse();
    PBResponse res =
        await requestFunc("/user/userGrantRole", param.writeToBuffer());
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
