import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_server_kk_etcd.pb.dart';

class ApiServer {
  /// list server
  static Future<void> serverList(
      ServerListParam param,
      ServerListResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/server/serverList", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// deregister server
  static Future<void> deregisterServer(
      DeregisterServerParam param,
      DeregisterServerResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest(
        "/server/deregisterServer", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
